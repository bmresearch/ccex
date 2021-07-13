package logger

import (
	"errors"
	"os"
	"testing"
)

func Test_Logger(t *testing.T) {
	f, _ := os.Create("logger.log")
	Log.SetOut(f)
	Log.SetLevel(DEBUG)
	Log.Debug("debug log")
	Log.Debugf("%.8f", 0.2912101221212)
	Log.Info("info log")
	Log.Infof("%.8f", 0.2912101221212)
	Log.Warn(errors.New("test warn"))
	Log.Warnf("%.8f", 0.2912101221212)
	Log.Error(errors.New("test error"))
	Debug("debug log2")
	Info("info log2")
	Warn(errors.New("test warn2"))
	Error(errors.New("test error2"))
}

func Test_NewLogger(t *testing.T) {
	f, _ := os.Create("logger.log")
	logger := NewLogger()
	logger.SetOut(f)
	logger.SetLevel(DEBUG)
	logger.Debug("debug newlogger log")
	logger.Debugf("%.8f", 0.2912101221212)
	logger.Info("info newlogger log")
	logger.Infof("%.8f", 0.2912101221212)
	logger.Warn(errors.New("test newlogger error"))
	logger.Warnf("%.8f", 0.2912101221212)
}
