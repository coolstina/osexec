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
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSuiteSeries(t *testing.T) {
	suite.Run(t, &SuiteSeries{})
}

type SuiteSeries struct {
	suite.Suite
}

func (suite *SuiteSeries) Test_RunSimpleCommand_NotExistsCommand() {
	actual, err := RunSimpleCommand("ll -l")
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), actual)
	assert.Equal(suite.T(), "/bin/sh: ll: command not found\n", actual)
}

func (suite *SuiteSeries) Test_RunSimpleCommand() {
	actual, err := RunSimpleCommand("ls -l")
	assert.NoError(suite.T(), err)
	split := strings.Split(actual, "\n")
	assert.GreaterOrEqual(suite.T(), len(split), 4)
}

func (suite *SuiteSeries) Test_RunSimpleCommand_UNAME() {
	actual, err := RunSimpleCommand("uname -a")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), actual)
}

func (suite *SuiteSeries) Test_RunSimpleCommand_SUDO() {
	actual, err := RunSimpleCommand("sudo chmod 0755 osexec.go")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	fmt.Println(actual)
}
