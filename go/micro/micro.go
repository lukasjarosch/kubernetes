// Package micro implements a go-micro service for k8s
package micro

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/go-plugins/selector/cache"
)

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()

	s := cache.NewSelector(
		selector.Registry(r),
	)

	// set the registry and selector
	options := []micro.Option{
		micro.Registry(r),
		micro.Selector(s),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return grpc.NewService(options...)
}