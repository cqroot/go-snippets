package storage

import "fmt"

type Storage interface {
	Write(content string)
}

type DiskStorageFactory struct{}

func (f DiskStorageFactory) New() Storage {
	return DiskStorage{}
}

type DiskStorage struct{}

func (s DiskStorage) Write(content string) {
	fmt.Println("Disk storage is written with:", content)
}

type MemoryStorageFactory struct{}

func (f MemoryStorageFactory) New() Storage {
	return MemoryStorage{}
}

type MemoryStorage struct{}

func (s MemoryStorage) Write(content string) {
	fmt.Println("Memory storage is written with:", content)
}

type CloudStorageFactory struct{}

func (f CloudStorageFactory) New() Storage {
	return CloudStorage{}
}

type CloudStorage struct{}

func (s CloudStorage) Write(content string) {
	fmt.Println("Cloud storage is written with:", content)
}
