package skylib

import (
	"os"
	"log"
	"flag"
	"launchpad.net/mgo"
)


var MongoServer *string = flag.String("mongoServer", "127.0.0.1", "address of mongo server")


// This is a doozer adapter to our skylib.Store interface.
// It's pretty trivial, as our API is doozer, but we
// need this because the Event structs are technically
// not the same type.
type MongoStore struct {
	session	*mgo.Session
}


// Constructor for DoozerStore.
// Connect based on DoozerServer cmd-line flag.
func MongoConnect() *MongoStore {
 	session, err := mgo.Mongo(*MongoServer)
        if err != nil {
                log.Panic(err)
        }
	ds := &MongoStore{session: session}
	DC = ds
	return ds
}

// Responds with the first change made to any file matching path, a glob pattern, on or after rev.
func (me *MongoStore) Wait(glob string, rev int64) (ev *Event, err os.Error) {

	return
}

// Returns the current revision.
func (me *MongoStore) Rev() (rev int64, err os.Error) {
	return
}

// Sets the contents of the file at path to value, as long as rev is greater than or equal to the file's revision. Returns the file's new revision.
func (me *MongoStore) Set(file string, oldrev int64, body []byte) (newrev int64, err os.Error) {
	return
}

// Del deletes the file at path if rev is greater than or equal to the file's revision.
func (me *MongoStore) Del(file string, rev int64) (err os.Error) {
	return
}

// Gets the contents (value) and revision (rev) of the file at path in the specified revision (rev). If rev is not provided, get uses the current revision.
func (me *MongoStore) Get(file string, prev *int64) (body []byte, rev int64, err os.Error) {
	return
}