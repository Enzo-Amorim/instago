package login

import "github.com/gin-gonic/gin"

func register() {

}

func connectAccount(c *gin.Context) {
	email := c.PostForm("email")
	senha := c.PostForm("senha")
}
