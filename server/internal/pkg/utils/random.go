package utils

import (
	"github.com/jaevor/go-nanoid"
	"github.com/sirupsen/logrus"
)

var GenerateKey func() string

func init() {
	var err error
	GenerateKey, err = nanoid.Standard(21)
	if err != nil {
		logrus.Fatalln(err)
	}
}
