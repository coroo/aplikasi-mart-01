package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github/coroo/aplikasi-mart-01/models"

	guuid "github.com/google/uuid"
)

type IProductRepository interface {
	Insert(product models.Product) (*models.Product, error)
	FindOneById(id string) (*models.Product, error)
	DeleteOneById(id string) (*models.Product, error)
}

var (
	productQueries = map[string]string{
		"deleteOneById":      "DELETE FROM products WHERE id=?",
		"productFindOneById": "select id,product_code,product_name,product_price from products where id=?",
		"insertProduct":      "insert into products(id,product_code,product_name,product_price) values(?,?,?,?)",
	}
)

type ProductRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewProductRepository(db *sql.DB) IProductRepository {
	ps := make(map[string]*sql.Stmt, len(productQueries))
	for n, v := range productQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &ProductRepository{
		db, ps,
	}
}

func (r *ProductRepository) Insert(product models.Product) (*models.Product, error) {
	id := guuid.New()
	product.Id = id.String()
	res, err := r.ps["insertProduct"].Exec(product.Id, product.ProductCode, product.ProductName, product.ProductPrice)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &product, nil
}

func (r *ProductRepository) FindOneById(id string) (*models.Product, error) {
	row := r.ps["productFindOneById"].QueryRow(id)
	res := new(models.Product)
	err := row.Scan(&res.Id, &res.ProductCode, &res.ProductName, &res.ProductPrice)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *ProductRepository) DeleteOneById(id string) (*models.Product, error) {
	row := r.ps["deleteOneById"].QueryRow(id)
	res := new(models.Product)
	err := row.Scan(&res.Id, &res.ProductCode, &res.ProductName, &res.ProductPrice)
	if err != nil {
		return res, err
	}
	return res, nil
}
