package go_databases

import (
	"github.com/PxyUp/go-databases/mongo"
	"gotest.tools/assert"
	"testing"
)

func TestGetConnector(t *testing.T) {
	t.Run("GetMongoConnector", func(t *testing.T) {
		t.Run("should return mongo connector", func(t *testing.T) {
			c := GetMongoConnector()
			assert.Equal(t, c, mongo.GetInstance())
		})
		t.Run("should return same instance", func(t *testing.T) {
			assert.Equal(t, mongo.GetInstance(), mongo.GetInstance())
		})
	})
}
