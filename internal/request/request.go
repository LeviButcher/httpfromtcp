package request

import (
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	raw, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(raw), "\r\n")

	rl, err := parseRequestLine(lines[0])
	if err != nil {
		return nil, err
	}
	request := Request{
		RequestLine: *rl,
	}

	return &request, nil
}

type ParseError struct{}

func (m *ParseError) Error() string {
	return "bad parse"
}

// GET /coffee HTTP/1.1
func parseRequestLine(line string) (*RequestLine, error) {
	words := strings.Split(line, " ")
	if len(words) != 3 {
		return nil, &ParseError{}
	}
	method := words[0]
	target := words[1]
	version := strings.TrimPrefix(words[2], "HTTP/")
	if version != "1.1" {
		return nil, &ParseError{}
	}

	return &RequestLine{
		HttpVersion:   version,
		RequestTarget: target,
		Method:        method,
	}, nil
}
