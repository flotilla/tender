// Copyright © 2018 Craig Tracey <craigtracey@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/flotilla/tender/pkg/tender"
	log "github.com/sirupsen/logrus"
)

const kernelCmdLineOpt = "flotilla-config"

func parseKernelCmdline() (*url.URL, error) {
	cmdLine, err := ioutil.ReadFile("/tmp/cmdline")
	if err != nil {
		return nil, err
	}

	paramStrs := strings.Fields(string(cmdLine))
	for _, param := range paramStrs {
		if strings.HasPrefix(param, kernelCmdLineOpt) {
			val := strings.Split(param, "=")
			if len(val) != 2 {
				return nil, fmt.Errorf("Invalid format for %s", kernelCmdLineOpt)
			}
			url, err := url.Parse(val[1])
			if err != nil {
				return nil, fmt.Errorf("Could not parse url %s", val[1])
			}
			return url, nil
		}
	}
	return nil, fmt.Errorf("No config option '%s' found on command line", kernelCmdLineOpt)
}

func main() {

	log.Info("tender provisioning agent started.")
	confURL, err := parseKernelCmdline()
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(log.Fields{
		"configURL": confURL,
	}).Info("Discovered config URL")

	err = tender.ApplyConfig(*confURL)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to apply configuration")
	}
	log.Info("Congratulations! Configuration successfully applied.")
}
