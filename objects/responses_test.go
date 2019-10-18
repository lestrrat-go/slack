package objects_test

import (
	"encoding/json"
	"testing"

	"github.com/lestrrat-go/slack/objects"
	"github.com/stretchr/testify/assert"
)

func TestGenericResponse(t *testing.T) {
	t.Run("building", func(t *testing.T) {
		t.Parallel()
		r := objects.BuildGenericResponse().
			OK(true).
			MustBuild()

		if !assert.True(t, r.OK(), "ok = true") {
			return
		}
	})
	t.Run("unmarshaling (false)", func(t *testing.T) {
		t.Parallel()
		const src = `{"ok": false}`
		var r objects.GenericResponse
		if !assert.NoError(t, json.Unmarshal([]byte(src), &r), "json.Unmarshal should succeed") {
			return
		}

		if !assert.True(t, !r.OK(), "r.OK() should be false") {
			return
		}
	})
	t.Run("unmarshaling", func(t *testing.T) {
		t.Parallel()
		const src = `{"ok": true, "channel": {"id": "foo"}}`

		var r struct {
			objects.GenericResponse
			objects.Channel `json:"channel"`
		}
		if !assert.NoError(t, json.Unmarshal([]byte(src), &r), "json.Unmarshal should succeed") {
			return
		}

		if !assert.True(t, r.OK(), "r.OK() should be true") {
			return
		}
		if !assert.Equal(t, r.ID(), "foo", "r.ID() should be foo") {
			return
		}
	})
}
