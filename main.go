package main

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/jonbodner-buf/test-server/gen/person/v1"
	"github.com/jonbodner-buf/test-server/gen/person/v1/personv1connect"
)

func main() {
	path, handler := personv1connect.NewPersonServiceHandler(&PersonServiceHandler{})
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	h := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println(
		"serving",
		path,
	)
	h.ListenAndServe()
}

type PersonServiceHandler struct{}

func (*PersonServiceHandler) GreetPerson(ctx context.Context, req *v1.GreetPersonRequest) (*v1.GreetPersonResponse, error) {
	return &v1.GreetPersonResponse{Message: fmt.Sprintf("Hello, %s, you look great for %d!", req.Name, req.Age)}, nil
}
