// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.
package spdx

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
  var foo strFunc = func(s string) string {
  	return s + "foo "
  }
  var bar strFunc = func(s string) string {
  	return s + "bar "
  }
  var baz strFunc = func(s string) string {
  	return s + "baz "
  }
  assert.Equal(t, "foo bar baz ", compose(foo, bar, baz)(""))
}
