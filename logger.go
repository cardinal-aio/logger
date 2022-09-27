package logger

import (
	"fmt"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	configureLogger()
}

func configureLogger() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&nested.Formatter{
		HideKeys:      true,
		ShowFullLevel: true,
		TrimMessages:  true,
		CallerFirst:   true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d, %s()", filename, f.Line, f.Function)
		},
	})
}
