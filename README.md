# Go File Server

## Purpose
A small file server using just golang.

Built using [grpc](https://grpc.io/)

## Starting

A docker compose is in the root, it builds, instantiates and start up a single rpc server, an rpc client, and a web server.

The web server serves up files from `files`

## Developing

In order to modify and change the RPC spec, you need to install `protobuf`. On OSX this can be done using homebrew: `brew install protobuf`

`protobuf` takes `.proto` spec files and converts then into an rpc interface that is imported in both the client and the server.

To generate that interface: `protoc -I rpc/ rpc/rpc.proto --go_out=plugins=grpc:rpc`

Then you should be able to run the client and server in seperate terminals and hit `localhost:3030/test.txt`


