package main

import (
	"github.com/DiLRandI/circle-ci/calc/intcalc"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("test CircleCI app started")

	c := intcalc.New()

	logrus.Info("Adding 2 numbers")
	logrus.Infof("2+5=%d", c.Add(2, 5))
}
