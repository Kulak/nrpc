package main

//go:generate protoc -I. -I ../../../../.. --plugin ../../protoc-gen-nrpc/protoc-gen-nrpc --go_out . --nrpc_out . alloptions.proto

func main() {

}
