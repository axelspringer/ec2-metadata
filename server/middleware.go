// Copyright © 2017 Axel Springer SE
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package server

import (
	"net/http"
	"strings"
)

// PlainText sets the content-type of responses to text/plain.
func PlainText(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// ValidateVersion checks for the listest versions in request
func ValidateVersion(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		urlPart := strings.Split(r.URL.Path, "/")

		// look for version in available version, or other endpoints
		for _, ver := range jsonObject.Versions {
			if ver == urlPart[1] || r.URL.Path == "/" {
				// pass along
				h.ServeHTTP(w, r)
				return
			}
		}

		// parse errror
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	return http.HandlerFunc(fn)
}
