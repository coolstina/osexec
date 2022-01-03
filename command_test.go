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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand_Simple(t *testing.T) {
	command := Command("ls", "-l")
	assert.NotNil(t, command.Args)

	actual, err := command.Output()
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
}

func TestCommand_UseSudo(t *testing.T) {
	command := Command("sudo ls", "-l")
	assert.NotNil(t, command.Args)

	actual, err := command.Output()
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
}

func TestCommand_UseSudo_AppendArgs_Multi_Args(t *testing.T) {
	command := Command("sudo ls", "-l")
	assert.NotNil(t, command.Args)
	assert.Len(t, command.Args, 3)
	assert.Equal(t, `/bin/sh -c sudo ls -l`, strings.Join(command.Args, " "))

	command.AppendArgs("|", "grep", "command.go")
	assert.Len(t, command.Args, 3)
	assert.Equal(t, `/bin/sh -c sudo ls -l | grep command.go`, strings.Join(command.Args, " "))

	actual, err := command.Output()
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
}

func TestCommand_UseSudo_AppendArgs_Raw(t *testing.T) {
	command := Command("sudo ls", "-l")
	assert.NotNil(t, command.Args)
	assert.Len(t, command.Args, 3)
	assert.Equal(t, `/bin/sh -c sudo ls -l`, strings.Join(command.Args, " "))

	command.AppendArgs("| grep command.go")
	assert.Len(t, command.Args, 3)
	assert.Equal(t, `/bin/sh -c sudo ls -l | grep command.go`, strings.Join(command.Args, " "))

	actual, err := command.Output()
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
}
