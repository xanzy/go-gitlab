//go:build go1.23
// +build go1.23

package gitlab

import (
	"iter"
)

// Paginatable is the type implemented by list functions that return paginated
// content (e.g. [UsersService.ListUsers]).
// It works for top-level entities (e.g. users). See [PaginatableForID] for
// entities that require a parent ID (e.g. tags).
type Paginatable[O, T any] func(*O, ...RequestOptionFunc) ([]*T, *Response, error)

// AllPages is a [iter.Seq2] iterator to be used with any paginated resource.
// E.g. [UsersService.ListUsers]
//
//	for user, err := range gitlab.AllPages(gl.Users.ListUsers, nil) {
//		if err != nil {
//			// handle error
//		}
//		// process individual user
//	}
//
// It is also possible to specify additional pagination parameters:
//
//	for mr, err := range gitlab.AllPages(
//		gl.MergeRequests.ListMergeRequests,
//		&gitlab.ListMergeRequestsOptions{
//			ListOptions: gitlab.ListOptions{
//				PerPage:    100,
//				Pagination: "keyset",
//				OrderBy:    "created_at",
//			},
//		},
//		gitlab.WithContext(ctx),
//	) {
//		// ...
//	}
//
// Errors while fetching pages are returned as the second value of the iterator.
// It is the responsibility of the caller to handle them appropriately, e.g. by
// breaking the loop. The iteration will otherwise continue indefinitely,
// retrying to retrieve the erroring page on each iteration.
func AllPages[O, T any](f Paginatable[O, T], opt *O, optFunc ...RequestOptionFunc) iter.Seq2[*T, error] {
	return func(yield func(*T, error) bool) {
		nextLink := ""
		for {
			page, resp, err := f(opt, append(optFunc, WithKeysetPaginationParameters(nextLink))...)
			if err != nil {
				if !yield(nil, err) {
					return
				}
				continue
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

// PaginatableForID is the type implemented by list functions that return
// paginated content for sub-entities (e.g. [TagsService.ListTags]).
// See also [Paginatable] for top-level entities (e.g. users).
type PaginatableForID[O, T any] func(any, *O, ...RequestOptionFunc) ([]*T, *Response, error)

// AllPagesForID is similar to [AllPages] but for paginated resources that
// require a parent ID (e.g. tags of a project).
func AllPagesForID[O, T any](id any, f PaginatableForID[O, T], opt *O, optFunc ...RequestOptionFunc) iter.Seq2[*T, error] {
	idFunc := func(opt *O, optFunc ...RequestOptionFunc) ([]*T, *Response, error) {
		return f(id, opt, optFunc...)
	}
	return AllPages(idFunc, opt, optFunc...)
}
