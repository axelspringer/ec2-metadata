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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var (
	// ServerPort is the port the server listens on
	ServerPort int
	// ServerInterface is the interface the server listens on
	ServerInterface string
	// JSONFile is the file refence to the mock json
	JSONFile string
)

var (
	jsonObject EC2MetaData
)

// Server contains the data for the server
type Server struct {
}

// NewServer returns a new server to be run
func NewServer() *Server {
	return &Server{}
}

// Run is running the server
func (s *Server) Run(cmd *cobra.Command, args []string) {
	// test if there is a json file
	if JSONFile == "" {
		os.Exit(1)
	}

	// read in file
	file, err := ioutil.ReadFile(JSONFile)

	// err ready
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	// unmarshal json
	err = json.Unmarshal(file, &jsonObject)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
		os.Exit(1)
	}

	// add routes to the global handler
	// goji.Get("/:version/meta-data/:", )

	// middleware
	goji.Use(PlainText)
	goji.Use(ValidateVersion)

	// handle versions
	goji.Handle("/", Root)
	goji.Handle("/:version", HandleVersion)

	// seems complicated
	goji.Handle("/:version/user-data", HandleUserData)
	goji.Handle("/:version/meta-data/:endpoint", HandleMetaData)
	goji.Handle("/:version/dynamic/:endpoint", HandleDynamicData)

	// serve the mocks
	goji.Serve()
}

// Root returns the available versions
func Root(w http.ResponseWriter, r *http.Request) {
	handleJSONEncode(w, &jsonObject.Versions)
}

// HandleUserData outputs user-data
func HandleUserData(w http.ResponseWriter, r *http.Request) {
	handleJSONEncode(w, &jsonObject.UserData)
}

// HandleVersion finds the valid version
func HandleVersion(c web.C, w http.ResponseWriter, r *http.Request) {
	for _, ver := range jsonObject.Versions {
		if c.URLParams["version"] == ver {
			return
		}
	}
	// parse errror
	ErrorHandler(w, r, http.StatusNotFound)
	return
}

// HandleMetaData returns available meta data
func HandleMetaData(c web.C, w http.ResponseWriter, r *http.Request) {
	// construct endpoints
	endpoints := map[string]interface{}{
		"ami-id":               &jsonObject.MetaData.AmiID,
		"ami-launch-index":     &jsonObject.MetaData.AmiLaunchIndex,
		"ami-manifest-path":    &jsonObject.MetaData.AmiManifestPath,
		"local-ipv4":           &jsonObject.MetaData.LocalIpv4,
		"availability-zone":    &jsonObject.MetaData.AvailabilityZone,
		"hostname":             &jsonObject.MetaData.Hostname,
		"instance-id":          &jsonObject.MetaData.InstanceID,
		"instance-action":      &jsonObject.MetaData.InstanceAction,
		"instance-type":        &jsonObject.MetaData.InstanceType,
		"mac":                  &jsonObject.MetaData.Mac,
		"profile":              &jsonObject.MetaData.Profile,
		"reservation-id":       &jsonObject.MetaData.ReservationID,
		"security-credentials": &jsonObject.MetaData.SecurityCredentials,
		"security-groups":      &jsonObject.MetaData.SecurityGroups,
	}

	// parse data to endpoint
	if endpoint, ok := endpoints[c.URLParams["endpoint"]]; ok {
		// return endpoint
		handleJSONEncode(w, endpoint)
		return
	}

	// parse errror
	ErrorHandler(w, r, http.StatusNotFound)
	return
}

// HandleDynamicData returns dynamic metadata
func HandleDynamicData(c web.C, w http.ResponseWriter, r *http.Request) {
	// construct endpoints
	endpoints := map[string]interface{}{
		"identity-document": &jsonObject.DynamicData.InstanceIdentity,
	}

	// endpoint
	if endpoint, ok := endpoints[c.URLParams["endpoint"]]; ok {
		// return endpoint
		handleJSONEncode(w, endpoint)
		return
	}

	// parse errror
	ErrorHandler(w, r, http.StatusNotFound)
	return
}

// handleJsonEncode encodes objects as json and writes to the output
func handleJSONEncode(w http.ResponseWriter, obj interface{}) {
	encoder := json.NewEncoder(w)
	encoder.Encode(&obj)
}
