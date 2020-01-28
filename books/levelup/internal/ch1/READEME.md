>int Type Is Special
The int type is special, and its size will depend on the architecture of the computer
you’re using. In practice most systems will be running a 64-bit architecture,
and int will be the same as an int64. This also applies to the uint type.

> Type Conversion
  To convert a type, commonly called “casting” in other languages, you pass the
  variable into parenthesis after the target type. This will only work with types of
  the same category, as these numbers are; if you try to convert between incompatible
  types you’ll get a compiler error.
  
  >nil Pointers
   One element to note when using pointers is that there’s no guarantee you’ll have
   a value passed into the function. You could also pass in nil, which represents
   no value (note, this is very different to false, which only applies to Boolean
   values). While it’s obvious in our example that we’re not passing nil into the
   function, more often than not you’re using variables that have been passed through
   from other areas of the code, so you must check whether a value is actually passed.
   If you try to access or alter a nil pointer, you’ll receive an error along the lines of
   "panic: runtime error: invalid memory address or nil pointer dereference". You
   can avoid this error by checking if the variable is equal to nil in situations where
   a nil value may be passed.
