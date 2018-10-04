package product

import (
	"fmt"

	"errors"

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
	query := `insert into t_product (
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

func (r *Resource) Products() ([]product.Product, error) {
	var products []product.Product

	querys := `select productid, name, descript, price, image, stock, condition, rating,
	 url, countreview, preorder, categoryid, merchantid, createat from t_product where isdelete = $1`
	rows, err := r.masterDB.Queryx(querys, false)
	if err != nil {
		fmt.Println("error query ", err)
		return nil, err
	}

	for rows.Next() {
		var m product.Product
		err := rows.Scan(&m.ProductId, &m.Name, &m.Descript, &m.Price, pq.Array(&m.Image), &m.Stock,
			&m.Condition, &m.Rating, &m.Url, &m.CountReview, &m.Preorder, &m.CategoryId, &m.MerchantId, &m.CreateAt)
		if err != nil {
			fmt.Println("error loop data ", err)
			return nil, err
		}
		products = append(products, m)
	}
	return products, nil
}

func (r *Resource) Product(idProduct int) (*product.Product, error) {
	var m product.Product
	querys := `select productid, name, descript, price, image, stock, condition, rating,
		url, countreview, preorder, categoryid, merchantid, createat from t_product where productid = $1 and isdelete = $2`

	if err := r.masterDB.QueryRow(querys, idProduct, false).Scan(&m.ProductId, &m.Name, &m.Descript, &m.Price, pq.Array(&m.Image), &m.Stock,
		&m.Condition, &m.Rating, &m.Url, &m.CountReview, &m.Preorder, &m.CategoryId, &m.MerchantId, &m.CreateAt); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *Resource) EditProduct(m *product.Product) error {
	querys := `update t_product set name = $1, descript= $2, price=$3, stock=$4, condition=$5, url=$6, preorder=$7,
	 categoryid=$8, merchantid=$9, updateat=$10 where productid = $11`

	sttmn, err := r.masterDB.Prepare(querys)
	if err != nil {
		return err
	}
	defer sttmn.Close()

	res, err := sttmn.Exec(m.Name, m.Descript, m.Price, m.Stock, m.Condition, m.Url, m.Preorder, m.CategoryId, m.MerchantId, m.UpdateAt, m.ProductId)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("id not found")
	}
	return nil
}

func (r *Resource) DeleteProduct(idproduct int) error {
	querys := `update t_product set isdelete = $1 where productid = $2`

	sttmn, err := r.masterDB.Prepare(querys)
	if err != nil {
		return err
	}
	defer sttmn.Close()

	res, err := sttmn.Exec(true, idproduct)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("id not found")
	}
	return nil
}

func (r *Resource) UpdateImage(idProduct int, image []string) error {
	querys := `update t_product set image = $1 where productid = $2`
	sttmn, err := r.masterDB.Prepare(querys)
	if err != nil {
		return err
	}
	defer sttmn.Close()

	result, err := sttmn.Exec(pq.Array(image), idProduct)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if count == 0 {
		return errors.New("id not found")
	}
	return nil
}

func (r *Resource) ProductLimit(page, limit int) ([]product.Product, error) {
	var products []product.Product

	querys := `select productid, name, descript, price, image, stock, condition, rating,
		url, countreview, preorder, categoryid, merchantid, createat from t_product
		where isdelete = $1
		 order by productid limit $2 offset $3`

	rows, err := r.masterDB.Queryx(querys, false, limit, page)
	if err != nil {
		fmt.Println("error query ", err)
		return nil, err
	}

	for rows.Next() {
		var m product.Product
		err := rows.Scan(&m.ProductId, &m.Name, &m.Descript, &m.Price, pq.Array(&m.Image), &m.Stock,
			&m.Condition, &m.Rating, &m.Url, &m.CountReview, &m.Preorder, &m.CategoryId, &m.MerchantId, &m.CreateAt)
		if err != nil {
			fmt.Println("error loop data ", err)
			return nil, err
		}
		products = append(products, m)
	}

	return products, nil
}
