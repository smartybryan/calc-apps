package handlers

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"

	calc "github.com/smartybryan/calc-lib"
)

func TestCSVHandler(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "[TEST]", 0)
	rawInput := `1,+,2
3,-,2
Nan,+,1
1,+,Nan
4,+,5
1,2,3,4
1,2
`
	input := strings.NewReader(rawInput)
	var output bytes.Buffer

	handler := NewCSVHandler(logger, input, &output, map[string]Calculator{"+": &calc.Addition{}})
	err := handler.Handle()
	assertError(t, err, nil)
	expected := `1,+,2,3
4,+,5,9
`
	if output.String() != expected {
		t.Errorf("expected: [%s], got: [%s]", expected, output.String())
	}

	t.Log(logBuffer.String())
}

func TestCSVHanlder_WriteError(t *testing.T) {
	reader := strings.NewReader("1,+,2")
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "[TEST]", 0)
	writer := &ErringWriter{err: boink}
	handler := NewCSVHandler(logger, reader, writer, map[string]Calculator{"+": &calc.Addition{}})
	err := handler.Handle()
	assertError(t, err, boink)
}

func TestCSVHandler_ReadError(t *testing.T) {
	reader := ErringReader{err: boink}
	handler := NewCSVHandler(nil, reader, nil, nil)
	err := handler.Handle()
	assertError(t, err, boink)
}

//////////////////////////////////////////////////////////

var boink = errors.New("boink")

type ErringReader struct {
	err error
}

func (this ErringReader) Read(_ []byte) (n int, err error) {
	return 0, this.err
}
