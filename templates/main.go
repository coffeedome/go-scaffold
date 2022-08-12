package templates

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/animals", CreateAnimal)

	router.Run()

}
