package main

import (
	"go-snippets/design-patterns/creational-simple_factory/storage"
)

func main() {
	var s storage.Storage

	for _, typ := range []storage.StorageType{
		storage.DiskStorageType,
		storage.MemoryStorageType,
		storage.CloudStorageType,
	} {
		s = storage.New(typ)
		s.Write("test")
	}
}
