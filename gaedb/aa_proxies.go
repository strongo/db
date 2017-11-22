package gaedb

import (
	"google.golang.org/appengine/datastore"
	"github.com/strongo/nds"
)

var (
	LoggingEnabled   = true // TODO: move to Context.WithValue()
	mockDB           MockDB
	NewIncompleteKey = datastore.NewIncompleteKey
	NewKey           = datastore.NewKey

	//dbRunInTransaction = datastore.RunInTransaction
	//dbGet = datastore.Get
	//dbGetMulti = datastore.GetMulti
	//dbPut = datastore.Put
	//dbPutMulti = datastore.PutMulti
	//dbDelete = datastore.Delete
	//dbDeleteMulti = datastore.DeleteMulti

	dbRunInTransaction = nds.RunInTransaction
	dbGet              = nds.Get
	dbGetMulti         = datastore.GetMulti
	dbPut              = nds.Put
	dbPutMulti         = nds.PutMulti
	dbDelete           = nds.Delete
	dbDeleteMulti      = nds.DeleteMulti
)