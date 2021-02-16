package interfaces

import "github.com/gin-gonic/gin"

type IGetArticles interface {
	GetArticles(c *gin.Context) ([]string, error)
}
