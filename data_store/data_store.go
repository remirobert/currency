package data_store

import "gopkg.in/mgo.v2"

type DataStore struct {
	Session *mgo.Session
}

func NewDataStore(url string) (store *DataStore, err error) {
	s, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &DataStore{Session: s}, nil
}
