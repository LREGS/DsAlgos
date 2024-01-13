# Functions 

Functions in Go can take zdero or more arguments. 

To make Go code easier to read, the variable type comes after the variable name. 

Variables in Go are passed by value - this means that when a variable is passed into a function that function recieves a copy of the variable. The function is unable to mutate the original data 

In go you need to ignore all return values that you do not use because the compilker will throw and error for every unused variable. 

For example, if a function returns the radius and middle point of a circle, you should ignore the middle point if you only want the radius 

# Structs 

# Interfaces 

Every now and then you might ned to access the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion. 

