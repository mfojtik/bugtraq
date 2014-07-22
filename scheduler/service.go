package scheduler

import (
	"time"

	"github.com/mfojtik/bugtraq/services"
)

type ServiceScheduler struct {
	Timeout time.Duration

	serviceScheduler *Scheduler
}

func (s *ServiceScheduler) GetListJSON() (string, error) {
	return s.serviceScheduler.CurrentCachedList()
}

func (s *ServiceScheduler) LastUpdate() time.Time {
	return s.serviceScheduler.LastUpdate()
}

func (s *ServiceScheduler) Id() string {
	return s.serviceScheduler.service.Id()
}

func (s *ServiceScheduler) Schedule(p services.Service) {
	s.serviceScheduler = &Scheduler{service: p, cache: NewCache()}
	s.serviceScheduler.Schedule(s.Timeout)
}

func NewServiceScheduler(t time.Duration) ServiceScheduler {
	return ServiceScheduler{Timeout: t}
}
