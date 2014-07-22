package bugzilla

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type RedHat struct {
	User     string
	Password string
	ListId   string
}

func (s RedHat) LastUpdate() time.Time {
	return time.Now()
}

func (s RedHat) GetListJSON() (r string, err error) {

	client := NewClient(s.User, s.Password, "https://bugzilla.redhat.com/xmlrpc.cgi")

	result, err := client.GetList(s.ListId)
	if err != nil {
		return
	}

	jsonResult, err := json.Marshal(&result)
	return string(jsonResult), err
}

func (s RedHat) String() string {
	return fmt.Sprintf("Red Hat Bugzilla - %s", s.ListId)
}

func (s RedHat) Id() string {
	return url.QueryEscape(s.ListId)
}
