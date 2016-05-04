// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.
package spdx

type strFunc func(s string) string

func compose2(a strFunc, b strFunc) strFunc {
	return func(s string) string {
		return b(a(s))
	}
}

func compose(fns ...strFunc) strFunc {
	return func(s string) string {
		var res strFunc
		res = fns[0]
		for i := 1; i < len(fns); i++ {
			res = compose2(res, fns[i])
		}
		return res(s)
	}
}
