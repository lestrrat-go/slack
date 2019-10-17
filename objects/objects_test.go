package objects_test

import (
	"testing"

	"github.com/lestrrat-go/slack/objects"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	o, err := objects.BuildText(objects.MarkdownTextType, "*Hello* World").Build()
	if !assert.NoError(t, err, `text object should return no errors`) {
		return
	}

	if !assert.Equal(t, o.Type(), objects.MarkdownTextType) {
		return
	}

	if !assert.Equal(t, o.Text(), "*Hello* World") {
		return
	}
}

func TestSectionBlock(t *testing.T) {
	b, err := objects.BuildSectionBlock(objects.MarkdownText("*Hello* World")).Build()
	if !assert.NoError(t, err, `building a section block should return no errors`) {
		return
	}

	if !assert.Equal(t, b.Type(), objects.SectionBlockType) {
		return
	}
}

func TestContextBlock(t *testing.T) {
	b, err := objects.BuildContextBlock(objects.MarkdownText("*Hello* World")).Build()
	if !assert.NoError(t, err, `building a context block should return no errors`) {
		return
	}

	if !assert.Equal(t, b.Type(), objects.ContextBlockType) {
		return
	}
}

