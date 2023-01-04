package utils

import (
	"github.com/sirupsen/logrus"
)

func InitLogrus() {
	logrus.SetLevel(logrus.InfoLevel)
}
func init() {
	InitLogrus()
}
