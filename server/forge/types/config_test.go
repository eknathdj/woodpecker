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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostFromURL(t *testing.T) {
	t.Parallel()

	t.Run("plain host", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "github.com", HostFromURL("https://github.com"))
	})

	t.Run("forge in sub-directory", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "example.com/gitea", HostFromURL("https://example.com/gitea"))
	})

	t.Run("trailing slash stripped", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "example.com/gitea", HostFromURL("https://example.com/gitea/"))
	})

	t.Run("unparseable input returned unchanged", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "not a url", HostFromURL("not a url"))
	})
}
