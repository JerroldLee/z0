// Copyright 2018 The zipper Authors
// This file is part of the z0 library.
//
// The z0 library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The z0 library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the z0 library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"io/ioutil"
	"path"
	"strings"
)

// GitCommit  Git SHA1 commit hash of the release (set via linker flags)
var GitCommit = func() string {
	head := readGitFile("HEAD")
	if splits := strings.Split(head, " "); len(splits) == 2 {
		head = splits[1]
		return readGitFile(head)
	}
	return ""
}

// readGitFile returns content of file in .git directory.
func readGitFile(file string) string {
	content, err := ioutil.ReadFile(path.Join(".git", file))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(content))
}
