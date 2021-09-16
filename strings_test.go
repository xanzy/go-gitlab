package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string
	Age int
	NickNames []string
	Address Address
	Company *Company
}

type Address struct {
	Street string
	City string
	Province string
	Country string
}

type Company struct {
	Name string
	Address Address
	Country string
}

var (
	response       = "gitlab.Person{Name:\"name\", Age:16, NickNames:[\"n\" \"a\" \"m\" \"e\"], Address:gitlab.Address{Street:\"\", City:\"\", Province:\"\", Country:\"\"}}"
	emptySliceResp = "gitlab.Person{Name:\"name\", Age:16, Address:gitlab.Address{Street:\"\", City:\"\", Province:\"\", Country:\"\"}}"
)

func TestStringify_nil(t *testing.T) {
	var person *Person
	resp := Stringify(person)
	assert.Equal(t, "<nil>", resp)
}

func TestStringify(t *testing.T) {
	person := &Person{"name", 16, []string {"n", "a", "m", "e"}, Address{}, nil}
	resp := Stringify(person)
	assert.Equal(t, response, resp)
}

func TestStringify_emptySlice(t *testing.T) {
	person := &Person{"name", 16, nil, Address{}, nil}
	resp := Stringify(person)
	assert.Equal(t, emptySliceResp, resp)
}
