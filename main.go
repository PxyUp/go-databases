package go_databases

import (
	"github.com/PxyUp/go-databases/mongo"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type MongoDbConnector interface {
	Connect(mongoUrl string, databaseName string) error
	Disconnect()
	Count(collectionName string, findPredicate bson.M) (int, error)
	InsertOne(collectionName string, entity interface{}) error
	Remove(collectionName string, findPredicate bson.M) error
	RemoveAll(collectionName string, findPredicate bson.M) (*mgo.ChangeInfo, error)
	UpdateOne(collectionName string, findPredicate bson.M, updatePredicate bson.M) error
	UpdateAll(collectionName string, findPredicate bson.M, updatePredicate bson.M) (*mgo.ChangeInfo, error)
	GetOne(collectionName string, findPredicate bson.M, structToDeserialize interface{}) error
	GetAll(collectionName string, findPredicate bson.M, structToDeserialize interface{}) error
	GetOneProject(collectionName string, findPredicate bson.M, projectFields bson.M, structToDeserialize interface{}) error
	GetAllProject(collectionName string, findPredicate bson.M, projectFields bson.M, structToDeserialize interface{}) error
	GetIterator(collectionName string, findPredicate bson.M, opts *mongo.MongoOptions) *mongo.MongoIterator
}

type CreateMongoConnector MongoDbConnector

const (
	MONGO_CONNECTOR = "MONGO_CONNECTOR"
)

func GetMongoConnector() CreateMongoConnector {
	return mongo.GetInstance()
}
