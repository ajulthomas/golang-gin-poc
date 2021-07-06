package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"head": gin.H{
				"apiname":    "test",
				"apiversion": "0.1.0",
				"source":     "github.com/ajulthomas/golang-gin-poc",
			},
			"body": gin.H{
				"message":   "success",
				"userName":  "Ajul Thomas",
				"frameWork": "gin-gonic",
				"language":  "Go (aka golang)",
				"editor":    "Microsoft Visual Studio Code",
			},
		})
	})
	server.Run(":8080")
}
