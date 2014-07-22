package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/mfojtik/bugtraq/services"
)

type Scheduler struct {
	service services.Service
	cache   *Cache
}

func (s *Scheduler) CurrentCachedList() (string, error) {
	if cache := s.cache.Get("list"); cache != nil {
		listCache := cache.(ListCache)
		log.Printf("INFO: Cache hit (%s)", listCache.LastUpdate())
		return listCache.String(), nil
	} else {
		return "", fmt.Errorf("The cache is not populated yet.")
	}
}

func (s *Scheduler) LastUpdate() time.Time {
	if cache := s.cache.Get("list"); cache != nil {
		listCache := cache.(ListCache)
		return listCache.LastUpdate()
	} else {
		return time.Now()
	}
}

func (s *Scheduler) Schedule(interval time.Duration) chan struct{} {
	t := time.NewTicker(interval)
	c := make(chan struct{})
	log.Println("INFO: Starting synchronization scheduler")
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

func (s *Scheduler) Execute() {
	log.Printf("INFO: Fetching updates [%s]...", s.service)
	if result, err := s.service.GetListJSON(); err != nil {
		log.Printf("WARNING: %s\n", err)
	} else {
		s.cache.Add(ListCache{cachedJson: result, lastUpdate: time.Now()})
	}
}
