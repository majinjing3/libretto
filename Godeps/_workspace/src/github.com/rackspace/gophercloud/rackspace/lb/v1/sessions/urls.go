package sessions

import (
	"strconv"

	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"
)

const (
	path   = "loadbalancers"
	spPath = "sessionpersistence"
)

func rootURL(c *gophercloud.ServiceClient, id int) string {
	return c.ServiceURL(path, strconv.Itoa(id), spPath)
}
