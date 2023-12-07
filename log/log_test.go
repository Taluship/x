package log_test

import (
	"testing"
	"time"

	"github.com/taluship/x/log"
)

func TestNewLogger(t *testing.T) {
	defer log.ProfileLogf(time.Now(), "profile log message")
	log.Debugf("Debugf: %v", struct{ string }{"log message"})
	log.Infof("Infof: %v", struct{ string }{"log message"})
	log.Warnf("Warnf: %v", struct{ string }{"log message"})
	log.Errorf("Errorf: %v", struct{ string }{"log message"})
}
