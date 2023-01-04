package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("launching server at port 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/**
// 1.构建镜像
sudo docker build . -t myimage

// 2.查看镜像
sudo docker image ls

// 3.运行镜像
sudo docker run -d -p 8080:8080 myimage

// 4. 查看服务启动
fengzhu@:~/deploy-on-docker$ sudo docker container ls
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                                       NAMES
deceb282986f   myimage   "/bin/sh -c /app/main"   4 seconds ago   Up 2 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   naughty_goldberg
*/
