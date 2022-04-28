package main

import (
	"attempt_gin/server"
)

func main() {
	r := server.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
