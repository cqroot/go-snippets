package storage

import "fmt"

type Storage interface {
	Write(content string)
}

type StorageType int

const (
	DiskStorageType StorageType = iota
	MemoryStorageType
	CloudStorageType
)

func New(typ StorageType) Storage {
	switch typ {
	case MemoryStorageType:
		return MemoryStorage{}
	case CloudStorageType:
		return CloudStorage{}
	}
	return DiskStorage{}
}

type DiskStorage struct{}

func (s DiskStorage) Write(content string) {
	fmt.Println("Disk storage is written with:", content)
}

type MemoryStorage struct{}

func (s MemoryStorage) Write(content string) {
	fmt.Println("Memory storage is written with:", content)
}

type CloudStorage struct{}

func (s CloudStorage) Write(content string) {
	fmt.Println("Cloud storage is written with:", content)
}
