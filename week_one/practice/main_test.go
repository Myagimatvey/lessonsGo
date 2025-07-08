package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
2
2
3
4
5`

var testOkResult = `1
2
3
4
5
`

var testFail = `1
2
3
1
2`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)

	err := uniq(in, out)

	if err != nil {
		t.Errorf("Test for OK failed - Error")
	}

	if out.String() != testOkResult {
		t.Errorf("Test for OK failed - results not match")
	}
}

func TestForError(t *testing.T)  {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)

	err := uniq(in, out)

	if err == nil {
		t.Errorf("Test for OK failed - Error")
	}
}