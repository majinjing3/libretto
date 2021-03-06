package databases

import (
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"
	os "github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud/openstack/db/v1/databases"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud/pagination"
)

func Create(client *gophercloud.ServiceClient, instanceID string, opts os.CreateOptsBuilder) os.CreateResult {
	return os.Create(client, instanceID, opts)
}

func List(client *gophercloud.ServiceClient, instanceID string) pagination.Pager {
	return os.List(client, instanceID)
}

func Delete(client *gophercloud.ServiceClient, instanceID, dbName string) os.DeleteResult {
	return os.Delete(client, instanceID, dbName)
}
