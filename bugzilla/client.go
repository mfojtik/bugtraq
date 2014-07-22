package bugzilla

import (
	"log"

	"github.com/kolo/xmlrpc"
)

var Bugzillas = map[string]string{
	"redhat": "https://bugzilla.redhat.com/xmlrpc.cgi",
}

type Bugzilla struct {
	User     string
	Password string

	client *xmlrpc.Client
}

type User struct {
	Id    int    `xmlrpc:"id"`
	Token string `xmlrpc:"token"`
}

type RpcHash map[string]string

func NewClient(user, password, service string) Bugzilla {
	client := Bugzilla{User: user, Password: password}
	if err := client.Connect(service); err != nil {
		log.Fatalf("Unable to connect to %s service: %s", service, err)
	}
	return client
}

func (b *Bugzilla) Connect(bugzilla string) (err error) {
	b.client, err = xmlrpc.NewClient(Bugzillas[bugzilla], nil)
	if b.User != "" && b.Password != "" {
		user := User{}
		err = b.client.Call("User.login", RpcHash{
			"login":    b.User,
			"password": b.Password,
		}, &user)
		return
	}
	return
}

func (b *Bugzilla) SavedSearch(name string) (result BugsResult, err error) {
	err = b.client.Call("Bug.search", RpcHash{"savedsearch": name}, &result)
	return
}
