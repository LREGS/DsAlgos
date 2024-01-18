Functions 

Functions in Go can take zdero or more arguments. 

To make Go code easier to read, the variable type comes after the variable name. 

Variables in Go are passed by value - this means that when a variable is passed into a function that function recieves a copy of the variable. The function is unable to mutate the original data 

In go you need to ignore all return values that you do not use because the compilker will throw and error for every unused variable. 

For example, if a function returns the radius and middle point of a circle, you should ignore the middle point if you only want the radius 

# Structs 

# Interfaces 

Every now and then you might ned to access the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion. 

# Channels 

Channels are a typed, thread safe queue, channels allow different gorountines to communicate with each other. 

You can create a channel like ch:= make(chan int)

Send data to a channel:

ch <- 69 - the arrow determins the direction the data is returning towards 

Recieve data from a channel v := <- channel 

This reads and removes a value from the channe, and saves it into the v variable. This operation will  block until there is a value in the channel to be read 

A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug. 

Empty structs are often used as tokens in Go programs. In this context, a token is a unary value. In other words in doesnt care what is passed it cares when and if something is passed. 

To wait and block until something is sent on a channel use the following syntax 
<-channel 

This will block until it pops a single item off the channel, then continue, discarding the item. 

Channels can be optionall buffered - you can provide a buffer length as the second argument to the make function. 

Sending on a buffered channel only blocks when the buffer is full. Recieving blocks on when the buffer is empty. 

## Closing channel 
Channels can be closed by a sender. 

You can check if a channel is closed similar to the ok syntax by doing v, ok := <-ch 

ok is false if the channel is empty or closed 

Sending on a closed channel will cause a paninc. A panic on the main gorouting will casue the entire program ot crash and a apanic in any other goroutine will cause that goroutine to crash. 

# Mutexed in GO 

Mutexes allow us ot lock access to data. Data sources like maps are not thread safe - if there are multiple go routines writing to the same map it can cause a lot of bugs. For this reason you need to lock the map so that whilst you're making edits you know you're editing the most up to date version. Anything else also wanted to edit it will in turn need to wait for the most up to date version. 

# Generics in GO 

Go doies not suport clases. For a while it wasn't easy to re-use code in Go. This is because functions had to be declared with types, however, often times a function doesnt care about the data its being given and similar functionality can be used for more than one type. 

Generics allow you to use variables to refer to specific types. This allows you to write abstractions of functions and write less code in the process.

 



