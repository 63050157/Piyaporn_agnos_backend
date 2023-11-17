package main

import (
  "github.com/gin-gonic/gin"
  "Piyaporn_agnos_backend/database"
	"Piyaporn_agnos_backend/handler"
)

func main() {
  db, err := database.InitDB()
  if err != nil {
    panic(err)
  }
  defer db.Close()

  r := gin.Default()

  apiGroup := r.Group("/api")
  apiGroup.POST("/strong_password_steps", handler.AddDataHandler(db))

  r.Run(":8080")
}
