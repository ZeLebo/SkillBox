package Service

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"user/pkg/Error"
)

// Request struct to parse the request from user
type Request struct {
	TargetID int32 `json:"target_id"`
	SourceID int32 `json:"source_id"`
	Age      int   `json:"new age"`
}

func (req *Request) Bind(w http.ResponseWriter, r *http.Request) error {
	content, err := ioutil.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			Error.HTTPErrorHandle(w, Error.HTTPErrorHandler{
				ErrorCode:   http.StatusInternalServerError,
				Description: "Error while closing file after reading note",
			})
		}
	}(r.Body)

	if err != nil {
		Error.HTTPErrorHandle(w, Error.HTTPErrorHandler{
			ErrorCode:   http.StatusInternalServerError,
			Description: "Cannot read the data from request",
		})
		return errors.New("cannot read the data from request")
	}

	if err = json.Unmarshal(content, &req); err != nil {
		Error.HTTPErrorHandle(w, Error.HTTPErrorHandler{
			ErrorCode:   http.StatusInternalServerError,
			Description: "Cannot parse data from JSON",
		})
		return errors.New("cannot parse data from JSON")
	}

	return nil
}
