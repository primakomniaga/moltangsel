package product

import (
	"fmt"

	"github.com/jemmycalak/mall-tangsel/service/product"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Resource struct {
	masterDB *sqlx.DB
	slaveDB  *sqlx.DB
}

func New(masterDB, slaveDB *sqlx.DB) *Resource {
	return &Resource{
		masterDB: masterDB,
		slaveDB:  slaveDB,
	}
}

func (r *Resource) CreateProduct(m *product.Product) error {
	fmt.Println(m)
	query := `insert into t_product 
			(
				name, descript, price, image, stock, condition, rating, url,
				countReview, preorder, categoryId, merchantId, insertBy, createAt, isDelete
			)values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
	stmn, err := r.masterDB.Prepare(query)
	if err != nil {
		fmt.Println("Error Prepare query ", err)
		return err
	}
	defer stmn.Close()

	_, err = stmn.Exec(m.Name, m.Descript, m.Price, pq.Array(m.Image), m.Stock, m.Condition, m.Rating, m.Url,
		m.CountReview, m.Preorder, m.CategoryId, m.MerchantId, m.InsertBy, m.CreateAt, m.IsDelete)
	if err != nil {
		fmt.Println("Error Exec ", err)
		return err
	}
	return nil
}
