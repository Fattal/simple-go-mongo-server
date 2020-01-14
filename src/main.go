package main

import (
  "github.com/gin-gonic/gin"
)

type User struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

func main() {
  r := gin.Default()

  r.GET("/api/get/user", getUser)
  r.POST("/api/post/user", createUser)

  r.Run(":8888")
}

func getUser(c *gin.Context) {
  c.JSON(200, gin.H{"message": "data"})
}

func createUser(c *gin.Context) {
  user := User{}
  c.BindJSON(&user)

  c.JSON(200, gin.H{ "user": user })
}
