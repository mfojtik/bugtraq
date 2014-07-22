package bugzilla

import (
	"encoding/json"
	"time"

	"github.com/mfojtik/bugtraq/scheduler"
	"github.com/mfojtik/bugtraq/services"
)

type BugzillaService struct {
	client *Bugzilla
}

func (s BugzillaService) List() (r string, err error) {
	var result BugsResult

	if err = s.client.Connect("redhat"); err != nil {
		return
	}

	if result, err = s.client.SavedSearch("openshift-blockers"); err != nil {
		return
	}

	jsonResult, err := json.Marshal(&result)
	return string(jsonResult), err
}

func RedHat(user, password string) BugzillaService {
	c := NewClient(user, password, "redhat")
	return BugzillaService{client: &c}
}

func ScheduledRedHat(user, password string, interval time.Duration) services.Service {
	redhat := RedHat(user, password)
	s := scheduler.ServiceScheduler{ServiceHandler: &redhat}
	s.Schedule(interval)
	return s
}
