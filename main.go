package main

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/jonbodner-buf/test-server/gen/greet/v1"
	"github.com/jonbodner-buf/test-server/gen/greet/v1/greetv1connect"
)

func main() {
	path, handler := greetv1connect.NewGreetServiceHandler(&GreetServiceHandler{})
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

type GreetServiceHandler struct{}

func (*GreetServiceHandler) GreetPerson(ctx context.Context, req *v1.GreetPersonRequest) (*v1.GreetPersonResponse, error) {
	return &v1.GreetPersonResponse{Message: fmt.Sprintf("Hello, %s, you are %d!", req.Person.Name, req.Person.Age)}, nil
}
