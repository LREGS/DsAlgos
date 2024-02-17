https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years

The backbone of any go service is the server. A new server function makes a 'main' http handler. Usually I have one per service and Irely pn http routes to divert traffic to the right handles within each service because:

- New Server takes all dependencies as arguments 
- It turns a handler which can be a dedicated type for more complex situations 
- it usually configures its own muxer and calls out to routes.go 

The new server constructor is responsible for all the top level http stuff that applies to endpoints like cors, middleware, and logging

// Constructor 
` func NewServer(
    logger * logger
    store
    config  
) http.handler{
  mux := http.NewServerMux()
add routes(
  routes
)

var handler http.Handler = mux 
handler = someMiddleware(handler)
return handler 
}

Because it is responsible for all the top level stuff:

handler = logging.LoggingMiddleWare(logger, handler)

return Handler

This would access the logger being passed into the new server and use its dependency 

You then set up the server by exposing it to the go http package 

srv := NewServer(args)

httpServer := &http.Server{
  addr : address
  handler : srv 
}

By passing the arguments as dependencies you get a little bit more type safety vs passing them as structs. 

# Mapping API Surfaces in routes.go 

This is the only place routes are listed in their codebase. It is very helpful to go to one point in the codebase and see the entire api surface. 

At the moment, I have a central api file, and then the complete api for each element within my app is stored within an api dir. But, all the routes are together in the store, and all the interfaces in the stores 

# func main only calls run 

Use run as the main function instead (I think so run() error {} and you can control when to stop the program if an error is returned )

Main should also create a context and can also pass os arguments to the run command. This makes programmers easier to test because code can call run to execute program, and control what happens just through passing different arguments. 

# Maker funcs return the handler 

`// handleSomething handles one of those web requests
// that you hear so much about.
func handleSomething(logger *Logger) http.Handler {
	thing := prepareThing()
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// use thing to handle request
			logger.Info(r.Context(), "msg", "handleSomething")
		}
	)
}`

This pattern gives each handler its own closure environment. You can do initialization work in this space, and the data will be available to the handlers when they are called.

In a pattern like this you should only read and use shared data, not modify the original of any of the data or you will need a mutex or something. 

# Handling decoding/encoding in one place 

Every service will need to decode request boadies and encode the response bodies. This is a solid abstraction. You can write very basic generics in your helper function that can encode and decode any type. 

# Validating data

A simple single function interface can be used to create a validator and this is then any struct that has the valid method is an object that can be validated. 

The valid method will check the object and return a map of any issues in the data it found - this could be a dodgy email string, empty username etc 

This should avoid checking any fields in databases because that is probably a little too complicated for this function 


# Adapter patter for middleware 

Middleware functions take a handler and return a new one that can run code before and or after calling the original handler. Or it can decide not to call the original handler at all.

`func adminOnly(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !currentUser(r).IsAdmin {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	})
}`

# Sometimes middleware can be returned 

Using the above approach is okay when the middleware function is simple. But if the function has a lot of dependencies (logger, db, api client etc) then it can be good to have a function that returns the middleware function. 

`func newMiddleware(
	logger Logger,
	db *DB,
	slackClient *slack.Client,
	rroll []byte,
) func(h http.Handler) http.Handler`

The return type of this function is the handler func we will call when setting up our routes.

`middleware := newMiddleware(logger, db, slackClient, rroll)
mux.Handle("/route1", middleware(handleSomething(handlerSpecificDeps))
mux.Handle("/route2", middleware(handleSomething2(handlerSpecificDeps))`

You can then create a type for middleware like 

`type middleware func(h http.Handler) http.handler`

we can see a similar pattern in the mattermost codebase. They define a type handler which consists of a server (which in this codebase also holds the service struct) and handlerfunc and then some options. This is then used as a basis for the different handlers required, such as those for tasks that require api authentication and those that dont. 

When these are then used within the api a handlerfunc (remember this is anything with that has a response writer and request) is passed into the APIhandler func, which initiates the handle type and returns the middleware function as the handler for this route. This means that this handler has all the required dependencies to complete its task, without you having to include the dependencies in the fufunction definition every time. 

In mattermost the user api probably had 30+ endpoints that all use the APISessionRequired handler func, it would be annoying to have to write the function out everytime to include auth, logger, db etc when instead you can wrap it within the APISessionRequired function, which creates and returns a handler type that i:: designed for this task. 


# The opportunity to hide the request/response types away 

IF an endopoint has its own request and response types, usually they're only useful for that handler. If thats the case they can be defined inside of the function. 


# Use inline requests/response types for additional storytelling in tests 


# Sync.one to defer setup 

If you have expensive code in some of your handlers it can be good toi defer this until the handler is actually called to speed up start-times of your application. sync.Once means the code is only executed once and other calls will block until its finished. 

If the handler is never called the expensive work will never be completed. Although this improves start-up times it also moves the initialization from start-up to run time and this might have other implications depending on the application. 

# Designing for testability 

It can be good to instaed of testing the individual handler funcs, to test the api endpoints and mimic the user of users. Instead of calling the functions we can test and actually call the run function, start the server etc and the endpoitns will be tested in their totality this way your code is going through all layers of the application and even using the db etc. 



