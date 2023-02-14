package testjsoniter_test

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type OverlappingKey1 struct {
	Foo string
	Bar string `json:"Baz"`
	Baz string
}

type OverlappingKey2 struct {
	Foo string
	Bar string `json:"Bar"`
	Baz string `json:"Bar"`
}

type OverlappingKey3 struct {
	Foo string
	Bar string `json:"Baz"`
	Baz string
	Qux string `json:"Baz"`
}

type Sub1 struct {
	Foo string
	Bar string `json:"Bar"`
}

type OverlappingKey4 struct {
	Foo string
	Bar string
	Baz string
	Sub1
}

type Recursive1 struct {
	R string `json:"r"`
	Recursive2
}

type Recursive2 struct {
	R  string `json:"r"`
	RR string `json:"rr"`
	*OverlappingKey5
}

type OverlappingKey5 struct {
	Foo string
	Recursive1
}

func TestMimic_overlapping_key(t *testing.T) {
	for _, config := range []jsoniter.API{
		jsoniter.ConfigCompatibleWithStandardLibrary,
		jsoniter.ConfigDefault,
	} {
		for _, v := range []any{
			OverlappingKey1{"foo", "bar", "baz"},
			OverlappingKey2{"foo", "bar", "baz"},
			OverlappingKey3{"foo", "bar", "baz", "qux"},
			OverlappingKey4{Foo: "foo", Bar: "bar", Baz: "baz", Sub1: Sub1{Foo: "foofoo", Bar: "barbar"}},
			// These cause stack overflow
			// OverlappingKey5{Foo: "foo", Recursive1: Recursive1{R: "r", Recursive2: Recursive2{R: "rr"}}},
			// OverlappingKey5{Foo: "foo", Recursive1: Recursive1{R: "r", Recursive2: Recursive2{R: "rr", OverlappingKey5: &OverlappingKey5{Foo: "foo"}}}},
		} {
			binOrg, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			binMimic, err := config.Marshal(v)
			if err != nil {
				panic(err)
			}

			str1, str2 := string(binOrg), string(binMimic)
			if str1 != str2 {
				t.Errorf("not same, type = %T. org = %s, jsoniter = %s\n", v, str1, str2)
			}

		}
	}
}
