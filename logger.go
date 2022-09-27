package nezukologger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	configureLogger()
}

func configureLogger() {
	Logger = logrus.New()

	Logger.SetFormatter(&nested.Formatter{
		HideKeys:      true,
		ShowFullLevel: true,
		TrimMessages:  true,
		CallerFirst:   true,
	})
}
