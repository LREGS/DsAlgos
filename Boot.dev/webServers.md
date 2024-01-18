# Web Servers 

A web server is just a computer that serves data over a network on the internet :

Servers run sofware that listsn for incoming requests from clients whne a request is recieved the server response with the requested data. 

Servers should be able to handle many requests at the same time 

Go routines make it easier to handle multiple request at once. 

Remember, go routines allow for concurrent programming - so if a function is waiting for data to come from a channel 

func handler(nums <- data){
  for num := range nums{
      go handleNums(num)
}
}

This would add a new go routine every time a new piece of data was added to the handler. 
Without this functionality you would have to wait for the nums to be handled one after the othe other. If you had 10,000 concurrent users on your site and each action took even 0.5 seconds it becomes very noticable. 

Go was built with concurrency in mind - it was made to handle googles web traffic / infrastrcuture and go routines are a good fit for servers before they're more leightweight than a traditional thread and still make use of multi cores. 

This is different to the node async model - this is a little bit slower than go concurrency. Althought it will do just fine for crud applications once you start doing cpu intensive work it will be slow. 

With django or flask you use WSGI that will spawn a new instance of pyto run concurrent tasks - again a little more complicated than go. 

:

Servers area always running - a lot of code I've written so far has been more run, do a thing, maybe another things and then exit. 

With servers they're sat waiting for requests 

Testing with servers is a little different too because you cannot just write a bunch of print statements to see what is happening you have to actually test the http endpoints to see if they're working. 

# Custom Hanndlers 

Any type with a ServeHTTP method that matches the http.HandlerFunc signature is a http.Handler. To handle a HTTP request all a function needs is a way to write the response and request itself. 


# Lost loads of notes  :(



fds:
