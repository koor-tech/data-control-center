// @generated by protoc-gen-es v1.6.0
// @generated from file api/services/stats/v1/stats.proto (package api.services.stats.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ClusterRadar, ClusterStats } from "../../../resources/ceph/v1/stats_pb.js";
import { NodeInfo, ResourceInfo } from "../../../resources/k8s/v1/resources_pb.js";

/**
 * @generated from message api.services.stats.v1.GetClusterStatsRequest
 */
export const GetClusterStatsRequest = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterStatsRequest",
  [],
);

/**
 * @generated from message api.services.stats.v1.GetClusterStatsResponse
 */
export const GetClusterStatsResponse = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterStatsResponse",
  () => [
    { no: 1, name: "stats", kind: "message", T: ClusterStats },
  ],
);

/**
 * @generated from message api.services.stats.v1.GetClusterResourcesRequest
 */
export const GetClusterResourcesRequest = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterResourcesRequest",
  [],
);

/**
 * @generated from message api.services.stats.v1.GetClusterResourcesResponse
 */
export const GetClusterResourcesResponse = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterResourcesResponse",
  () => [
    { no: 1, name: "resources", kind: "message", T: ResourceInfo, repeated: true },
    { no: 2, name: "deployments", kind: "message", T: ResourceInfo, repeated: true },
  ],
);

/**
 * @generated from message api.services.stats.v1.GetClusterNodesRequest
 */
export const GetClusterNodesRequest = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterNodesRequest",
  [],
);

/**
 * @generated from message api.services.stats.v1.GetClusterNodesResponse
 */
export const GetClusterNodesResponse = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterNodesResponse",
  () => [
    { no: 1, name: "nodes", kind: "message", T: NodeInfo, repeated: true },
  ],
);

/**
 * @generated from message api.services.stats.v1.GetClusterRadarRequest
 */
export const GetClusterRadarRequest = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterRadarRequest",
  [],
);

/**
 * @generated from message api.services.stats.v1.GetClusterRadarResponse
 */
export const GetClusterRadarResponse = proto3.makeMessageType(
  "api.services.stats.v1.GetClusterRadarResponse",
  () => [
    { no: 1, name: "radar", kind: "message", T: ClusterRadar },
  ],
);

