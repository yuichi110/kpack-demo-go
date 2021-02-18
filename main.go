package main
import "github.com/gin-gonic/gin"

func main() {
    engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.String(200, "Hello CNB: Go with Gin!!")
    })
    engine.Run(":3000")
}