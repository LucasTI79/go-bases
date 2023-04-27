package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func main() {
	product := Product{"1", "Product 1", 100.00}

	err := SaveProduct(product)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.POST("/products", createProduct)
	e.Logger.Fatal(e.Start(":3000"))
}

func createProduct(c echo.Context) error {
	product := Product{}
	if err := c.Bind(&product); err != nil {
		return err
	}

	err := SaveProduct(product)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, product)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello world!"))
}

func SaveProduct(product Product) error {
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES(?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}
