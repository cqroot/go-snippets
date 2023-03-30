package main

import (
	"github.com/cqroot/go-snippets/design-patterns/creational-simple_factory/storage"
)

func main() {
	var s storage.Storage

	s = storage.New(storage.DiskStorageType)
	s.Write("test")

	s = storage.New(storage.MemoryStorageType)
	s.Write("test")

	s = storage.New(storage.CloudStorageType)
	s.Write("test")
}
