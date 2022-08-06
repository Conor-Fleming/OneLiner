package main

import (
	"testing"
)

/*helper functions to make reading tests easier

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}*/

type addTest struct {
	arg1     []string
	expected string
}

var addTests = []addTest{
	addTest{[]string{`"" n lastFilledTTR,filledTodayTTR,output`, `s lastFilledTTR=$$Main^SNETLINK(patID,$$zoDTap^EAXLIB2("EPT",patID,25231),"EPT25231")`, `s lastFilledTTR=$JUSTIFY(lastFilledTTR/365,0,1)`},
		`"" n lastFilledTTR,filledTodayTTR,output s lastFilledTTR=$$Main^SNETLINK(patID,$$zoDTap^EAXLIB2("EPT",patID,25231),"EPT25231") s lastFilledTTR=$JUSTIFY(lastFilledTTR/365,0,1)`},
}

func TestCoversionToOne(t *testing.T) {
	for _, test := range addTests {
		if output := ToOne(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestConversionToMany(t *testing.T) {

}
