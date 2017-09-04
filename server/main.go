// Copyright © 2017 Sebastian Döll <sebastian@katallaxie.me>
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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/zenazn/goji"
)

// Version represents the semver of the server
const Version = "0.0.1"

var (
	ServerPort      int
	ServerInterface string
	JsonFile        string
)

type Server struct {
	bind string
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(cmd *cobra.Command, args []string) {
	if JsonFile == "" {
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(JsonFile)

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(file))

	//var json interface{}
	// println(rootCmd.Data)

	// println(cmd.Args)

	// Add routes to the global handler
	goji.Get("/", Root)

	goji.Serve()
}

// Root route (GET "/"). Print a list of greets.
func Root(w http.ResponseWriter, r *http.Request) {
	// In the real world you'd probably use a template or something.
	io.WriteString(w, "Gritter\n======\n\n")
}
