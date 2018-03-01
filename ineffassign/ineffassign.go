// Copyright 2016 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ineffassign

import (
	"io"

	"github.com/palantir/okgo/checker"
	"github.com/palantir/okgo/okgo"
)

const (
	TypeName okgo.CheckerType     = "ineffassign"
	Priority okgo.CheckerPriority = 0
)

func Creator() checker.Creator {
	return checker.NewCreator(
		TypeName,
		Priority,
		func(cfgYML []byte) (okgo.Checker, error) {
			return &ineffassignCheck{}, nil
		},
	)
}

type ineffassignCheck struct{}

func (c *ineffassignCheck) Type() (okgo.CheckerType, error) {
	return TypeName, nil
}

func (c *ineffassignCheck) Priority() (okgo.CheckerPriority, error) {
	return Priority, nil
}

func (c *ineffassignCheck) Check(pkgPaths []string, projectDir string, stdout io.Writer) {
	cmd, wd := checker.AmalgomatedCheckCmd(string(TypeName), append([]string{"-n"}, pkgPaths...), stdout)
	if cmd == nil {
		return
	}
	checker.RunCommandAndStreamOutput(cmd, func(line string) okgo.Issue {
		return okgo.NewIssueFromLine(line, wd)
	}, stdout)
}

func (c *ineffassignCheck) RunCheckCmd(args []string, stdout io.Writer) {
	checker.AmalgomatedRunRawCheck(string(TypeName), args, stdout)
}
