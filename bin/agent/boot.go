package main

import "nolan/storage"

func main() {
	m := storage.Master{
		Name: "hello",
	}
	println(m.Name)
}
