package main

import (
	"testing"
)

func TestMake_URL_by_index(t *testing.T) {
	type T1 struct {
		k1, v string
		k2    int
	}
	test_data := []T1{
		T1{k1: "http://example.com/imgs/*.jpg", k2: 1, v: "http://example.com/imgs/1.jpg"},
		T1{k1: "http://example.com/imgs/***.jpg", k2: 0, v: "http://example.com/imgs/1.jpg"},
	}

	for _, e := range test_data {
		result := Make_URL_by_index(e.k1, e.k2)
		if result != e.v {
			t.Errorf("err %s %d->%s  But %s", e.k1, e.k2, e.v, result)
		}
	}

}
