package testjsoniter_test

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type Quote struct {
	Foo int  `json:",string"`
	Bar *int `json:",string"`
	Baz Sub  `json:",string"`
}

type Sub struct {
	Qux string
}

func TestQuote_behavior(t *testing.T) {
	v := Quote{
		Foo: 10,
		Bar: nil,
		Baz: Sub{
			Qux: "qux",
		},
	}

	binOrg, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	binMimic, err := jsoniter.Marshal(v)
	if err != nil {
		panic(err)
	}

	str1, str2 := string(binOrg), string(binMimic)
	if str1 != str2 {
		t.Errorf("not same, type = %T. org = %s, jsoniter = %s\n", v, str1, str2)
	}
}
