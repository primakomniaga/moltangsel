package product

import (
	"time"
)

type Product struct {
	ProductId   string    `form:"productId"`
	Name        string    `form:"name" binding:"required"`
	Descript    string    `form:"desc" binding:"required"`
	Price       int       `form:"price" binding:"required"`
	Image       []string  `form:"image"`
	Stock       int       `form:"stock"`
	Condition   string    `form:"condition"`
	Url         string    `form:"url"`
	CountReview int       `form:"contReview"`
	Preorder    int       `form:"preOrder"`
	Rating      int       `form:"rating"`
	CategoryId  int       `form:"categoryId" binding:"required"`
	MerchantId  int       `form:"merchantId" binding:"required"`
	InsertBy    int       `form:"insertBy" binding:"required"`
	CreateAt    time.Time `form:"createAt"`
	UpdateAt    time.Time `form:"updateAt"`
	IsDelete    bool      `form:"isDelete"`
}

type Images struct {
	Image string `json:"image"`
}

func (s *Service) NewProduct() *Product {
	return &Product{
		CountReview: 0,
		Rating:      0,
		CreateAt:    time.Now(),
		IsDelete:    false,
	}
}
