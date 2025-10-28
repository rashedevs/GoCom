package repo

import (
	"database/sql"
	"fmt"
	"gocom/domain"
	"gocom/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (title, description, price, img_url) 
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r productRepo) Get(id int) (*domain.Product, error) {
	var product domain.Product
	query := `
		SELECT id, title, description, price, img_url FROM products WHERE id = $1
	`
	err := r.db.Get(&product, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r productRepo) List(page, limit int64) ([]*domain.Product, error) {
	var products []*domain.Product
	offset := (page - 1) * limit
	query := `
		SELECT id, title, description, price, img_url FROM products LIMIT $1 OFFSET $2
	`
	err := r.db.Select(&products, query, limit, offset)
	if err != nil {
		fmt.Println("Inside list method err", err)
		return nil, err
	}
	return products, nil
}

func (r productRepo) Count() (int64, error) {
	var count int64
	query := `
		SELECT COUNT(*) FROM products
	`
	// query := `
	// 	SELECT COUNT(md5(title || description || price || img_url || random() :: text )) FROM products
	// `

	err := r.db.Get(&count, query)
	if err != nil {
		fmt.Println("Inside count method err", err)
		return 0, err
	}
	return count, nil
}

func (r productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `UPDATE products SET title=$1, description=$2, price=$3, img_url=$4 WHERE id=$5`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r productRepo) Delete(id int) error {
	query := `
		DELETE FROM products WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
