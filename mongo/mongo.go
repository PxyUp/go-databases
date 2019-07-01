package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

var (
	once            sync.Once
	instance        *MongoConnector
	MISSING_DB_URI  = errors.New("MISSING_DB_URI")
	MISSING_DB_NAME = errors.New("MISSING_DB_NAME")
)

type MongoConnector struct {
	session *MgoSession
}

type MgoSession struct {
	session      *mgo.Session
	isConnected  bool
	dataBaseName string
}

func connect(url string) (*mgo.Session, error) {
	conStr := strings.Replace(url, "+", "%2B", -1)
	session, err := mgo.Dial(conStr)
	if err != nil {
		return session, err
	}
	return session, nil
}

func GetInstance() *MongoConnector {
	once.Do(func() {
		instance = createInstance()
	})
	return instance
}

func createInstance() *MongoConnector {
	return &MongoConnector{
		session: &MgoSession{
			isConnected: false,
		},
	}
}

func (c *MongoConnector) Connect(mongoUrl string, databaseName string) error {
	if c.session.isConnected {
		return nil
	}
	c.session.dataBaseName = databaseName
	if len(mongoUrl) == 0 {
		return MISSING_DB_URI
	}
	if len(databaseName) == 0 {
		return MISSING_DB_NAME
	}
	session, err := connect(mongoUrl)
	c.session.session = session
	if err != nil {
		return err
	}
	c.session.isConnected = true
	return nil

}

func (c *MongoConnector) InsertOne(collectionName string, entity interface{}) error {
	sessionCopy := c.session.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(c.session.dataBaseName).C(collectionName)
	err := collection.Insert(&entity)
	return err
}

func (c *MongoConnector) GetOne(collectionName string, findPredicate bson.M, structToDeserialize interface{}) error {
	sessionCopy := c.session.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(c.session.dataBaseName).C(collectionName)
	err := collection.Find(findPredicate).One(structToDeserialize)
	return err
}

func (c *MongoConnector) GetAll(collectionName string, findPredicate bson.M, structToDeserialize interface{}) error {
	sessionCopy := c.session.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(c.session.dataBaseName).C(collectionName)
	err := collection.Find(findPredicate).All(structToDeserialize)
	return err
}

func (c *MongoConnector) GetOneProject(collectionName string, findPredicate bson.M, projectFields bson.M, structToDeserialize interface{}) error {
	sessionCopy := c.session.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(c.session.dataBaseName).C(collectionName)
	err := collection.Find(findPredicate).Select(projectFields).One(structToDeserialize)
	return err
}

func (c *MongoConnector) GetAllProject(collectionName string, findPredicate bson.M, projectFields bson.M, structToDeserialize interface{}) error {
	sessionCopy := c.session.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(c.session.dataBaseName).C(collectionName)
	err := collection.Find(findPredicate).Select(projectFields).All(structToDeserialize)
	return err
}

func (c *MongoConnector) UpdateAll(collectionName string, findPredicate bson.M, updatePredicate bson.M) (*mgo.ChangeInfo, error) {
	sessionCopy := c.session.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(c.session.dataBaseName).C(collectionName)
	info, err := collection.UpdateAll(findPredicate, updatePredicate)
	return info, err
}

func (c *MongoConnector) UpdateOne(collectionName string, findPredicate bson.M, updatePredicate bson.M) error {
	sessionCopy := c.session.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(c.session.dataBaseName).C(collectionName)
	err := collection.Update(findPredicate, updatePredicate)
	return err
}

func (c *MongoConnector) Disconnect() {
	c.session.session.Close()
	c.session.isConnected = false
}
