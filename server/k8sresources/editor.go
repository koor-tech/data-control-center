package k8sresources

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	k8sv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/k8s/v1"
	k8sresourcespb "github.com/koor-tech/data-control-center/gen/go/api/services/k8sresources/v1"
)

func (s *Server) GetResources(ctx context.Context, _ *connect.Request[k8sresourcespb.GetResourcesRequest]) (*connect.Response[k8sresourcespb.GetResourcesResponse], error) {
	var resources []*k8sv1.Resource

	cephResources, err := s.k.GetYamlCephResources(ctx, s.Namespace)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by %w", err))
	}

	for _, resource := range cephResources {
		resources = append(resources, &k8sv1.Resource{
			Name:      resource.Name,
			Namespace: resource.Namespace,
			Content:   resource.Content,
			Kind:      resource.Kind,
			Object:    resource.Object,
		})
	}

	res := connect.NewResponse(&k8sresourcespb.GetResourcesResponse{
		Resources: &k8sv1.Resources{
			Resources: resources,
		},
	})
	return res, nil
}

func (s *Server) SaveResource(ctx context.Context, req *connect.Request[k8sresourcespb.SaveResourceRequest]) (*connect.Response[k8sresourcespb.SaveResourceResponse], error) {
	requestResource := req.Msg.Resource
	resource := k8sv1.Resource{
		Name:      requestResource.Name,
		Namespace: requestResource.Namespace,
		Content:   requestResource.Content,
		Kind:      requestResource.Kind,
		Object:    requestResource.Object,
	}

	if len(resource.Name) == 0 || len(resource.Content) == 0 {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by empty file"))
	}

	err := s.k.SaveResource(ctx, &resource)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by %w", err))
	}

	res := connect.NewResponse(&k8sresourcespb.SaveResourceResponse{
		Resource: &resource,
	})
	return res, nil
}
