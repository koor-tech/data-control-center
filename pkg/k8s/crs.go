package k8s

import (
	"context"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func GetCustomResource[T any](ctx context.Context, client *kubernetes.Clientset, resSchame schema.GroupVersionResource, resource string, name string, opts metav1.GetOptions) (*T, error) {
	cliSet := dynamic.New(client.RESTClient())
	obj, err := cliSet.Resource(resSchame).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	bytes, err := obj.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var crdObj T
	if err := json.Unmarshal(bytes, &crdObj); err != nil {
		return nil, err
	}

	return &crdObj, nil
}

func ListCustomResource[T any](ctx context.Context, client *kubernetes.Clientset, resSchame schema.GroupVersionResource, resource string, opts metav1.ListOptions) ([]*T, error) {
	cliSet := dynamic.New(client.RESTClient())
	list, err := cliSet.Resource(resSchame).List(ctx, opts)
	if err != nil {
		return nil, err
	}

	items := []*T{}
	for _, obj := range list.Items {
		bytes, err := obj.MarshalJSON()
		if err != nil {
			return nil, err
		}

		var crdObj T
		if err := json.Unmarshal(bytes, &crdObj); err != nil {
			return nil, err
		}

		items = append(items, &crdObj)
	}

	return items, nil
}
