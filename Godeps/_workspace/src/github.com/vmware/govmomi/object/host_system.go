/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import (
	"fmt"
	"net"

	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/vim25"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/vim25/mo"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/vim25/types"
	"github.com/apcera/libretto/Godeps/_workspace/src/golang.org/x/net/context"
)

type HostSystem struct {
	Common

	InventoryPath string
}

func (h HostSystem) String() string {
	if h.InventoryPath == "" {
		return h.Common.String()
	}
	return fmt.Sprintf("%v @ %v", h.Common, h.InventoryPath)
}

func NewHostSystem(c *vim25.Client, ref types.ManagedObjectReference) *HostSystem {
	return &HostSystem{
		Common: NewCommon(c, ref),
	}
}

func (h HostSystem) Name(ctx context.Context) (string, error) {
	var mh mo.HostSystem

	err := h.Properties(ctx, h.Reference(), []string{"name"}, &mh)
	if err != nil {
		return "", err
	}

	return mh.Name, nil
}

func (h HostSystem) ConfigManager() *HostConfigManager {
	return NewHostConfigManager(h.c, h.Reference())
}

func (h HostSystem) ResourcePool(ctx context.Context) (*ResourcePool, error) {
	var mh mo.HostSystem

	err := h.Properties(ctx, h.Reference(), []string{"parent"}, &mh)
	if err != nil {
		return nil, err
	}

	var mcr *mo.ComputeResource
	var parent interface{}

	switch mh.Parent.Type {
	case "ComputeResource":
		mcr = new(mo.ComputeResource)
		parent = mcr
	case "ClusterComputeResource":
		mcc := new(mo.ClusterComputeResource)
		mcr = &mcc.ComputeResource
		parent = mcc
	default:
		return nil, fmt.Errorf("unknown host parent type: %s", mh.Parent.Type)
	}

	err = h.Properties(ctx, *mh.Parent, []string{"resourcePool"}, parent)
	if err != nil {
		return nil, err
	}

	pool := NewResourcePool(h.c, *mcr.ResourcePool)
	return pool, nil
}

func (h HostSystem) ManagementIPs(ctx context.Context) ([]net.IP, error) {
	var mh mo.HostSystem

	err := h.Properties(ctx, h.Reference(), []string{"config.virtualNicManagerInfo.netConfig"}, &mh)
	if err != nil {
		return nil, err
	}

	var ips []net.IP
	for _, nc := range mh.Config.VirtualNicManagerInfo.NetConfig {
		if nc.NicType == "management" && len(nc.CandidateVnic) > 0 {
			ip := net.ParseIP(nc.CandidateVnic[0].Spec.Ip.IpAddress)
			if ip != nil {
				ips = append(ips, ip)
			}
		}
	}

	return ips, nil
}
