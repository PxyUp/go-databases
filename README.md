# Databases Golang
[![codecov](https://codecov.io/gh/PxyUp/go-databases/branch/master/graph/badge.svg)](https://codecov.io/gh/PxyUp/go-databases)

This repository will have connectors for most popular databases

# How to use

All instance it is singleton, you can multiple use them in another file, without any worries about connections counts

```bash
go get github.com/PxyUp/go-databases
```

```go
import (
	"github.com/PxyUp/go-databases"
)

func main() {
    instance := go_databases.GetMongoConnector()
    err := instance.Connect(mongoString, mongoDbName)
    assert.Equal(t, err, nil)
    user := &user{
        "Test",
    }
    err = instance.InsertOne(collection, user)
    assert.Equal(t, err, nil)
}
```

# Mongo

```bash
type MongoDbConnector interface {
	Connect(mongoUrl string, databaseName string) error
	Disconnect()
	InsertOne(collectionName string, entity interface{}) error
	UpdateOne(collectionName string, findPredicate bson.M, updatePredicate bson.M) error
	UpdateAll(collectionName string, findPredicate bson.M, updatePredicate bson.M) (*mgo.ChangeInfo, error)
	GetOne(collectionName string, findPredicate bson.M, structToDeserialize interface{}) error
	GetAll(collectionName string, findPredicate bson.M, structToDeserialize interface{}) error
	GetOneProject(collectionName string, findPredicate bson.M, projectFields bson.M, structToDeserialize interface{}) error
	GetAllProject(collectionName string, findPredicate bson.M, projectFields bson.M, structToDeserialize interface{}) error
}
```
