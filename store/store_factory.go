package store

import "github.com/go_server/store/sqlite"

type DatastoreOptions struct {
	StoreName string
	DbDir     string
	DbName    string
	Addr      string
}

var Store DataStore

var availDataStores = map[string]func(options DatastoreOptions) DataStore{
	"sqlite": func(options DatastoreOptions) DataStore {
		client, err := sqlite.NewClient(options.DbDir)
		if err != nil {
			panic(err)
		}
		return client
	},
}

func NewDataStoreClient(options DatastoreOptions) DataStore {
	return availDataStores[options.StoreName](options)
}
