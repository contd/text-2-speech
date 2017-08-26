# Text 2 Speech

This is a result of following a tutorial that showed how to use `cgo` to use the api for the `flite` text to speech program.  The `flite` folder is the go package that does this and the `backend` folder is the gRPC server that uses it to generate a wav file given a client request.  The `say` folder is the client implementation that sends whatever text is passed into it to the backend gRPC server.

## Docker

This project also made use of docker and multi-stage docker builds to enable development and deployment of this on non-linux based systems.  The code could be used without docker on a linux system as long as the `flite-dev` packages are installed.

