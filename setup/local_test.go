package setup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txsvc/apikit/config"
)

func TestInitLocalProvider(t *testing.T) {
	config.SetProvider(NewLocalConfigProvider())

	cfg := config.GetConfig()
	assert.NotNil(t, cfg)

	assert.NotNil(t, cfg.Info())
	assert.NotNil(t, cfg.Settings())
	assert.NotEmpty(t, cfg.Settings().GetScopes())
}

func TestConfigLocation(t *testing.T) {
	config.SetProvider(NewLocalConfigProvider())

	cfg := config.GetConfig()
	assert.Equal(t, cfg, config.GetConfig())

	path := cfg.ConfigLocation()
	assert.NotEmpty(t, path)
	assert.Equal(t, config.DefaultConfigLocation, path)

	cfg.SetConfigLocation("$HOME/.config")
	assert.Equal(t, "$HOME/.config", cfg.ConfigLocation())
}

func TestGetSettings(t *testing.T) {
	conf := NewLocalConfigProvider().(*localConfig)
	assert.NotNil(t, conf)

	config.SetProvider(conf)
	ds := config.GetConfig().Settings()
	assert.NotNil(t, ds)
	assert.NotEmpty(t, ds)
}
