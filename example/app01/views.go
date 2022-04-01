package app01

import (
	"brisk"
)

func AppGet(c *brisk.Context) {

	// c.WriteString(http.StatusAccepted, "app01 hello")
	c.ResponseWriter.Write([]byte("123"))
}
