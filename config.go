// Copyright 2022 DomineCore.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package brisk

import "github.com/spf13/viper"

type Conf struct {
	viper.Viper
}

func NewConf() *Conf {
	return &Conf{*viper.New()}
}
