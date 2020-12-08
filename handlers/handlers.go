package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// ErrMsg ...
type ErrMsg struct {
	Error string
}

// ResMsg ...
type ResMsg struct {
	Result bool
}

// SearchData ...
type SearchData struct {
	Val int
}

func sendError(w http.ResponseWriter, desc string) {

	w.WriteHeader(http.StatusInternalServerError)

	em := &ErrMsg{
		Error: desc,
	}

	b, err := json.Marshal(em)
	if err != nil {
		w.Write([]byte("Error generating the server response"))
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	w.Write(b)
}

func sendResult(w http.ResponseWriter, res bool) {

	w.WriteHeader(http.StatusOK)

	r := &ResMsg{
		Result: res,
	}

	b, err := json.Marshal(r)
	if err != nil {
		w.Write([]byte("Error generating the server response"))
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	w.Write(b)
}
