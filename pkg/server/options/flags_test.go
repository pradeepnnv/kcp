/*
Copyright 2022 The KCP Authors.

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

package options

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
)

func TestNamedFlagSetOrder(t *testing.T) {
	fss := NewOptions().Flags()
	var names []string
	for name, fs := range fss.FlagSets {
		if !fs.HasFlags() {
			continue
		}
		fmt.Printf("%q,\n", name)
		names = append(names, name)
	}

	require.Subset(t, names, namedFlagSetOrder, "namedFlagSetOrder has extra entries")
	require.Subset(t, namedFlagSetOrder, names, "namedFlagSetOrder in incomplete")
}

func TestAllowedFlagList(t *testing.T) {
	o := NewOptions()
	fss := o.rawFlags()

	missing := map[string][]*pflag.Flag{}

	for section, fs := range fss.FlagSets {
		fs.VisitAll(func(f *pflag.Flag) {
			if !allowedFlags.Has(f.Name) && !disallowedFlags.Has(f.Name) {
				t.Errorf("--%s in section \"%s flags\" is neither allowed nor disallowed", f.Name, section)
				missing[section] = append(missing[section], f)
			}
		})
	}

	// print out for easy copy/paste during rebase
	if len(missing) > 0 {
		t.Log("Flags missing from allowlist and denylist:")
		for section, flags := range missing {
			fmt.Printf("\n\t// %s flags\n", section)
			for _, flag := range flags {
				fmt.Printf("\t\"%s\", // %s\n", flag.Name, strings.Split(flag.Usage, "\n")[0])
			}
		}
	}
}

func TestAllowedFlagListCleanup(t *testing.T) {
	o := NewOptions()
	fss := o.rawFlags()

	allFlags := map[string]*pflag.Flag{}
	for _, fs := range fss.FlagSets {
		fs.VisitAll(func(f *pflag.Flag) {
			allFlags[f.Name] = f
		})
	}

	for _, flag := range allowedFlags.List() {
		if _, ok := allFlags[flag]; !ok {
			t.Errorf("flag --%s is allowed but not in any flag set", flag)
		}
	}
	for _, flag := range disallowedFlags.List() {
		if _, ok := allFlags[flag]; !ok {
			t.Errorf("flag --%s is allowed but not in any flag set", flag)
		}
	}
}
