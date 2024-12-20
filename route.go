package mbr

import "context"

type RouterHandleFunc func(ctx context.Context)

type Route struct {
	name    string
	Path    string
	HandleF RouterHandleFunc
}