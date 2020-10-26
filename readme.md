## golang-grpc
this is tell you how to build GRPC server & client using golang, and how to create simple request and response
between server and client. 

## Installation

install `protobuf-compiler` and `golang-goprotobuf-dev` depedencies in ubuntu
``` bash
# this command required sudo
# make sure you have golang installed in your system

## update you packages
sudo apt update 

## install protobuf compiler in ubuntu
$ sudo apt install protobuf-compiler

## install golang-goprotobuf-dev in ubuntu
$ sudo apt install golang-goprotobuf-dev

```


## Usage
you should add a few golang package using command `go get`
dont forget to change directory to your project

``` bash
# get package from golang repository
$ go get -u github.com/golang/protobuf/protoc_gen
```

### create protobuf file
you can use `make` command to create new protobuf file.
``` bash
## kucing is a file name and will generate protobuf file
## check in your project folder, filename kucing.proto will create
## change kucing to another file you want
$ proto=kucing make create-proto
```

for generate you can use 

``` bash
## kucing is a file name and will generate protobuf file
## check in your project folder, filename kucing.proto will create
## change kucing to another file you want
$ proto=kucing protoc-generate
```
after that, you should manually create kucing.go in to kucing folder.
dont forget to register you new service into main.go

``` go
    ............
    kucing := chatt.Server{}
    kucing.RegisterChattServiceServer(grpcServer, &kucing)
    ............
```
## author
- Fajrul Aulia

