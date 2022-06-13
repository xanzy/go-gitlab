
package gitlab

type ListableOptions interface{
	GetPage() int
	SetPage(int)
}

func (o *ListOptions) GetPage() int {
	return o.Page
}

func (o *ListOptions) SetPage(page int) {
	o.Page = page
}

func Collect[T any, U ListableOptions](f func(U, ...RequestOptionFunc) ([]*T, *Response, error), opt U, options ...RequestOptionFunc) ([]*T, error) {
	var collection []*T
	opt.SetPage(1)
	for opt.GetPage() != 0 {
		page, resp, err := f(opt, options...)
		if err != nil {
			return nil, err
		}

		collection = append(collection, page...)
		opt.SetPage(resp.NextPage)
	}

	return collection, nil
}