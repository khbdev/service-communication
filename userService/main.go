package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}


type Product struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    UserID string `json:"user_id"`
}


type UserWithProducts struct {
    User     User      `json:"user"`
    Products []Product `json:"products"`
}


var users = []User{
    {ID: "1", Name: "Azizbek", Email: "azizbek@example.com"},
    {ID: "2", Name: "Ali", Email: "ali@example.com"},
}

func main() {
    r := gin.Default()

    
    r.GET("/users/:id/details", func(c *gin.Context) {
        userID := c.Param("id")

        // 1️⃣ User topish
        var foundUser *User
        for _, u := range users {
            if u.ID == userID {
                foundUser = &u
                break
            }
        }

        if foundUser == nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
  products, err := fetchProductsWithRetry(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
       
 
      
        result := UserWithProducts{
            User:     *foundUser,
            Products: products,
        }

       
        c.JSON(http.StatusOK, result)
    })

 
    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
    })

    r.Run(":8082") 
}


func fetchProductsWithRetry(userID string) ([]Product, error) {
    var products []Product
  productServiceURL := fmt.Sprintf("http://product-service:8083/products?user_id=%s", userID)

   
    for attempt := 1; attempt <= 3; attempt++ {

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        req, err := http.NewRequestWithContext(ctx, "GET", productServiceURL, nil)
        if err != nil {
            return nil, err
        }

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            fmt.Printf("Attempt %d failed: %v\n", attempt, err)
            time.Sleep(1 * time.Second)
            continue
        }

        defer resp.Body.Close()
        if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
            return nil, err
        }

     
        return products, nil
    }

    return nil, fmt.Errorf("failed to fetch products after 3 attempts")
}
