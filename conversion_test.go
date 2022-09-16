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

func TestToMany(t *testing.T) {
	input := []string{`n return,count d smarty2^JVITALS4("","","","",1,"","","",1,I) f n1=1:1:V(0) {s V(n1)=$REPLACE(V(n1),":",",") s V(n1)=$REPLACE(V(n1),",",",    --")_$CHAR(13,10) s V(n1)=V(n1)_"," s numPieces=$L(V(n1),",") f n2=1:1:numPieces {i ($P(V(n1),",",n2)'["- Other (fill in)")&($P(V(n1),",",n2)'="") {s count=count+1 s return(count)=$P(V(n1),",",n2) } }} f n3=1:1:count {s V(n3)=return(n3)} s V(0)=count`}
	result := ToMany(input)
	expected := `n return,count 
	d smarty2^JVITALS4("","","","",1,"","","",1,I) 
	f n1=1:1:V(0) {
		s V(n1)=$REPLACE(V(n1),":",",") 
		s V(n1)=$REPLACE(V(n1),",",",    --")_$CHAR(13,10) 
		s V(n1)=V(n1)_"," 
		s numPieces=$L(V(n1),",") 
		f n2=1:1:numPieces {
			i ($P(V(n1),",",n2)'["- Other (fill in)")&($P(V(n1),",",n2)'="") {
				s count=count+1 
				s return(count)=$P(V(n1),",",n2) 
			} 
		}
	} 
	f n3=1:1:count {
		s V(n3)=return(n3)
	} 
	s V(0)=count`

	if result != expected {
		t.Errorf("got %q, wanted %q", result, expected)
	}

}
