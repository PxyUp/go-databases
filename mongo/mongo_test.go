package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const (
	mongoString = "mongodb://localhost:27017"
	mongoDbName = "testing"
	collection  = "user"
)

type user struct {
	Name string `bson:"name"`
}

func TestGetInstance(t *testing.T) {
	t.Run("GetInstance", func(t *testing.T) {
		t.Run("should return instance", func(t *testing.T) {
			instance := GetInstance()
			assert.Equal(t, reflect.TypeOf(instance), reflect.TypeOf(&MongoConnector{}))
		})
		t.Run("shoud return same instance", func(t *testing.T) {
			assert.Equal(t, GetInstance(), GetInstance())
		})
	})
}

func TestMongoConnector(t *testing.T) {
	t.Run("MongoConnector", func(t *testing.T) {
		instance := GetInstance()
		t.Run("Connect", func(t *testing.T) {
			t.Run("should return error", func(t *testing.T) {
				err := instance.Connect("", "")
				assert.Equal(t, err, MISSING_DB_URI)
			})
			t.Run("should return error", func(t *testing.T) {
				err := instance.Connect("asdasd", "")
				assert.Equal(t, err, MISSING_DB_NAME)
			})
			t.Run("should return error", func(t *testing.T) {
				err := instance.Connect("test", "test")
				assert.NotNil(t, err)
			})
			t.Run("should connect", func(t *testing.T) {
				err := instance.Connect(mongoString, mongoDbName)
				assert.Equal(t, err, nil)
				assert.Equal(t, instance.session.isConnected, true)
			})
			t.Run("should return nil", func(t *testing.T) {
				err := instance.Connect(mongoString, mongoDbName)
				assert.Equal(t, err, nil)
			})

		})
		t.Run("InsertOne", func(t *testing.T) {
			t.Run("should insert record", func(t *testing.T) {
				user := &user{
					"Test",
				}

				err := instance.InsertOne(collection, user)
				assert.Equal(t, err, nil)
			})
		})
		t.Run("GetOne", func(t *testing.T) {
			t.Run("should return record", func(t *testing.T) {
				user := &user{}
				err := instance.GetOne(collection, bson.M{
					"name": "Test",
				}, user)
				assert.Equal(t, err, nil)
				assert.Equal(t, "Test", user.Name)
			})
		})
		t.Run("GetOneProject", func(t *testing.T) {
			t.Run("should return record", func(t *testing.T) {
				user := &user{}
				err := instance.GetOneProject(collection, bson.M{
					"name": "Test",
				}, bson.M{
					"name": 1,
				}, user)
				assert.Equal(t, err, nil)
				assert.Equal(t, "Test", user.Name)
			})
		})
		t.Run("GetAll", func(t *testing.T) {
			t.Run("should return array of record", func(t *testing.T) {
				users := []user{}
				err := instance.GetAll(collection, bson.M{
					"name": "Test",
				}, &users)
				assert.Equal(t, err, nil)
				assert.Equal(t, "Test", users[0].Name)
				assert.Equal(t, 1, len(users))
			})
		})
		t.Run("GetAllProject", func(t *testing.T) {
			t.Run("should return array of record", func(t *testing.T) {
				users := []user{}
				err := instance.GetAllProject(collection, bson.M{
					"name": "Test",
				}, bson.M{
					"name": 1,
				}, &users)
				assert.Equal(t, err, nil)
				assert.Equal(t, "Test", users[0].Name)
				assert.Equal(t, 1, len(users))
			})
		})
		t.Run("UpdateOne", func(t *testing.T) {
			t.Run("should update record", func(t *testing.T) {
				err := instance.UpdateOne(collection, bson.M{
					"name": "Test",
				}, bson.M{
					"$set": bson.M{
						"name": "Test4",
					},
				}, )
				assert.Equal(t, err, nil)
				user := &user{}
				err = instance.GetOne(collection, bson.M{
					"name": "Test4",
				}, user)
				assert.Equal(t, err, nil)
				assert.Equal(t, "Test4", user.Name)
			})
		})
		t.Run("UpdateAll", func(t *testing.T) {
			t.Run("should update all record", func(t *testing.T) {
				usr := &user{
					"Test",
				}

				err := instance.InsertOne(collection, usr)
				assert.Equal(t, err, nil)
				_, err = instance.UpdateAll(collection, bson.M{}, bson.M{
					"$set": bson.M{
						"name": "Test4",
					},
				}, )
				assert.Equal(t, err, nil)
				users := []user{}
				err = instance.GetAll(collection, bson.M{
					"name": "Test4",
				}, &users)
				assert.Equal(t, err, nil)
				assert.Equal(t, "Test4", users[0].Name)
				assert.Equal(t, "Test4", users[1].Name)
				assert.Equal(t, 2, len(users))
			})
		})
		t.Run("Disconnect", func(t *testing.T) {
			t.Run("should disconnect", func(t *testing.T) {
				instance.Disconnect()
				assert.Equal(t, instance.session.isConnected, false)
			})
		})
	})
}
