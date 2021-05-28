package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type dataInfo struct {
	Msg1 string `json:"msg1"`
	Msg2 string `json:"msg2"`
	Msg3 string `json:"msg3"`
	Msg4 string `json:"msg4"`
	Msg5 string `json:"msg5"`
	Msg6 string `json:"msg6"`
	Msg7 string `json:"msg7"`
}

func greetings(c *gin.Context) {
	var (result gin.H)
	dataStruct := dataInfo{
		Msg1: "Hey Angelia, I love you",
		Msg2: "I will love you like i love my mom",
		Msg3: "It means, that i won't hurt you forever",
		Msg4: "Because u are the second woman after my mom",
		Msg5: "One more, i won't hurt you",
		Msg6: "Forever...",
		Msg7: "Made with ❣ - Herly",
	}
	result = gin.H{
		"error" : nil,
		"data": dataStruct,
	}

	c.JSON(http.StatusOK, result)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	//not found
	router.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found", "data":nil})
	})
	router.GET("/api/greetings", greetings)
	router.Run(":5000")
}