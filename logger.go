package logger

import (
	"fmt"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

var Logger *_Logger

type _Logger struct {
	logrus.Logger
}

func init() {
	configureLogger()
}

func configureLogger() {
	l := logrus.New()
	l1 := _Logger{*l}
	Logger = &l1
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&nested.Formatter{
		HideKeys:      true,
		ShowFullLevel: true,
		TrimMessages:  true,
		CallerFirst:   true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			filename := path.Base(f.File)
			return fmt.Sprintf(" | %s:%d %s()", filename, f.Line, f.Function)
		},
	})
}

func (l *_Logger) SpewDebug(obj any) {
	Logger.Debug(spew.Sdump(obj))
}
