package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/mfojtik/bugtraq/services"
)

type ListCache struct {
	cachedJson string
	lastUpdate time.Time
}

func (l ListCache) GetId() string         { return "list" }
func (l ListCache) String() string        { return l.cachedJson }
func (l ListCache) LastUpdate() time.Time { return l.lastUpdate }

type ServiceScheduler struct {
	ServiceHandler services.Service
	cache          *Cache
}

func (s ServiceScheduler) List() (string, error) {
	if cache := s.cache.Get("list"); cache != nil {
		listCache := cache.(ListCache)
		log.Printf("Serving cache from %s", listCache.LastUpdate())
		return listCache.String(), nil
	} else {
		return "", fmt.Errorf("Empty cache")
	}
}

func (s *ServiceScheduler) Schedule(interval time.Duration) chan struct{} {
	s.cache = NewCache()
	t := time.NewTicker(interval)
	c := make(chan struct{})
	log.Println("Scheduler started.")
	go func() {
		s.Execute()
		for {
			select {
			case <-t.C:
				s.Execute()
			case <-c:
				break
			}
		}
	}()
	return c
}

func (s *ServiceScheduler) Execute() {
	if result, err := s.ServiceHandler.List(); err != nil {
		log.Printf("WARNING: %s\n", err)
	} else {
		s.cache.Add(ListCache{cachedJson: result, lastUpdate: time.Now()})
		log.Println("INFO: Synchronized successfully")
	}
}
