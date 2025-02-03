package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product_Upsert_Model struct {
	ID             string                  `json:"id" bson:"id" validate:"required,uuid4"`
	Name           string                  `json:"name" bson:"name" validate:"required"`
	ImageURL       string                  `json:"image_url" bson:"image_url" validate:"required,url"`
	Price          float64                 `json:"price" bson:"price" validate:"required,gt=0"`
	Discount       float64                 `json:"discount" bson:"discount" validate:"gte=0,lte=100"`
	Currency       string                  `json:"currency" bson:"currency" validate:"required,oneof=IDR USD EUR"`
	Description    string                  `json:"description" bson:"description" validate:"required"`
	IsAvailable    bool                    `json:"is_available" bson:"is_available" validate:"required"`
	CategoryID     string                  `json:"category_id" bson:"category_id" validate:"required"`
	SubcategoryID  string                  `json:"subcategory_id" bson:"subcategory_id" validate:"required"`
	RegisteredDate primitive.DateTime      `bson:"registereddate,omitempty" json:"registereddate,omitempty"`
	UpdatedDate    primitive.DateTime      `bson:"updateddate,omitempty" json:"updateddate,omitempty"`
	Variants       []Product_Variant_Model `json:"variants,omitempty" bson:"variants,omitempty"`
}

type Product_Variant_Model struct {
	ID          string                     `json:"id" bson:"id" validate:"required,uuid4"`
	Name        string                     `json:"name" bson:"name" validate:"required"`
	ImageURL    string                     `json:"image_url" bson:"image_url" validate:"required,url"`
	Subvariant  []Product_Subvariant_Model `json:"subvariant" bson:"subvariant" validate:"dive"`
	Stock       int                        `json:"stock" bson:"stock" validate:"gte=0"`
	IsAvailable bool                       `json:"is_available" bson:"is_available" validate:"required"`
}

type Product_Subvariant_Model struct {
	ID      string `json:"id" bson:"id" validate:"required,uuid4"`
	Subname string `json:"subname" bson:"subname" validate:"required"`
	Stock   int    `json:"stock" bson:"stock" validate:"gte=0"`
}
