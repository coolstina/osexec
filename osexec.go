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
)

// RunSimpleCommand execute system command.
func RunSimpleCommand(name string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", name)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return string(result), err
	}

	return string(result), nil
}
