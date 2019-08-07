• func append(slice []Type, elems ...Type) []Type: This function appends elements to the end of a slice.
 If the underlying array is full, it reallocates the content to a bigger slice before appending.
• func cap(v Type) int: Returns the number elements of an array, or of the underlying array if the argument is a slice.
• func close(c chan<- Type): Closes a channel.
• func complex(r, i FloatType) ComplexType: Given two floating points, this returns a complex number.
• func copy(dst, src []Type) int: Copies elements from a slice to another.
• func delete(m map[Type]Type1, key Type): Removes an entry from a map.
• func imag(c ComplexType) FloatType: Returns the imaginary part of a complex number.
• func len(v Type) int: Returns the length of an array, slice, map, string, or channel.
• func make(t Type, size ...IntegerType) Type: Creates a new slice, map, or channel.
• func new(Type) *Type: Returns a pointer to a variable of the specified type, initialized with a zero value.
• func panic(v interface{}): Stops the execution of the current goroutine and, if it's not intercepted, the program.
• func print(args ...Type): Writes the arguments to the standard error.
• func println(args ...Type): Writes the arguments to the standard error and adds a new line at the end.
• func real(c ComplexType) FloatType: Returns the real part of a complex number.
• func recover() interface{}: Stops a panic sequence and captures the panic value.