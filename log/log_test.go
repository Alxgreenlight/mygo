package log

import "testing"

func TestForExample(t *testing.T) {
	ForExample()
	if log.logger == nil || log.sugar == nil {
		t.Errorf("ForExample function failed to create loggers :(")
	}
}
