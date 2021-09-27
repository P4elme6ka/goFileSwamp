package filesRegistry

import (
	"errors"
	"github.com/google/uuid"
	"log"
)

var registry map[uuid.UUID]FileDescription

type FileDescription struct {
	Name string
}

func NewFile(name string) FileDescription {
	nuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	registry[nuid] = FileDescription{
		Name: name,
	}
	return registry[nuid]
}

func (fd *FileDescription) GetFileName(id uuid.UUID) (string, error) {
	if fd, ok := registry[id]; ok {
		return fd.Name, nil
	} else {
		return "", errors.New("no such file")
	}
}

func GetFilesList() map[uuid.UUID]FileDescription {
	return registry
}
