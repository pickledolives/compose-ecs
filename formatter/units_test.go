/*
   Copyright 2020 Docker, Inc.

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

package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	b  = 1
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
)

func TestMemBytes(t *testing.T) {
	var m MemBytes
	assert.Nil(t, m.Set("42"))
	assert.Equal(t, int64(42), m.Value())
	assert.Equal(t, "42B", m.String())

	assert.Nil(t, m.Set("1b"))
	assert.Equal(t, int64(1), m.Value())
	assert.Equal(t, "1B", m.String())

	assert.Nil(t, m.Set("1k"))
	assert.Equal(t, int64(kb), m.Value())
	assert.Equal(t, "1KiB", m.String())

	assert.Nil(t, m.Set("1m"))
	assert.Equal(t, int64(mb), m.Value())
	assert.Equal(t, "1MiB", m.String())

	assert.Nil(t, m.Set("1g"))
	assert.Equal(t, int64(gb), m.Value())
	assert.Equal(t, "1GiB", m.String())

	assert.Nil(t, m.Set("42g"))
	assert.Equal(t, int64(42*gb), m.Value())
	assert.Equal(t, "42GiB", m.String())

	assert.Error(t, m.Set("###"))
}