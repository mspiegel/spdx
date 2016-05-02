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

func TestStandardDashes(t *testing.T) {
  assert.Equal(t, "", standardDashes(""))
  assert.Equal(t, "-", standardDashes("-")) // hyphen-minus
  assert.Equal(t, "-", standardDashes("\u2010")) // hyphen
  assert.Equal(t, "-", standardDashes("\u2011")) // non-breaking hyphen
  assert.Equal(t, "-", standardDashes("\u2012")) // figure dash
  assert.Equal(t, "-", standardDashes("\u2013")) // en dash
  assert.Equal(t, "-", standardDashes("\u2014")) // em dash
  assert.Equal(t, "-", standardDashes("\u2015")) // em dash
}

func TestStandardVarietals(t *testing.T) {
  assert.Equal(t, "", standardVarietals(""))
  assert.Equal(t, "center", standardVarietals("centre"))
  assert.Equal(t, "center center", standardVarietals("centre centre"))
  assert.Equal(t, "center offense", standardVarietals("centre offence"))
}

func TestStandardize(t *testing.T) {
  assert.Equal(t, "center offense", Standardize("centre  OFFENCE"))  
}
