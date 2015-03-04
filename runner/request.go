package runner

import (
	"encoding/json"

	"github.com/drone/drone-cli/common"
	"github.com/samalba/dockerclient"
)

// A Request represents a build request.
type Request struct {
	Repo   *common.Repo   `json:"repo"`
	Commit *common.Commit `json:"commit"`
	Config *common.Config `json:"-"`
	Clone  *common.Clone  `json:"clone"`
	User   *common.User   `json:"user"`

	Client dockerclient.Client `json:"-"`
}

func (r *Request) Encode() string {
	out, _ := json.Marshal(r)
	return string(out)
}