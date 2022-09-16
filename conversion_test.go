package main

import (
	"testing"
)

func TestToOne(t *testing.T) {
	input := []string{`"" n lastFilledTTR,filledTodayTTR,output`, `s lastFilledTTR=$$Main^SNETLINK(patID,$$zoDTap^EAXLIB2("EPT",patID,25231),"EPT25231")`, `s lastFilledTTR=$JUSTIFY(lastFilledTTR/365,0,1)`}
	result := ToOne(input)
	expected := `"" n lastFilledTTR,filledTodayTTR,output s lastFilledTTR=$$Main^SNETLINK(patID,$$zoDTap^EAXLIB2("EPT",patID,25231),"EPT25231") s lastFilledTTR=$JUSTIFY(lastFilledTTR/365,0,1)`

	if result != expected {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}

/*func TestToMany(t *testing.T) {
	result := ToMany()
	expected :=

}*/
