package parser

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type URL struct {
	Scheme string `json:"scheme,omitempty"`
	Host   string `json:"host,omitempty"`
	Port   string `json:"port,omitempty"`
	Path   string `json:"path,omitempty"`
	Query  Query  `json:"query,omitempty"`
}

func (u *URL) WriteJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(u)
}

type Query map[string][]string

func Parse(s string) (*URL, error) {
	n, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	u := &URL{
		Scheme: n.Scheme,
		Host:   n.Host,
		Port:   n.Port(),
		Path:   n.Path,
	}

	q, err := parseRawQuery(n.RawQuery)
	u.Query = q

	return u, err
}

func parseRawQuery(raw string) (Query, error) {
	q := Query{}

	for len(raw) > 0 {
		elem := ""
		at := strings.Index(raw, "&")
		if at < 0 {
			elem = raw
			raw = ""
		} else {
			elem = raw[:at]
			raw = raw[at+1:]
		}

		at = strings.Index(elem, "=")
		if at < 1 {
			continue
		}

		key := elem[:at]
		if k, err := url.QueryUnescape(key); err == nil {
			key = k
		}

		val := elem[at+1:]
		if v, err := url.QueryUnescape(val); err == nil {
			val = v
		}

		if v, ok := q[key]; ok {
			q[key] = append(v, val)
		} else {
			q[key] = []string{val}
		}
	}

	return q, nil
}
