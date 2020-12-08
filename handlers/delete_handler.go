package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"server_bst/logger"
	"server_bst/trees"
	"strconv"
)

// DeleteHeandler ...
type DeleteHeandler struct {
	Name string
	T    trees.Tree
}

// DeleteHeandler ...
func (h *DeleteHeandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var desc string

	if r.Method != http.MethodDelete {
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
	res := h.T.DeleteTree(val)
	desc = fmt.Sprintln("Deleting a node with the value = ", valStr, ", result = ", res)
	logger.SendInfo(h.Name, desc)
	sendResult(w, res)
}
