package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	task "github.com/jimweng/gogolookTask"
	"github.com/jimweng/gogolookTask/db"
	thttp "github.com/jimweng/gogolookTask/http"
)

const filePath = "/storage"

func main() {
  const defaultPort = 8080
  port := flag.Int("port", defaultPort, "listen port")
  flag.Parse()

  if err := run(*port); err != nil {
    log.Fatal(err)
  }
}

func run(port int) error {
  fs := db.NewFileSystem()
  repo := db.NewRepository(filePath, fs)
  sv := task.NewService(repo)
  srv := thttp.NewServer(sv)

  fmt.Println("server started")
  return http.ListenAndServe(fmt.Sprintf(":%d", port), srv)
}