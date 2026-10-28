// Copyright 2026 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"net/url"
	"strings"
)

// ForgeConfig holds the generic configuration of a forge instance.
type ForgeConfig struct {
	// Host is the host the forge is reachable at, including any
	// sub-path it lives under.
	// Examples: "github.com", "example.com/gitea"
	Host string `json:"host"`
}

// HostFromURL returns the host of the given forge URL including any
// sub-path the forge lives under. The raw input is returned unchanged
// if it cannot be parsed as an absolute URL.
func HostFromURL(rawURL string) string {
	uri, err := url.Parse(rawURL)
	if err != nil || uri.Host == "" {
		return rawURL
	}

	host := uri.Host
	if path := strings.Trim(uri.Path, "/"); path != "" {
		host += "/" + path
	}

	return host
}
