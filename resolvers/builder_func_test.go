// +build func_test

package resolvers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		version string
		archive string
	}{
		{desc: "should build a new builder from demo", name: "github.com/cgentron/plugin-demo", version: "master", archive: "https://github.com/cgentron/plugin-demo/archive/main.zip"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := NewClient()
			assert.NoError(t, err)

			ctx := context.Background()

			str, err := c.Fetch(ctx, tt.name, tt.version, tt.archive)
			assert.NoError(t, err)
			assert.Empty(t, "", str)

			err = c.Unzip(tt.name, tt.version)
			assert.NoError(t, err)

			resolvers := map[string]Descriptor{
				"demo": {ModuleName: tt.name, Version: tt.version},
			}

			b, err := NewBuilder(c, resolvers)
			assert.NoError(t, err)
			assert.NotNil(t, b)
		})
	}
}
