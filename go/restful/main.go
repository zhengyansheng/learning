// https://github.com/emicklei/go-restful/blob/master/examples/restful-multi-containers.go
// GET http://localhost:8080/hello
// GET http://localhost:8081/hello
package main

import (
	"io"
	"log"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

/*
	container -> webserver -> route
*/

func main() {
	// 创建一个默认的container
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws) //添加到默认的container中
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// 创建一个新的container
	container2 := restful.NewContainer()
	ws2 := new(restful.WebService)
	ws2.Route(ws2.GET("/hello").To(hello2))
	container2.Add(ws2) //添加到container2中
	server := &http.Server{Addr: ":8081", Handler: container2}
	log.Fatal(server.ListenAndServe())
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "default world")
}

func hello2(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "second world")
}
