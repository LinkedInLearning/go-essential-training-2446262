/* Key/Value database

$ curl -d'hello' http://localhost:8080/k1
$ curl http://localhost:8080/k1
hello
$ curl -i http://localhost:8080/k2
404 not found

Limit value size to 1k
*/
package main

import (
	"net/http"
)

type Server struct {
	db DB
}

// POST /key Store request body as value
// GET /<key> Send back value, or 404 if key not found
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func main() {
	// TODO: Routing & start server
}
