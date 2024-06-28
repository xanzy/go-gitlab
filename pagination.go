package gitlab

// AllPages can be used to fetch all pages of a paginated resource, e.g.: (assuming gl is a gitlab client instance)
//
//	allUsers, err := gitlab.AllPages(gl.Users.List, nil)
//
// It is also possible to specify additional pagination parameters:
//
//	mrs, err := gitlab.AllPages(
//		gl.MergeRequests.ListMergeRequests,
//		&gitlab.ListMergeRequestsOptions{
//			ListOptions: gitlab.ListOptions{
//				PerPage:    100,
//				Pagination: "keyset",
//				OrderBy:    "created_at",
//			},
//		},
//		gitlab.WithContext(ctx),
//	)
func AllPages[O, T any](f Paginatable[O, T], opt *O, optFunc ...RequestOptionFunc) ([]*T, error) {
	all := make([]*T, 0)
	nextLink := ""
	for {
		page, resp, err := f(opt, append(optFunc, WithKeysetPaginationParameters(nextLink))...)
		if err != nil {
			return nil, err
		}
		all = append(all, page...)
		if resp.NextLink == "" {
			break
		}
		nextLink = resp.NextLink
	}
	return all, nil
}

// Paginatable is the type implemented by list functions that return paginated content (e.g. [UsersService.ListUsers]).
// It works for top-level entities (e.g. users). See [PaginatableForID] for entities that require a parent ID (e.g.
// tags).
type Paginatable[O, T any] func(*O, ...RequestOptionFunc) ([]*T, *Response, error)

// AllPagesForID is similar to [AllPages] but for paginated resources that require a parent ID (e.g. tags of a project).
func AllPagesForID[O, T any](id any, f PaginatableForID[O, T], opt *O, optFunc ...RequestOptionFunc) ([]*T, error) {
	idFunc := func(opt *O, optFunc ...RequestOptionFunc) ([]*T, *Response, error) {
		return f(id, opt, optFunc...)
	}
	return AllPages(idFunc, opt, optFunc...)
}

// PaginatableForID is the type implemented by list functions that return paginated content for sub-entities (e.g.
// [TagsService.ListTags]).
// See also [Paginatable] for top-level entities (e.g. users).
type PaginatableForID[O, T any] func(any, *O, ...RequestOptionFunc) ([]*T, *Response, error)
