package config

import "github.com/spf13/viper"

// RCyclesConfig wrapper around Viper config, planning for future functions
type RCyclesConfig struct{
	*viper.Viper
}

func New() *RCyclesConfig {
	v := viper.New()
	rcc := new(RCyclesConfig)
	rcc.Viper = v
	return rcc
}
