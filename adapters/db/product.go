package db

import (
	"github.com/georgetonjr/arquitetura-hexagonal/application"
  "context"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where product id=?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
