package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://tamanno:pass@localhost:5432/Foodshop")
	if err != nil {
		log.Println(err)
	}
	return db, nil
}

func CreateProduct(name string, price int) error {
	_, err := DB.Exec("INSERT INTO products(name , price) VALUES ($1, $2)", name, price)
	if err != nil {
		return err
	}
	return nil
}

func Getproducts(name string, price int) (*sql.Rows, error) {
	res, err := DB.Query("SELECT * FROM products ")
	if err != nil {
		return nil, err
	}
	return res, nil
}
