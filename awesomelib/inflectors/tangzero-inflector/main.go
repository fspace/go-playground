package main

import "github.com/tangzero/inflector"

func main() {
	print(inflector.Pluralize("person"), "\n")        // "people"
	print(inflector.Singularize("posts"), "\n")       // "post"
	print(inflector.Camelize("posts_hello"), "\n")    // "PostsHello"
	print(inflector.Underscorize("PostsHello"), "\n") // "posts_hello"
}
