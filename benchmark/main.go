package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

type Response struct {
	Elapsed  int64      `json:"elapsed"`
	Average  float64    `json:"average"`
	Products []*Product `json:"products"`
}

var pool *sql.DB // Database connection pool.
var poolTime int64 = 0
var poolCount int64 = 0

func main() {
	dsn := "postgres://root:secret@localhost:5432/testbenchmark?sslmode=disable"
	var err error
	pool, err = sql.Open("postgres", dsn)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(2 * time.Minute)
	pool.SetMaxIdleConns(50)
	pool.SetMaxOpenConns(90)

	_, stop := context.WithCancel(context.Background())
	defer stop()

	router := gin.Default()
	router.StaticFile("/", "./index.html")

	// Handlers...

	router.GET("/products/pooled", func(c *gin.Context) {
		startTime := time.Now()
		// Query the database for all products
		rows, err := pool.Query("SELECT id, name, price, description, category FROM mock_data limit 1000")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products, err := scanProducts(rows)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		elapsed := time.Since(startTime).Microseconds()
		poolCount++
		poolTime += elapsed
		c.JSON(http.StatusOK, Response{Elapsed: elapsed, Average: float64(poolTime / poolCount), Products: products})
	})

	// Start the HTTP server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Unable to start HTTP server: %v\n", err)
	}

}

func scanProducts(rows *sql.Rows) ([]*Product, error) {
	defer rows.Close()

	products := make([]*Product, 0)
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}
