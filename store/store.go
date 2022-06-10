package store

import "github.com/go_server/models"

type DataStore interface {
	Connection
	Write
}

type Connection interface {
	CloseClient() error
}

type Write interface {
	Write(data models.FileData) (string, error)
}
