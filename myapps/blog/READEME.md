https://vsupalov.com/go-folder-structure/

Go code lives in cmd, internal and pkg. Everything else has a clear place as well. Here’s why there are separate folders
 for subsets of .go files:

The cmd directory contains subfolders with a single main.go each. Each folder inside of cmd. If your project produces an
 executable binary, you will know exactly that your cmd folder will be the place to look. When you want to build or run 
 something, it will look like go run cmd/binaryname/main.go - pretty obvious!

The internal folder contains packages which are specific to your project. They are not meant to be imported outside of 
this project, unlike the contents of the pkg folder.

The pkg directory contains packages which you might want to make accessible to other projects. It can have subfolders 
by topic.

Our test directory does not contain Go tests! Unit tests live right besides the code they are supposed to test. Instead, 
this is the place to put scripts for external blackbox and smoke tests. I like to use Python for my scripting needs, 
as you see.

The ui folder is used for the frontend part of the project. Templates, static assets or the project’s SPA can go here. 
In this case, I added be a single HTML file which is served straight out of the Go code.