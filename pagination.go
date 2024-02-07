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

// Paginatable is the type that is implemented by all functions used to paginated content.
type Paginatable[O, T any] func(*O, ...RequestOptionFunc) ([]*T, *Response, error)
