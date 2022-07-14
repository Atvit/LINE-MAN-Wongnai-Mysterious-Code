package mysterious

import "encoding/base64"

type Mysterious interface {
	Decode(s string) (string, error)
	Reverse(s string) string
}

type mysterious struct{}

func NewMysterious() Mysterious {
	return mysterious{}
}

func (m mysterious) Decode(s string) (string, error) {
	sd, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", ErrorIllegalData
	}

	return string(sd), nil
}

func (m mysterious) Reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}

	return result
}
