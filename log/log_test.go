package log_test

import (
	"testing"

	"github.com/taluship/x/log"
)

func TestNewLogger(t *testing.T) {
	log.Debugf("Debugf: %v", struct{ string }{"log message"})
	log.Infof("Infof: %v", struct{ string }{"log message"})
	log.Warnf("Warnf: %v", struct{ string }{"log message"})
	log.Errorf("Errorf: %v", struct{ string }{"log message"})
}
