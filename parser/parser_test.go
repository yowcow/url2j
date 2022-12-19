package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	cases := []struct {
		input    string
		expected *URL
	}{
		{
			"/hoge/",
			&URL{
				Path:  "/hoge/",
				Query: Query{},
			},
		},
		{
			"http://example.com/",
			&URL{
				Scheme: "http",
				Host:   "example.com",
				Path:   "/",
				Query:  Query{},
			},
		},
		{
			"/?foo&bar",
			&URL{
				Path:  "/",
				Query: Query{},
			},
		},
		{
			"/?foo=",
			&URL{
				Path: "/",
				Query: Query{
					"foo": []string{""},
				},
			},
		},
		{
			"/?foo=1",
			&URL{
				Path: "/",
				Query: Query{
					"foo": []string{"1"},
				},
			},
		},
		{
			"/?foo=&bar=",
			&URL{
				Path: "/",
				Query: Query{
					"foo": []string{""},
					"bar": []string{""},
				},
			},
		},
		{
			"/?foo=1&",
			&URL{
				Path: "/",
				Query: Query{
					"foo": []string{"1"},
				},
			},
		},
		{
			"/?foo=1&bar=2",
			&URL{
				Path: "/",
				Query: Query{
					"foo": []string{"1"},
					"bar": []string{"2"},
				},
			},
		},
		{
			"/?foo=1&bar=2&bar=3&buz=",
			&URL{
				Path: "/",
				Query: Query{
					"foo": []string{"1"},
					"bar": []string{"2", "3"},
					"buz": []string{""},
				},
			},
		},
		{
			"/?foo=%26%3D%25",
			&URL{
				Path: "/",
				Query: Query{
					"foo": []string{"&=%"},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			actual, err := Parse(c.input)
			if err != nil {
				t.Fatal("failed parsing url:", err)
			}
			if d := cmp.Diff(c.expected, actual); d != "" {
				t.Error("expected no diff but got:", d)
			}
		})
	}
}
