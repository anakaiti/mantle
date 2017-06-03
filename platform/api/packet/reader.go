// Copyright 2017 CoreOS, Inc.
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

package packet

import (
	"io"
)

type blockingReader struct {
	done chan interface{}
}

func newBlockingReader() *blockingReader {
	return &blockingReader{
		done: make(chan interface{}),
	}
}

func (r *blockingReader) Read(b []byte) (int, error) {
	<-r.done
	return 0, io.EOF
}

func (r *blockingReader) Close() error {
	close(r.done)
	return nil
}
