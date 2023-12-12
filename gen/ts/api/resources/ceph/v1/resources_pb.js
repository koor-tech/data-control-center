// @generated by protoc-gen-es v1.5.1
// @generated from file api/resources/ceph/v1/resources.proto (package api.resources.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message api.resources.ceph.v1.Resource
 */
export const Resource = proto3.makeMessageType(
  "api.resources.ceph.v1.Resource",
  () => [
    { no: 1, name: "Name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "Namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "Kind", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "Object", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.Resources
 */
export const Resources = proto3.makeMessageType(
  "api.resources.ceph.v1.Resources",
  () => [
    { no: 1, name: "resources", kind: "message", T: Resource, repeated: true },
  ],
);
