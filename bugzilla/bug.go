package bugzilla

import "time"

// For more information about the fields look here:
// http://www.bugzilla.org/docs/4.4/en/html/api/Bugzilla/WebService/Bug.html#get

type Bug struct {
	Id             int       `xmlrpc:"id"`
	Url            string    `xmlrpc:"url"`
	Summary        string    `xmlrpc:"summary"`
	Status         string    `xmlrpc:"status"`
	CreationTime   time.Time `xmlrpc:"creation_time"`
	AssignedTo     string    `xmlrpc:"assigned_to"`
	Component      []string  `xmlrpc:"component"`
	Product        string    `xmlrpc:"product"`
	Platform       string    `xmlrpc:"platform"`
	Priority       string    `xmlrpc:"priority"`
	Severity       string    `xmlrpc:"severity"`
	Creator        string    `xmlrpc:"creator"`
	IsOpen         bool      `xmlrpc:"is_open"`
	Keywords       []string  `xmlrpc:"keywords"`
	LastChangeTime time.Time `xmlrpc:"last_change_time"`
	Resolution     string    `xmlrpc:"resolution"`
	DependsOn      []int     `xmlrpc:"depends_on"`
	Blocks         []int     `xmlrpc:"blocks"`
	DupeOf         int       `xmlrpc:"dupe_of"`
}

type BugsResult struct {
	Bugs []Bug `xmlrpc:"bugs"`
}
