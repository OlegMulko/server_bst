package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server_bst/logger"
	"server_bst/trees"
	"strconv"
)

// InsertHeandler ...
type InsertHeandler struct {
	Name string
	T    trees.Tree
}

// ServeHTTP ...
func (h *InsertHeandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var desc string

	if r.Method != http.MethodPost {
		desc = fmt.Sprintln("Invalid method, request metod - ", r.Method)
		logger.SendError(h.Name, desc)
		sendError(w, desc)
		return
	}
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		desc = "Missing content-type header"
		logger.SendError(h.Name, desc)
		sendError(w, desc)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.SendError(h.Name, err.Error())
		sendError(w, err.Error())
		return
	}
	var sd SearchData
	err = json.Unmarshal(b, &sd)
	if err != nil {
		logger.SendError(h.Name, err.Error())
		sendError(w, err.Error())
		return
	}
	h.T.InsertTree(sd.Val)
	desc = fmt.Sprintln("Adding a new node with the value = ", strconv.Itoa(sd.Val), " , result = ", true)
	logger.SendInfo(h.Name, desc)
	sendResult(w, true)
}
