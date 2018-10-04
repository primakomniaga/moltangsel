package product

import (
	"time"
)

type Product struct {
	ProductId   int       `form:"productId" json:"productId"`
	Name        string    `form:"name" json:"name" binding:"required"`
	Descript    string    `form:"desc" json:"desc" binding:"required"`
	Price       int       `form:"price" json:"price" binding:"required"`
	Image       []string  `form:"image" json:"image"`
	Stock       int       `form:"stock" json:"stock"`
	Condition   string    `form:"condition" json:"condition"`
	Url         string    `form:"url" json:"url"`
	CountReview int       `form:"contReview" json:"contReview"`
	Preorder    int       `form:"preOrder" json:"preOrder"`
	Rating      int       `form:"rating" json:"rating"`
	CategoryId  int       `form:"categoryId" json:"categoryId" binding:"required"`
	MerchantId  int       `form:"merchantId" json:"merchantId" binding:"required"`
	InsertBy    int       `form:"insertBy" json:"insertBy"`
	CreateAt    time.Time `form:"createAt" json:"createAt, omitempty"`
	UpdateAt    time.Time `form:"updateAt" json:"-"`
	IsDelete    bool      `form:"isDelete" json:"-"`
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
