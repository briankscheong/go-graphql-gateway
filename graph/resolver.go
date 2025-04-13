package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/briankscheong/go-graphql-gateway/graph/model"
	"k8s.io/client-go/kubernetes"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos     []*model.Todo
	K8sClient *kubernetes.Clientset
}
