// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	srv := gin.New()
// 	srv.Use(gin.Logger(), gin.Recovery())

// 	srv.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message":"Hello Mars!",
// 		})
// 	})
// 	srv.Run(":8080")
// }

package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second * 3)
	fmt.Println("done")
    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
	
}