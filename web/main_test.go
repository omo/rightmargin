package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type TestPerson struct {
	Name  string
	Phone string
}

func TestHello(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if nil != err {
		t.Fail()
		return
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	coll := session.DB("rima-test").C("people")
	to_insert := &TestPerson{"Ale", "+55 53 8116 9639"}
	err = coll.Insert(to_insert)
	if err != nil {
		t.Fail()
		return
	}

	sel := bson.M{"name": "Ale"}
	result := []TestPerson{}

	err = coll.Find(sel).All(&result)
	if err != nil {
		t.Fail()
		return
	}

	if 0 == len(result) {
		t.Fail()
		return
	}

	if result[0].Name != to_insert.Name {
		t.Fail()
		return
	}

	info, err := coll.RemoveAll(sel)
	if err != nil {
		t.Fail()
		return
	}

	if 0 == info.Removed {
		t.Fail()
		return
	}
}
