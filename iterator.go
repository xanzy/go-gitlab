//go:build go1.22 && goexperiment.rangefunc
// +build go1.22,goexperiment.rangefunc

package gitlab

import (
	"iter"
)

// PageIterator is an EXPERIMENTAL iterator as defined in the "rangefunc" experiment for go 1.22.
// See https://go.dev/wiki/RangefuncExperiment for more details.
//
// It can be used as:
//
//	for user, err := range gitlab.PageIterator(gl.Users.List, nil) {
//		if err != nil {
//			// handle error
//		}
//		// process individual user
//	}
func PageIterator[O, T any](f Paginatable[O, T], opt *O, optFunc ...RequestOptionFunc) iter.Seq2[*T, error] {
	return func(yield func(*T, error) bool) {
		nextLink := ""
		for {
			page, resp, err := f(opt, append(optFunc, WithKeysetPaginationParameters(nextLink))...)
			if err != nil {
				yield(nil, err)
				return
			}
			for _, p := range page {
				if !yield(p, nil) {
					return
				}
			}
			if resp.NextLink == "" {
				break
			}
			nextLink = resp.NextLink
		}
	}
}

// PageIteratorForID is similar to [PageIterator] but for paginated resources that require a parent ID (e.g. tags of a project).
func PageIteratorForID[O, T any](id any, f PaginatableForID[O, T], opt *O, optFunc ...RequestOptionFunc) iter.Seq2[*T, error] {
	idFunc := func(opt *O, optFunc ...RequestOptionFunc) ([]*T, *Response, error) {
		return f(id, opt, optFunc...)
	}
	return PageIterator(idFunc, opt, optFunc...)
}
