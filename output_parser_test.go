package main

import (
	"reflect"
	"strings"
	"testing"
)

type MockOutputParser struct {
	OutputParser
}

type ExpectedOutput struct {
	Target  string
	Key     string
	Message []string
	Seen    bool
}

type TestOutputHandler struct {
	ExpectedLines []*ExpectedOutput
}

func (parser MockOutputParser) Run(output string, handler outputHandler) {
	for _, line := range strings.Split(output, "\n") {
		parser.parseLine(line, handler)
	}
}

func (expected *ExpectedOutput) Matches(target, key string, message []string) bool {
	return expected.Target == target &&
		expected.Key == key &&
		reflect.DeepEqual(expected.Message, message)
}

func (handler *TestOutputHandler) handleOutput(target, key string, message []string) {
	for _, expected := range handler.ExpectedLines {
		if expected.Matches(target, key, message) {
			expected.Seen = true
		}
	}
}

func (handler TestOutputHandler) AssertAllSeen(t *testing.T) {
	for _, expected := range handler.ExpectedLines {
		if !expected.Seen {
			t.Errorf("Line not seen: %v", expected)
		}
	}
}

func TestVagrantOutputParser_parseLine(t *testing.T) {
	parser := OutputParser{}
	handler := TestOutputHandler{[]*ExpectedOutput{
		&ExpectedOutput{Target: "default", Key: "ui", Message: []string{"info", "This is a test, with commas.\nAnd newlines."}},
	}}
	parser.parseLine("123,default,ui,info,This is a test%!(VAGRANT_COMMA) with commas.\\nAnd newlines.", &handler)
	handler.AssertAllSeen(t)
}
