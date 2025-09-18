package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Product struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    UserID  string `json:"user_id"`  
}

var products = []Product{
    {ID: "101", Name: "Laptop", UserID: "1"},
    {ID: "102", Name: "Phone", UserID: "1"},
    {ID: "103", Name: "Book", UserID: "2"},
}

func main() {
    r := gin.Default()

    r.GET("/products", func(c *gin.Context) {
        userID := c.Query("user_id")
        var result []Product

        for _, p := range products {
            if p.UserID == userID {
                result = append(result, p)
            }
        }

        c.JSON(http.StatusOK, result)
    })

    r.Run(":8083")
}
