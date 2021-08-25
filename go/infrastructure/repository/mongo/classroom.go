package mongo

import "gopkg.in/mgo.v2/bson"

type ClasroomModel struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}
