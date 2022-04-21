// Copyright 2022 DomineCore.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package brisk

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config Conf

type Conf struct {
	viper.Viper
}

func NewConf() *Conf {
	return &Conf{*viper.New()}
}

func SetConf(path string, name string, filetype string) {
	Config = *NewConf()
	Config.AddConfigPath(path)
	Config.SetConfigName(name)
	Config.SetConfigType(filetype)
	err := Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
