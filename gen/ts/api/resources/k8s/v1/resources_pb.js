// @generated by protoc-gen-es v1.7.2
// @generated from file api/resources/k8s/v1/resources.proto (package api.resources.k8s.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum api.resources.k8s.v1.ReliabilityScore
 */
export const ReliabilityScore = proto3.makeEnum(
  "api.resources.k8s.v1.ReliabilityScore",
  [
    {no: 0, name: "RELIABILITY_SCORE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "RELIABILITY_SCORE_UNKNOWN", localName: "UNKNOWN"},
    {no: 2, name: "RELIABILITY_SCORE_NONE", localName: "NONE"},
    {no: 3, name: "RELIABILITY_SCORE_DEGRADED", localName: "DEGRADED"},
    {no: 4, name: "RELIABILITY_SCORE_OK", localName: "OK"},
  ],
);

/**
 * @generated from enum api.resources.k8s.v1.ResourceStatus
 */
export const ResourceStatus = proto3.makeEnum(
  "api.resources.k8s.v1.ResourceStatus",
  [
    {no: 0, name: "RESOURCE_STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "RESOURCE_STATUS_UNKNOWN", localName: "UNKNOWN"},
    {no: 2, name: "RESOURCE_STATUS_READY", localName: "READY"},
    {no: 3, name: "RESOURCE_STATUS_PROGRESSING", localName: "PROGRESSING"},
    {no: 4, name: "RESOURCE_STATUS_NOT_READY", localName: "NOT_READY"},
  ],
);

/**
 * @generated from message api.resources.k8s.v1.Resource
 */
export const Resource = proto3.makeMessageType(
  "api.resources.k8s.v1.Resource",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "kind", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "object", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.resources.k8s.v1.Resources
 */
export const Resources = proto3.makeMessageType(
  "api.resources.k8s.v1.Resources",
  () => [
    { no: 1, name: "resources", kind: "message", T: Resource, repeated: true },
  ],
);

/**
 * @generated from message api.resources.k8s.v1.ResourceInfo
 */
export const ResourceInfo = proto3.makeMessageType(
  "api.resources.k8s.v1.ResourceInfo",
  () => [
    { no: 1, name: "apiversion", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "kind", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "status", kind: "enum", T: proto3.getEnumType(ResourceStatus) },
    { no: 6, name: "replicas", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "reliability", kind: "enum", T: proto3.getEnumType(ReliabilityScore) },
    { no: 8, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * @generated from message api.resources.k8s.v1.NodeInfo
 */
export const NodeInfo = proto3.makeMessageType(
  "api.resources.k8s.v1.NodeInfo",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(ResourceStatus) },
    { no: 3, name: "roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "internal_ip", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "external_ip", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "age", kind: "message", T: Timestamp, opt: true },
  ],
);

