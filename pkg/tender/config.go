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
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	DiskImages []DiskImage `json:"disk_images"`
}

type DiskImage struct {
	Device   string `json:"device"`
	ImageURL string `json:"image_url"`
}

func NewConfig(configURL url.URL) (*Config, error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(configURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var config Config

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
