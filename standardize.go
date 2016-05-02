package spdx

import (
  "regexp"
  "strings"
)

// Refer to SPDX License List Matching Guidelines, v2.0
// http://spdx.org/spdx-license-list/matching-guidelines

var multispace = regexp.MustCompile(`[[:space:]]+`)

// quotation mark
// apostrophe
// left single quotation mark
// right single quotation mark
// left double quotation mark
// right double quotation mark
var anyquote = regexp.MustCompile("[\u0022\u0027\u2018\u2019\u201c\u201d]")

// hyphen-minus
// hyphen
// non-breaking hyphen
// figure dash
// en dash
// em dash
// horizontal bar
var anydash = regexp.MustCompile("[\u002d\u2010-\u2015]")

var varietals = map[string] string {
"acknowledgment" : "acknowledgement",
"analogue" : "analog",
"analyse" : "analyze",
"artefact" : "artifact",
"authorisation" : "authorization",
"authorised" : "authorized",
"calibre" : "caliber",
"cancelled" : "canceled",
"capitalisations" : "capitalizations",
"catalogue" : "catalog",
"categorise" : "categorize",
"centre" : "center",
"emphasised" : "emphasized",
"favour" : "favor",
"favourite" : "favorite",
"fulfil" : "fulfill",
"fulfilment" : "fulfillment",
"initialise" : "initialize",
"judgment" : "judgement",
"labelling" : "labeling",
"labour" : "labor",
"licence" : "license",
"maximise" : "maximize",
"modelled" : "modeled",
"modelling" : "modeling",
"offence" : "offense",
"optimise" : "optimize",
"organisation" : "organization",
"organise" : "organize",
"practise" : "practice",
"programme" : "program",
"realise" : "realize",
"recognise" : "recognize",
"signalling" : "signaling",
"sub-license" : "sublicense",
"sub license" : "sublicense",
"utilisation" : "utilization",
"whilst" : "while",
"wilful" : "wilfull",
"non-commercial" : "noncommercial",
"per cent" : "percent",
"copyright owner" : "copyright holder",
}

// 3.1.1 Guideline: All whitespace should be treated as a single blank space.
func removeMultispace(s string) string {
  return multispace.ReplaceAllString(s, " ")
}

// 4.1.1 Guideline: All upper case and lower case letters should be
// treated as lower case letters.
func removeUppercase(s string) string {
  return strings.ToLower(s)
}

// 5.1.3 Guideline: Quotes. Any variation of quotations (single, double,
// curly, etc.) should be considered equivalent.
func standardQuotes(s string) string {
  return anyquote.ReplaceAllString(s, `'`)
}

// 5.1.2 Guideline: Hyphens, Dashes.  Any hyphen, dash, en dash, em dash,
// or other variation should be considered equivalent.
func standardDashes(s string) string {
  return anydash.ReplaceAllString(s, `-`)
}

// 8.1.1 Guideline: The words in the following columns are considered
// equivalent and interchangeable.
// Precondition: input is all lowercase characters
func standardVarietals(s string) string {
  for k, v := range varietals {
    s = strings.Replace(s, k, v, -1)
  }
  return s
}

var Standardize = compose(removeMultispace, removeUppercase,
  standardQuotes, standardDashes,standardVarietals)
