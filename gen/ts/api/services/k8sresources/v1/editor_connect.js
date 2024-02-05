// @generated by protoc-gen-connect-es v1.3.0
// @generated from file api/services/k8sresources/v1/editor.proto (package api.services.k8sresources.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetResourcesRequest, GetResourcesResponse, SaveResourceRequest, SaveResourceResponse } from "./editor_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.services.k8sresources.v1.K8sResourcesService
 */
export const K8sResourcesService = {
  typeName: "api.services.k8sresources.v1.K8sResourcesService",
  methods: {
    /**
     * @generated from rpc api.services.k8sresources.v1.K8sResourcesService.GetResources
     */
    getResources: {
      name: "GetResources",
      I: GetResourcesRequest,
      O: GetResourcesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.k8sresources.v1.K8sResourcesService.SaveResource
     */
    saveResource: {
      name: "SaveResource",
      I: SaveResourceRequest,
      O: SaveResourceResponse,
      kind: MethodKind.Unary,
    },
  }
};

