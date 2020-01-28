
> While you can’t add your own methods to built-in types, you can create your own
  types from other types.
  
  ### Error Handling
  So what exactly is an error? Well, it’s a great example of an interface. An error is
  actually just a normal Go interface, and is any type that implements the Error
  method:
  ~~~go
  
  type error interface {
  Error() string
  }
~~~

This leads the way for us to create our own error types. We can either extend an
existing type, or create a struct with extra data.

> Constructors
  Notice how we created an empty string array for the Actors? These new functions
  provide the opportunity to initialize empty structures that the type requires to
  function, and it’s common practice to initialize empty arrays or maps within the
  new function.
