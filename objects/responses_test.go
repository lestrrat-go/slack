package objects_test

import (
	"testing"

	"github.com/lestrrat-go/slack/objects"
	"github.com/stretchr/testify/assert"
)

func TestGenericResponse(t *testing.T) {
	r := objects.BuildGenericResponse().
		OK(true).
		MustBuild()

	if !assert.True(t, r.OK(), "ok = true") {
		return
	}
}
