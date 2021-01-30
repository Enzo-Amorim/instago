package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type account struct {
	nome  string `json: name`
	email string `json: name`
	senha string `json: name`
}

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		nome := c.PostForm("nome")
		email := c.PostForm("email")
		senha := c.PostForm("senha")
		account := account{
			nome,
			email,
			senha,
		}

		fmt.Println(account)

		c.JSON(200, gin.H{
			"nome":  account.nome,
			"email": account.email,
			"senha": account.senha,
		})

	})

	r.Run()
}
