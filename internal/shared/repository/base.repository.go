package repository

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ================================================================================MONGDO DB==================================================================================== */
type MongoDBDatabase struct {
	DatabaseName   string
	CollectionName string
	Payload        map[string]interface{}
	Sort           map[string]interface{}
	PayloadList    []interface{}
	Filter         interface{}
}

func (mongodbdatabase MongoDBDatabase) ConnectMongoDB() (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		return nil, err
	}

	return client.Database(mongodbdatabase.DatabaseName), nil
}

func (mongodbdatabase MongoDBDatabase) DisconnectMongoDB(mongoClient *mongo.Client) error {
	return mongoClient.Disconnect(context.TODO())
}

func (mongodbdatabase MongoDBDatabase) GetMongoDB() ([]map[string]interface{}, error) {
	ctx := context.TODO()
	db, err := mongodbdatabase.ConnectMongoDB()
	if err != nil {
		return nil, err
	}

	findOptions := options.Find()
	findOptions.SetSort(mongodbdatabase.Sort)

	defer mongodbdatabase.DisconnectMongoDB(db.Client())
	cursor, err := db.Collection(mongodbdatabase.CollectionName).Find(ctx, mongodbdatabase.Filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]map[string]interface{}, 0)
	for cursor.Next(ctx) {
		var row map[string]interface{}
		err := cursor.Decode(&row)
		if err == nil {
			result = append(result, row)
		}
	}
	return result, nil
}

func (mongodbdatabase MongoDBDatabase) InsertMongoDB() error {
	db, err := mongodbdatabase.ConnectMongoDB()
	if err != nil {
		return err
	}

	defer mongodbdatabase.DisconnectMongoDB(db.Client())
	_, err = db.Collection(mongodbdatabase.CollectionName).InsertOne(context.TODO(), mongodbdatabase.Payload)
	if err != nil {
		return err
	}
	return nil
}
func (mongodbdatabase MongoDBDatabase) InsertBulkMongoDB() error {
	db, err := mongodbdatabase.ConnectMongoDB()
	if err != nil {
		return err
	}

	defer mongodbdatabase.DisconnectMongoDB(db.Client())
	_, err = db.Collection(mongodbdatabase.CollectionName).InsertMany(context.TODO(), mongodbdatabase.PayloadList)
	if err != nil {
		return err
	}
	return nil
}

func (mongodbdatabase MongoDBDatabase) UpdateMongoDB() error {
	db, err := mongodbdatabase.ConnectMongoDB()
	if err != nil {
		return err
	}

	defer mongodbdatabase.DisconnectMongoDB(db.Client())
	_, err = db.Collection(mongodbdatabase.CollectionName).UpdateOne(context.TODO(), mongodbdatabase.Filter, bson.M{"$set": mongodbdatabase.Payload})
	if err != nil {
		return err
	}
	return nil
}

func (mongodbdatabase MongoDBDatabase) DeleteMongoDB() error {
	db, err := mongodbdatabase.ConnectMongoDB()
	if err != nil {
		return err
	}

	defer mongodbdatabase.DisconnectMongoDB(db.Client())
	_, err = db.Collection(mongodbdatabase.CollectionName).DeleteOne(context.TODO(), mongodbdatabase.Filter)
	if err != nil {
		return err
	}
	return nil
}

/* ================================================================================REDIS==================================================================================== */
type RedisDatabase struct {
	key             string
	value           interface{}
	expiredDuration time.Duration
}

func (redisDatabase RedisDatabase) ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

func (redisDatabase RedisDatabase) SetRedis() error {
	db := redisDatabase.ConnectRedis()
	err := db.Set(context.TODO(), redisDatabase.key, redisDatabase.value, redisDatabase.expiredDuration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisDatabase RedisDatabase) GetRedis() (string, error) {
	db := redisDatabase.ConnectRedis()
	val, err := db.Get(context.TODO(), redisDatabase.key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
