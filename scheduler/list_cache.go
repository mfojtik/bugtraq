package scheduler

import "time"

type ListCache struct {
	cachedJson string
	lastUpdate time.Time
}

func (l ListCache) GetId() string         { return "list" }
func (l ListCache) String() string        { return l.cachedJson }
func (l ListCache) LastUpdate() time.Time { return l.lastUpdate }
