package log

import "testing"

func TestL(t *testing.T) {
	l := L()
	if l == nil {
		t.Errorf("L() function retuned nil pointer")
	}
}

func TestS(t *testing.T) {
	s := S()
	if s == nil {
		t.Errorf("S() function retuned nil pointer")
	}
}

func TestForExample(t *testing.T) {
	ForExample()
	if log.logger == nil || log.sugar == nil {
		t.Errorf("ForExample function failed to create loggers :(")
	}
}

func TestForDevelopmant(t *testing.T) {
	ForDevelopment()
	if log.logger == nil || log.sugar == nil {
		t.Errorf("ForExample function failed to create loggers :(")
	}
}

func TestForProduction(t *testing.T) {
	ForProduction()
	if log.logger == nil || log.sugar == nil {
		t.Errorf("ForExample function failed to create loggers :(")
	}
}
