// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package osexec

import (
	"os/exec"
	"strings"
)

// Cmd is internal package exec.Cmd wrapper.
type Cmd struct {
	*exec.Cmd
}

// AppendArgs appends args.
func (cmd *Cmd) AppendArgs(arg ...string) {
	if len(cmd.Args) != 2 {
		single := cmd.Args[len(cmd.Args)-1]
		args := strings.Split(single, " ")
		args = append(args, arg...)

		cmd.Args[len(cmd.Args)-1] = strings.Join(args, " ")
	}
}

// Command initialize internal package exec.Cmd wrapper command instance.
func Command(name string, arg ...string) *Cmd {
	all := strings.Join(append([]string{name}, arg...), " ")
	args := append([]string{"-c"}, all)

	return &Cmd{exec.Command("/bin/sh", args...)}
}
