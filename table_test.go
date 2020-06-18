// Copyright 2019 Yunion
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

package sqlchemy

import (
	"reflect"
	"testing"
)

func TestTableSpecDataGet(t *testing.T) {
	type S struct {
		Id   string
		Zone string
	}
	ts := NewTableSpecFromStruct(S{}, "s")

	for i := 0; i < 4; i++ {
		el := ts.DataPoolGet().(*S)
		if el.Id != "" || el.Zone != "" {
			t.Fatalf("get none zero el: %#v", el)
		}
		el.Id = "id"
		el.Zone = "zone"
		ts.DataPoolPut(el)
	}
}

func BenchmarkTableSpecDataGet(b *testing.B) {
	type S struct {
		Id   string
		Name string

		Project string
		Domain  string

		Zone string
	}
	const sliceLen = 20

	b.Run("pool", func(b *testing.B) {
		ts := NewTableSpecFromStruct(S{}, "s")
		li := make([]interface{}, sliceLen)
		for i := 0; i < b.N; i++ {
			for i := range li {
				li[i] = ts.DataPoolGet().(*S)
			}
			for i := range li {
				ts.DataPoolPut(li[i])
			}
		}
	})

	b.Run("reflect", func(b *testing.B) {
		ts := NewTableSpecFromStruct(S{}, "s")
		li := make([]interface{}, sliceLen)
		for i := 0; i < b.N; i++ {
			for i := range li {
				li[i] = reflect.New(ts.DataType()).Interface()
			}
		}
	})
}
