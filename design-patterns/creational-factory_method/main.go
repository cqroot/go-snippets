package main

import (
	"go-snippets/design-patterns/creational-factory_method/storage"
)

func main() {
	var s storage.Storage

	s = storage.DiskStorageFactory{}.New()
	s.Write("test")

	s = storage.MemoryStorageFactory{}.New()
	s.Write("test")

	s = storage.CloudStorageFactory{}.New()
	s.Write("test")
}
