package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type OutputParser struct {
	// If true, vagrant output will be echoed to stdout. Default: false
	Verbose bool
}

func (parser OutputParser) startParser(reader io.ReadCloser, handler outputHandler, done func()) {
	defer done()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		parser.parseLine(scanner.Text(), handler)
	}
}

func (parser OutputParser) parseLine(line string, handler outputHandler) {
	parts := strings.Split(line, ",")
	if len(parts) < 4 {
		return
	}

	// parts[0] is a timestamp - we don't care about it
	target := parts[1]
	key := parts[2]
	message := make([]string, len(parts)-3)
	for i, part := range parts[3:] {
		message[i] = strings.Replace(part, "\\n", "\n", -1)
		message[i] = strings.Replace(message[i], "\\r", "\r", -1)
		message[i] = strings.Replace(message[i], "%!(VAGRANT_COMMA)", ",", -1)
	}

	if parser.Verbose && key == "ui" {
		outf := os.Stdout
		level := "info"
		msg := message
		if len(msg) > 1 {
			level = msg[0]
			msg = msg[1:]
		}
		if level == "error" {
			outf = os.Stderr
		}
		outf.WriteString(strings.Join(msg, ", "))
		outf.Write([]byte{byte('\n')})
	}

	handler.handleOutput(target, key, message)
}
