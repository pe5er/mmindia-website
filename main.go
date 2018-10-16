// main.go

package main

import (
  "io"
  "os"
  "github.com/gin-gonic/gin"
)

func main() {

  gin.SetMode(gin.ReleaseMode)
  gin.DisableConsoleColor()
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f)


  r := routes()

  r.LoadHTMLGlob("templates/*")

  r.Run(":8081")

}
