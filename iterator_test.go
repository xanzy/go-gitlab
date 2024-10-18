//go:build go1.23
// +build go1.23

package gitlab

import (
	"errors"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPages(t *testing.T) {
	type foo struct{ string }
	type listFooOpt struct{}

	type iteration struct {
		foo *foo
		err error
	}

	sentinelError := errors.New("sentinel error")

	// assertSeq is a helper function to assert the sequence of iterations.
	// It is necessary because the iteration may be endless (e.g. in the error
	// case).
	assertSeq := func(t *testing.T, expected []iteration, actual iter.Seq2[*foo, error]) {
		t.Helper()
		i := 0
		for actualFoo, actualErr := range actual {
			if i >= len(expected) {
				t.Errorf("unexpected iteration: %v, %v", actualFoo, actualErr)
				break
			}
			assert.Equal(t, expected[i].foo, actualFoo)
			assert.Equal(t, expected[i].err, actualErr)
			i++
		}

		if i < len(expected) {
			t.Errorf("expected %d more iterations", len(expected)-i)
		}
	}

	type args struct {
		f       Paginatable[listFooOpt, foo]
		opt     *listFooOpt
		optFunc []RequestOptionFunc
	}
	tests := []struct {
		name string
		args args
		want []iteration
	}{
		{
			name: "empty",
			args: args{
				f: func() Paginatable[listFooOpt, foo] {
					return func(*listFooOpt, ...RequestOptionFunc) ([]*foo, *Response, error) {
						return []*foo{}, &Response{}, nil
					}
				}(),
			},
			want: []iteration{},
		},
		{
			name: "single element, no errors",
			args: args{
				f: func() Paginatable[listFooOpt, foo] {
					return func(*listFooOpt, ...RequestOptionFunc) ([]*foo, *Response, error) {
						return []*foo{{"foo"}}, &Response{}, nil
					}
				}(),
			},
			want: []iteration{
				{foo: &foo{"foo"}, err: nil},
			},
		},
		{
			name: "one error than success",
			args: args{
				f: func() Paginatable[listFooOpt, foo] {
					called := false
					return func(*listFooOpt, ...RequestOptionFunc) ([]*foo, *Response, error) {
						if !called {
							called = true
							return []*foo{}, &Response{}, sentinelError
						}
						return []*foo{{"foo"}}, &Response{}, nil
					}
				}(),
			},
			want: []iteration{
				{foo: nil, err: sentinelError},
				{foo: &foo{"foo"}, err: nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertSeq(t, tt.want, AllPages(tt.args.f, tt.args.opt, tt.args.optFunc...))
		})
	}
}
