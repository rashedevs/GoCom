package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"` //add tag
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"image_url" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(pId int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(pId int) error
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

func (r productRepo) Create(p Product) (*Product, error) {
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

func (r productRepo) Get(id int) (*Product, error) {
	var product Product
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

func (r productRepo) List() ([]*Product, error) {
	var products []*Product
	query := `
		SELECT id, title, description, price, img_url FROM products
	`
	err := r.db.Select(&products, query)
	if err != nil {
		fmt.Println("inside list method err", err)
		return nil, err
	}
	return products, nil
}

func (r productRepo) Update(p Product) (*Product, error) {
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
