# Communication on the Web 

HTTP is hte standard protocol for communicating over the web. 

It has not always been easy to share photos over the internet and it is really important to understant that photo sharing has become really simple these days because of the use of HTTP.

Most URLs have HTTP at the beginning of the address, however, it is not the only protocol that exists. 

# Web Addresses 

IP addresses help your computer tell the router what website it is tring to connect too 

Domain Name System (DNS) maps urls to IP addresses so someone can just type google.com rather than knowing all the ip addresses by heart 

# URLS
a Uniform Resource Identifier (URI) is a unique sequence of characters that identifies a resource that is accessed via the internet. URI come in two types: URLs/URNs

## URLS 

URLS are made up of: protocol, username, password, domain, port, path, query, fragment 

http://william:password@domain:8080/path?sort=query#id 

Not all the segements are required everytime - its uncommon to enter your login details in this way for facebook for example

The required aspects are protocol and domain - the rest are optional

The protocol (or scheme) is the first part of the url and defined the rules in which data being communicated is displayed, encoded or formatted 

### Ports 

ports are managed by the os and allow the server to segment incoming requests and datastreams, so you can direct all web traffic to web server, and you can access your db on a different port etc 

You cant bind multiple pieces of software to the same port 

You can have up to 65k ports at the same time 

In static sites url paths work closer to like a file path system of the website when websites basically just served pages of html to the client 

Now, urls represent pages/resources 

Query parameters are not always present and they rarely change the page itself, but will often change the contents on the page 

Although, query parameters can be used for anything that chooses to use them for just like the urls path. 

# Async 

Async is when you want the code to continue executing until the other function has finished. 

When reading from RAM its really quick - so when working with ram you can do it synchronously because you dont really need to wait for it 

If you're working from a Disk - you can do this async because it would say allow the ui to continue working whilst the file is loading in the background for example 

If you're working on the network - like http request its going to be depend on how far away the server is from the client to the spped. 

HTTP/network communication is slow enough to be noticed by humans, so its almost always done async because otherwise all our websites would be pausing every time we wait for data to be inputted or outputted 

JS have promises that can resolved or rejected - with ansync its important to have proper error handling because if your request is never returned your ui needs to know and handle this event correctly so its not waiting for data indefinitely 

IN JS you can create a promise object with Promise() but more likely is just async function name(){
    function 
}

This then can be used with async functions and await is used to wait for the response from the async function 

Await can be used in place of the resolve syntax to say assign a new variable to the results of an async function 

await can only be used inside of asny functions

# Error Handling

try/catch in js to handle errors that might arise from the code 

If it fails an error object will be returned, and its up to you how to handle it - you can return the error message by accessing err.message

you can use await within a try catch block

# HTTP HEADERS

a HTTP header allows clients and servers to pass additional information with each request or response. 
Essentially they're just case sensitive key-value pairs

HTTP requests from browsers using include client, operating system, preffered language 

As devs we can also define custom headers in each request 

The headers API allows us to perform varios actions on our requests and response headers such as retrieving, setting and removing them,. we can access headers through the request.headers and response.headers properties. 

# URL PATHS 

The path comes immediately after the domain name/port if one is provided 

# RESTful APIs 

REST is a set of conventions about how http should be used based on the request functions. 

## Seperate and Agnostic 

Resources are transfered via well-recognised language agnostic client/server interactions. A RESTful style means the implementation of the client and server can be done independently of one another. 

## Stateless 

This means the server does not need to know what state the client is in nor does the client need to know what state the server is in. This is forced by acting with resources and not commands. 

## Paths in REST 

in a RESTful API the last section of the path of a URL should specify which resource is being accessed. Then, as we talked about in the methods chatper, depending on whether the request is a GET, POST, PUT, DELETE the resource is read, created, updated or deleted. 


# HTTPS 

S stands for secure - this will encrpt the data being sent via the http protocol and allow server/client to share sensitive data 



