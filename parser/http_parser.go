package parser

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type HttpRequest struct {
	Method      string
	Resource    string
	Protocol    string
	QueryParams map[string]string
	Headers     map[string]string
	Body        string
}

func ParseRequest(conn net.Conn) (*HttpRequest, error) {
	reader := bufio.NewReader(conn)

	routeLine, _ := reader.ReadString('\n')
	routeParts := strings.Fields(routeLine)
	if len(routeParts) != 3 {
		fmt.Printf("Error while parsing http request line: %s\n", routeLine)
		return nil, errors.New("wrong route line")
	}

	request := createRequest(routeParts)
	extractHeaders(request, reader)

	return request, nil
}

func extractHeaders(request *HttpRequest, reader *bufio.Reader) {
	request.Headers = make(map[string]string)
	for {
		header, _ := reader.ReadString('\n')

		header = strings.TrimSpace(header)
		if header == "" {
			break
		}
		headerParts := strings.Split(header, ":")
		request.Headers[headerParts[0]] = strings.TrimSpace(headerParts[1])
	}
}

func createRequest(requestParts []string) *HttpRequest {
	var request = &HttpRequest{}
	request.Method = requestParts[0]
	request.Protocol = requestParts[2]

	parts := strings.FieldsFunc(requestParts[1],
		func(r rune) bool {
			return r == '?' || r == '&'
		},
	)

	request.Resource = parts[0]

	for i := 1; i < len(parts); i++ {
		request.QueryParams = make(map[string]string)
		queryParam := strings.Split(parts[i], "=")
		request.QueryParams[queryParam[0]] = queryParam[1]
	}
	return request
}
