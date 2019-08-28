https://github.com/irahardianto/service-pattern-go

https://github.com/golang-standards/project-layout

https://github.com/katzien/go-structure-examples

https://flaviocopes.com/go-filesystem-structure/

https://www.youtube.com/watch?v=oL6JBUk6tj0

https://github.com/Mindinventory/Golang-Project-Structure        可以的 https://www.mindinventory.com/blog/golang-project-structure/

https://github.com/manuelkiessling/go-cleanarchitecture/tree/master/src

## 博客 
https://aaf.engineering/go-web-application-structure-pt-1/

https://blog.golang.org/modules2019

https://dave.cheney.net/2018/07/14/taking-go-modules-for-a-spin

https://www.reddit.com/r/golang/comments/a1ivp7/help_go_modules_folder_structure/

https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2

https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

https://vsupalov.com/go-folder-structure/

https://forum.golangbridge.org/t/how-should-i-structure-packages-for-a-multiple-binary-web-application/665/9
Calmh jakob borg
>
    This is something we’ve been fighting and iterating on… Current layout on a reasonably sized project looks like this:
    
    Godeps/
      ... lots of dependencies. This would probably be vendor/ if we started again today.
    cmd/
      ... lots of binaries / main packages
    etc/
      ... various example configs and stuff
    gui/
      ... a web app that is compiled into one of the binaries
    lib/
      ... our internal packages, some with subdirectories of their own.
      ... this was "internal" at some point, but since this is nowadays enforced and we
      ... actually have a few external uses, it became "lib"
    script/
      ... various build supporting Go scripts
    build.sh
    build.go
    README
    AUTHORS
    ... etc standard toplevel stuff
    There are a few more top level directories for stuff like graphics assets and so on as well, but not relevant to the Go side of things. So all the Go code lives under cmd/ and lib/, apart from build scripts.
    
    This all builds with standard GOPATH (plus a prepend for Godeps), so internal packages are seen as github.com/organization/project/lib/thepackage.
    
    I’ve been looking into gb as well, but I’m not entirely convinced yet.
    
    
https://hackernoon.com/basic-monorepo-design-in-go-e9ba1cb8e4e6

>  ├──Makefile
   ├──Readme.md
   ├──cmd
   │  ├──parser
   │  │  └──parser.go
   │  └──terminator
   │     └──terminator.go
   ├──pkg
   │  ├──apps
   │  │  ├──parser
   │  │  │  └──parser.go
   │  │  └──terminator
   │  │     └──terminator.go
   │  ├──metrics
   │  │  └──metrics.go
   │  └──sentry
   │     └──sentry.go
   ├──bin
   │  ├──parser
   │  └──terminator
   ├──scripts
   │  └──deploy.sh
   └──docs
      ├──components
      │  ├──sentry.md
      │  └──redshift.md
      └──apps
         ├──parser.md
         └──terminator.md
         
> Makefile
  When projects become large and more individuals start contributing code it becomes difficult to enforce everyone to use the same commands. A Makefile can help standardize a nice shortcut list for contributors to use. Here are some tasks you might want to include in your projects:
  
  build: Builds the binary for your go project and places it in the bin folder.
  test: Runs tests using specific flags or tools. With Go’s testing package you can also specify test groups if your code base become very large.
  deploy: Runs a deploy script or trigger to deploy a specific project. It usually makes sense to keep your build and deploy tasks separate. Your Makefile can always specify a specific order to run them.
  clean: Cleans out old build objects and artifacts.
  cmd Folder
  The cmd folder is meant to hold your application entry points (main funcs). These entry point files are responsible for setting up and configuring each application.
  
  pkg Folder
  All internal packages which do not make sense as a standalone package should be located in a pkg folder. The organization of files within this folder will depend on your organization. Here are some general tips.
  
  Keep the KISS principle in mind while organizing.
  Teams will often be tempted to design projects in familiar language folder styles, but it is important to remember Go is the language and having a package for every data model might not make sense.
  Having generic packages named utils or misc are not clear to developers unfamiliar with your code base. If the code base becomes large enough developers might even forget what functions are held in these generic packages and end up adding duplicate functions.
  Splitting up a single package into separate files with good names often helps organization, but having 20 files which are only 20 lines long probably doesn’t make sense and will most likely annoy developers. Sometimes one file is enough.
  Designing a package which holds all error types might seem like a good organizational tactic, but it increases package entanglement and forces developers to jump file to file.
  bin Folder
  For many projects the bin folder is responsible for holding binary files. So it makes sense to store any binary files within this folder. Especially our compiled projects.
  
  scripts Folder
  Having a folder dedicated to scripts works well when paired with a Makefile. Teams can have complex deploy scripts which trigger multiple actions and keep the code separate from the Makefile. This keeps your Makefile clear and simple.
  
  docs Folder
  Every Go project should use Godoc’s methods, but there are benefits to making documentation which is not contained in your Go files. Having higher level docs can be tailored to users who do not care to see the inner workings. They also can house deployment processes, contributing guidelines, and other information which wouldn’t make sense to contain within the packages documentation.             