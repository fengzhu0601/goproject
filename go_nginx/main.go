package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func run(idx int) {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg":          fmt.Sprintf("hello from host %v", idx),
			"handler host": idx,
		})
	})

	r.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg":          fmt.Sprintf("hello from host %v", idx),
			"handler host": idx,
			"additional":   "/index",
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":808%d", idx),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("The service cannot start %v", err)
		return
	}
	wg.Done()
}

func main() {
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go run(i)
	}
	wg.Wait()
}
