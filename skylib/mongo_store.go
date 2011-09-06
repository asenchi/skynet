package skylib

import (
	"log"
	"os"
	"fmt"
	"flag"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
	"sync"
	"time"
)


var MongoServer *string = flag.String("mongoServer", "127.0.0.1", "address of mongo server")
var update sync.Mutex


func RemoveServiceAt(i int) {

	s := NOS.Services[i]
	c := MC.DB("skynet").C("config")

	err := c.Remove(bson.M{"ipaddress": s.IPAddress, "provides": s.Provides, "port": s.Port,  "protocol": s.Protocol})
	if err != nil {
		log.Panic(err)
	}

}

func MongoConnect() {
	var err os.Error
	MC, err = mgo.Mongo("127.0.0.1")
	if err != nil {
		panic(err)
	}

}
// on startup load the configuration file. 
// After the config file is loaded, we set the global config file variable to the
// unmarshaled data, making it useable for all other processes in this app.
func LoadRegistry() {
	NOS = &RegisteredNetworkServers{}
	NOS.Services = make([]*RpcService, 0)
	var service RpcService
	c := MC.DB("skynet").C("config")
	iter, err := c.Find(nil).Iter()
	if err != nil {
		log.Panic(err)
	}
	for {
		err = iter.Next(&service)
		if err != nil {
			break
		}
		fmt.Println("Loaded from MGO: ", service)
		newService := service
		NOS.Services = append(NOS.Services, &newService)
	}
	if err != mgo.NotFound {
		log.Panic(err)
	}
	fmt.Println(len(NOS.Services))

}

func (r *RpcService) AddToRegistry() {

	c := MC.DB("skynet").C("config")

	_,err := c.Upsert(bson.M{"ipaddress": r.IPAddress, "provides": r.Provides, "port": r.Port,  "protocol": r.Protocol},r)
	if err != nil {
		log.Panic(err.String())
	}

}

func (r *RpcService) RemoveFromRegistry() {

	c := MC.DB("skynet").C("config")

	err := c.Remove(bson.M{"ipaddress": r.IPAddress, "provides": r.Provides, "port": r.Port,  "protocol": r.Protocol})
	if err != nil {
		log.Panic(err)
	}

}

func RemoveService(i int) {

	s := NOS.Services[i]
	c := MC.DB("skynet").C("config")

	err := c.Remove(bson.M{"ipaddress": s.IPAddress, "provides": s.Provides, "port": s.Port,  "protocol": s.Protocol})
	if err != nil {
		log.Panic(err)
	}
}

// Watch for remote changes to the config file.  When new changes occur
// reload our copy of the config file.
// Meant to be run as a goroutine continuously.
func WatchRegistry() {

for {

	NewNOS := &RegisteredNetworkServers{}
	NewNOS.Services = make([]*RpcService, 0)
	var service RpcService
	c := MC.DB("skynet").C("config")
	iter, err := c.Find(nil).Iter()
	if err != nil {
		log.Panic(err)
	}
	for {
		err = iter.Next(&service)
		if err != nil {
			break
		}
		fmt.Println("Loaded from MGO: ", service)
		newService := service
		NewNOS.Services = append(NOS.Services, &newService)
	}
	if err != mgo.NotFound {
		log.Panic(err)
	}

	NOS = NewNOS	
	LogDebug("Reloading Services from Registry")
	time.Sleep(3e9)
}

}


var MC *mgo.Session
//  We *could* use this instead someday.
// var DC *doozer.Conn

// Any Store drop-in file would need to define this global function.
func ConnectStore() {
    MongoConnect()
}
