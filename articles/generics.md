# Generics in go 

To support generics a single function will need a way to declare what types it supports, calling code, on the other hand, will need a way to specigy whether it is calling with an integre or a float map

When you write a function you need to declare type parameters in addition to its ordinary function parameters. These type parameters make the function generic, enabling it to work with arguments of different types, youll call the function with type arguments and ordinary arguments. 

You can simplyfy the calling code by ommitting the type parameters and the compiler will infer the type - this is not possible if the argument has no params 

Type constaints can be moved into their own interface so they can be reused. 

The constraint allows any type implementing the interface 
