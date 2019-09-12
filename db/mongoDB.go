package db

import "gopkg.in/mgo.v2"

func mongoConnect() {
	session, err := mgo.Dial("127.0.0.1")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("usergo").C("profiles")

	// Index
	index := mgo.Index{
		Key:        []string{"msisdn"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	return c
}
