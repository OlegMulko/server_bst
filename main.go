package main

import (
	"fmt"
	"net/http"
	"server_bst/handlers"
	"server_bst/trees"
	"sync"
)

func main() {

	tree := &trees.TreeBst{
		Mux: &sync.RWMutex{},
	}

	SearchHeandler := &handlers.SearchHeandler{
		Name: "HttpHeandlerSearch",
		T:    tree,
	}
	http.Handle("/search", SearchHeandler)

	InsertHeandler := &handlers.InsertHeandler{
		Name: "HttpHeandlerSearch",
		T:    tree,
	}
	http.Handle("/insert", InsertHeandler)

	DeleteHeandler := &handlers.DeleteHeandler{
		Name: "HttpHeandlerSearch",
		T:    tree,
	}
	http.Handle("/delete", DeleteHeandler)

	fmt.Println("Http server bst start on port 8080")
	http.ListenAndServe(":8080", nil)
}
