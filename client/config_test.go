package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromFile(t *testing.T) {
	config, err := FromFile("./config_test.json")
	if err != nil {
		t.Errorf("Can't load test config: %q", err)
	}

	assert.Equal(t, config.ApiToken, "12345qwerty")
	assert.Equal(t, config.ProjectId, "6789")
	assert.Equal(t, "android_strings", config.Type)

	assert.Equal(t, 2, len(config.Languages))

	assert.Equal(t, "en", config.Languages[0].LangCode)
	assert.Equal(t, "de", config.Languages[1].LangCode)

	assert.Equal(t, "app/src/main/res/values/strings.xml", config.Languages[0].ExportPath)
	assert.Equal(t, "app/src/main/res/values-de/strings.xml", config.Languages[1].ExportPath)

	assert.Equal(t, 2, len(config.Tags))
	assert.Equal(t, "android", config.Tags[0])
	assert.Equal(t, "v2.5", config.Tags[1])

}
