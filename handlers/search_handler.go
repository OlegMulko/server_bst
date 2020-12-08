package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"server_bst/logger"
	"server_bst/trees"
	"strconv"
)

// SearchHeandler ...
type SearchHeandler struct {
	Name string
	T    trees.Tree
}

// ServeHTTP ...
func (h *SearchHeandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var desc string

	if r.Method != http.MethodGet {
		desc = fmt.Sprintln("Invalid method, request metod - ", r.Method)
		logger.SendError(h.Name, desc)
		sendError(w, desc)
		return
	}
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		logger.SendError(h.Name, err.Error())
		sendError(w, err.Error())
		return
	}
	valStr := values.Get("val")
	if valStr == "" {
		desc = "The val parameter is missing"
		logger.SendError(h.Name, desc)
		sendError(w, desc)
		return
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		logger.SendError(h.Name, err.Error())
		sendError(w, err.Error())
		return
	}
	res := h.T.SearchTree(val)
	desc = fmt.Sprintln("Search for a node with a value = ", valStr, ", result = ", res)
	logger.SendInfo(h.Name, desc)
	sendResult(w, res)
}
