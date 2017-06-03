package log

import "testing"

func TestLogs(t *testing.T) {
	Info("camembert")
	Infof("coeur de %v", "lion")

	Warn("J'aime les filles")
	Warnf("Bancs %v", "publics")

	Error("Bierre et chips")
	Errorf("les frites ne viennent pas de %v, c'est Français !", "Belgique")
}

func TestPanic(t *testing.T) {
	defer func() { t.Log(recover()) }()
	Panic("Et une panique !")
	t.Error("should panic")
}

func TestPanicf(t *testing.T) {
	defer func() { t.Log(recover()) }()
	Panicf("Fiiouuuuuuuuu")
	t.Error("should panic")
}

func TestAddCallerFormat(t *testing.T) {
	format := "hello"
	expected := "%vhello"

	if format = addCallerFormat(format); format != expected {
		t.Errorf("format should be %v: %v", expected, format)
	}
}
