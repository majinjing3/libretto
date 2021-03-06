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

package fields

import (
	"flag"
	"fmt"

	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/govc/cli"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/govc/flags"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/object"
	"github.com/apcera/libretto/Godeps/_workspace/src/golang.org/x/net/context"
)

type add struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("fields.add", &add{})
}

func (cmd *add) Register(f *flag.FlagSet) {}

func (cmd *add) Process() error { return nil }

func (cmd *add) Usage() string {
	return "NAME"
}

func (cmd *add) Run(f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	ctx := context.TODO()

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetCustomFieldsManager(c)
	if err != nil {
		return err
	}

	name := f.Arg(0)

	def, err := m.Add(ctx, name, "", nil, nil)
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", def.Key)

	return nil
}
