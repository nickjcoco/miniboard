package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8849")
	viper.SetDefault("THEME", "flatly")
	viper.SetDefault("COLOR", "dark")
	viper.SetDefault("COLORON", "#89ff89")
	viper.SetDefault("COLOROFF", "#ff7171")
	viper.SetDefault("BTNWIDTH", "180px")
	viper.SetDefault("WEBREFRESH", "60")
	viper.SetDefault("DBTRIMDAYS", "30")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Color, _ = viper.Get("COLOR").(string)
	config.ColorOn, _ = viper.Get("COLORON").(string)
	config.ColorOff, _ = viper.Get("COLOROFF").(string)
	config.BtnWidth, _ = viper.Get("BTNWIDTH").(string)
	config.WebRefresh, _ = viper.Get("WEBREFRESH").(string)
	config.DBTrimDays, _ = viper.Get("DBTRIMDAYS").(string)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("color", config.Color)
	viper.Set("coloron", config.ColorOn)
	viper.Set("coloroff", config.ColorOff)
	viper.Set("btnwidth", config.BtnWidth)
	viper.Set("webrefresh", config.WebRefresh)
	viper.Set("dbtrimdays", config.DBTrimDays)

	err := viper.WriteConfig()
	check.IfError(err)
}
