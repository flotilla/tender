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

package tender

import (
	"io"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

func ApplyConfig(configURL url.URL) error {
	log.WithFields(log.Fields{
		"configURL": &configURL,
	}).Info("Starting to execute config")

	config, err := NewConfig(configURL)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"config": config,
	}).Info("Running with config")

	for _, di := range config.DiskImages {
		err = writeImage(di.Device, di.ImageURL)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeImage(devicePath, imageURL string) error {

	out, err := os.OpenFile(devicePath, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
