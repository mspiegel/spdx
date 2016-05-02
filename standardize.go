package spdx

import (
  "regexp"
  "strings"
)

// Refer to SPDX License List Matching Guidelines, v2.0
// http://spdx.org/spdx-license-list/matching-guidelines

var multispace = regexp.MustCompile(`[[:space:]]+`)

// QUOTATION MARK, APOSTROPHE
// LEFT SINGLE QUOTATION MARK, RIGHT SINGLE QUOTATION MARK
// LEFT DOUBLE QUOTATION MARK, RIGHT DOUBLE QUOTATION MARK
var anyquote = regexp.MustCompile("[\u0022\u0027\u2018\u2019\u201c\u201d]")

// 3.1.1 Guideline: All whitespace should be treated as a single blank space.
func removeMultispace(input string) string {
  return multispace.ReplaceAllString(input, " ")
}

// 4.1.1 Guideline: All upper case and lower case letters should be
// treated as lower case letters.
func removeUppercase(input string) string {
  return strings.ToLower(input)
}

// 5.1.3 Guideline: Quotes. Any variation of quotations (single, double,
// curly, etc.) should be considered equivalent.
func standardQuotes(input string) string {
  return anyquote.ReplaceAllString(input, `'`)
}
