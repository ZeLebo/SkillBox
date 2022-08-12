package validators

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	user "user/pkg/user"
)

// Request struct to parse the request from user
type Request struct {
	TargetID int32 `json:"target_id"`
	SourceID int32 `json:"source_id"`

	Name string `json:"name"`
	Age  int    `json:"new age"`

	Friends []*user.User `json:"friends"`
}

func (req *Request) Bind(r *http.Request) error {
	content, err := ioutil.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			return
		}
	}(r.Body)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &req); err != nil {
		return errors.New("cannot parse data from JSON")
	}

	return nil
}
