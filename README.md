# go-boiler
> You have some eggs? This boiler lets you boil them...

This repository implements a clean/screaming architecture, exposing core logic through an HTTP server using Gin.

In a screaming architecture (like a feature oriented code base) the most important parts which are entities and actions that can be performed within entities are immediately visible thanks to the folder structure. Much like in an onion architecture, each layer adds more abstraction.

Layers of this onion must communicate via clear interfaces. This abstraction allows use of dependency injection, which in turn facilitates testing of the codebase.
