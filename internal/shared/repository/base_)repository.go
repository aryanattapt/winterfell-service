package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
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

/* ================================================================================ELASTIC SEARCH==================================================================================== */
var esClient *elasticsearch.Client

func initElasticsearch() {
	// Define custom Elasticsearch URL and options
	cfg := elasticsearch.Config{
		Addresses: []string{
			// Replace this with the URL of your Elasticsearch cluster
			"http://localhost:9200", // Localhost, change for remote cluster if needed
			// Add more addresses if you have multiple nodes
			// "http://es-node-1:9200",
			// "http://es-node-2:9200",
		},
		// Uncomment to add authentication if needed
		// Username: "your_username",
		// Password: "your_password",
	}

	// Initialize the Elasticsearch client with the custom configuration
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Assign the client to the global variable
	esClient = client

}

type Model map[string]interface{}

func CreateDocument(index string, id string, body Model) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("Error marshaling data: %s", err)
	}

	req, error := esClient.Index(
		index,
		bytes.NewReader(jsonData),
		esClient.Index.WithDocumentID(id),
		esClient.Index.WithRefresh("true"),
	)

	if error != nil {
		return error
	}

	if req.IsError() {
		return fmt.Errorf("Error indexing document: %s", req)
	}
	return nil
}

func GetDocumentByID(index string, id string) (Model, error) {
	var doc Model
	req, error := esClient.Get(index, id)
	if error != nil {
		return nil, error
	}

	if req.IsError() {
		return nil, fmt.Errorf("Error getting document: %s", req)
	}

	if err := json.NewDecoder(req.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("Error parsing response body: %s", err)
	}
	return doc, nil
}

func UpdateDocument(index string, id string, body Model) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("Error marshaling data: %s", err)
	}

	req, error := esClient.Update(index, id, bytes.NewReader(jsonData))
	if error != nil {
		return error
	}

	if req.IsError() {
		return fmt.Errorf("Error updating document: %s", req)
	}
	return nil
}

func DeleteDocument(index string, id string) error {
	req, error := esClient.Delete(index, id)
	if error != nil {
		return error
	}

	if req.IsError() {
		return fmt.Errorf("Error deleting document: %s", req)
	}
	return nil
}

type SearchQueryElasticModel struct {
	Query    string            `json:"query"`
	Filters  map[string]string `json:"filters"`
	Fields   []string          `json:"fields"`
	Sort     string            `json:"sort"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
}

func buildSearchQueryElasticSearch(query SearchQueryElasticModel) (map[string]interface{}, error) {
	// Base query structure
	queryBody := map[string]interface{}{
		"from": query.Page * query.PageSize,
		"size": query.PageSize,
	}

	// Add full-text query if specified
	if query.Query != "" {
		queryBody["query"] = map[string]interface{}{
			"match": map[string]interface{}{
				"_all": query.Query, // You can specify other fields here
			},
		}
	}

	// Add filters if specified
	if len(query.Filters) > 0 {
		filter := []map[string]interface{}{}
		for field, value := range query.Filters {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					field: value,
				},
			})
		}
		if len(filter) > 0 {
			if queryBody["query"] == nil {
				queryBody["query"] = map[string]interface{}{}
			}
			queryBody["query"].(map[string]interface{})["bool"] = map[string]interface{}{
				"filter": filter,
			}
		}
	}

	// Add fields if specified
	if len(query.Fields) > 0 {
		queryBody["stored_fields"] = query.Fields
	}

	// Add sorting if specified
	if query.Sort != "" {
		queryBody["sort"] = []map[string]interface{}{
			{
				query.Sort: map[string]interface{}{
					"order": "asc", // or "desc"
				},
			},
		}
	}

	return queryBody, nil
}

// Function to query Elasticsearch dynamically
func GetElasticSearch(index string, searchQuery SearchQueryElasticModel) ([]Model, error) {
	// Build the query
	queryBody, err := buildSearchQueryElasticSearch(searchQuery)
	if err != nil {
		return nil, err
	}

	// Convert the query body to JSON
	queryJSON, err := json.Marshal(queryBody)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling query: %s", err)
	}

	// Perform the search request
	req, error := esClient.Search(
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(bytes.NewReader(queryJSON)),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)

	if error != nil {
		return nil, error
	}

	// Check for errors
	if req.IsError() {
		return nil, fmt.Errorf("Error executing query: %s", req)
	}

	// Parse the response
	var res map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("Error parsing response body: %s", err)
	}

	// Extract the hits
	hits := res["hits"].(map[string]interface{})["hits"].([]interface{})
	var documents []Model
	for _, hit := range hits {
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})
		documents = append(documents, doc)
	}

	return documents, nil
}
