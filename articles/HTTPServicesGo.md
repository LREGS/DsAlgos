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


