package repository

import (
	"wintefell-service/internal/product/model"
	"wintefell-service/internal/shared/repository"
	"wintefell-service/pkg"
)

var mongoDBProductRepository = repository.MongoDBDatabase{DatabaseName: "winterfell", CollectionName: "products"}

func InsertProductToMongoDB(payload model.Product_Upsert_Model, resultChan chan error) {
	data, _ := pkg.StructToMap(payload)
	mongoDBProductRepository.Payload = data
	resultChan <- mongoDBProductRepository.InsertMongoDB()
}

func InsertProductToElasticSearch(payload model.Product_Upsert_Model, resultChan chan error) {
	data, _ := pkg.StructToMap(payload)
	resultChan <- repository.CreateDocument("products", payload.ID, data)
}
