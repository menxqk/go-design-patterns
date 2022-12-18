package command

import (
	"testing"
)

func TestMementoFacadeSaveSettings(t *testing.T) {
	m := MementoFacade{}

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	c := m.RestoreSettings()
	cType, ok := c.(Mute)
	if !ok {
		t.Errorf("should return Mute type, got %T", cType)
	}

	d := m.RestoreSettings()
	dType, ok := d.(Volume)
	if !ok {
		t.Errorf("should return Volume type, got %T", dType)
	}

}
