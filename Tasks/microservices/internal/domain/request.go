package domain

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"user/internal/user"
)

// Request struct to parse the request from user
type Request struct {
	TargetID int          `json:"target_id"`
	SourceID int          `json:"source_id"`
	Name     string       `json:"name"`
	Age      int          `json:"age"`
	Friends  []*user.User `json:"friends"`
}

func (req *Request) Bind(r *http.Request) error {
	//goland:noinspection ALL
	buff, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buff, req)
	if err != nil {
		return err
	}

	return nil
}

func (req *Request) BindRequestParams(r *http.Request) error {
	vars := mux.Vars(r)
	if vars["id"] != "" {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}
		req.TargetID = id
	}
	return nil
}
