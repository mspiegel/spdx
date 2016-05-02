package spdx

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestRemoveMultispace(t *testing.T) {
  assert.Equal(t, "", removeMultispace(""))
  assert.Equal(t, "a", removeMultispace("a"))
  assert.Equal(t, " a ", removeMultispace(" a  "))
  assert.Equal(t, "a b", removeMultispace("a     b"))
}

func TestRemoveUppercase(t *testing.T) {
  assert.Equal(t, "", removeUppercase(""))
  assert.Equal(t, "abc", removeUppercase("ABC"))
}

func TestStandardQuotes(t *testing.T) {
  assert.Equal(t, "", standardQuotes(""))
  assert.Equal(t, "'", standardQuotes("'"))
  assert.Equal(t, "'", standardQuotes("\""))
  assert.Equal(t, "'", standardQuotes("‘"))
  assert.Equal(t, "'", standardQuotes("’"))
  assert.Equal(t, "'", standardQuotes("“"))
  assert.Equal(t, "'", standardQuotes("”"))
}
