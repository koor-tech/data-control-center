// @generated by protoc-gen-es v1.3.1
// @generated from file api/services/stats/stats.proto (package stats, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ClusterRadar, ClusterStats, KoorCluster, NodeInfo, ResourceInfo } from "../../resources/stats/stats_pb.js";

/**
 * @generated from message stats.GetClusterStatsRequest
 */
export const GetClusterStatsRequest = proto3.makeMessageType(
  "stats.GetClusterStatsRequest",
  [],
);

/**
 * @generated from message stats.GetClusterStatsResponse
 */
export const GetClusterStatsResponse = proto3.makeMessageType(
  "stats.GetClusterStatsResponse",
  () => [
    { no: 1, name: "stats", kind: "message", T: ClusterStats },
  ],
);

/**
 * @generated from message stats.GetClusterResourcesRequest
 */
export const GetClusterResourcesRequest = proto3.makeMessageType(
  "stats.GetClusterResourcesRequest",
  [],
);

/**
 * @generated from message stats.GetClusterResourcesResponse
 */
export const GetClusterResourcesResponse = proto3.makeMessageType(
  "stats.GetClusterResourcesResponse",
  () => [
    { no: 1, name: "resources", kind: "message", T: ResourceInfo, repeated: true },
    { no: 2, name: "deployments", kind: "message", T: ResourceInfo, repeated: true },
  ],
);

/**
 * @generated from message stats.GetClusterNodesRequest
 */
export const GetClusterNodesRequest = proto3.makeMessageType(
  "stats.GetClusterNodesRequest",
  [],
);

/**
 * @generated from message stats.GetClusterNodesResponse
 */
export const GetClusterNodesResponse = proto3.makeMessageType(
  "stats.GetClusterNodesResponse",
  () => [
    { no: 1, name: "nodes", kind: "message", T: NodeInfo, repeated: true },
  ],
);

/**
 * @generated from message stats.GetClusterRadarRequest
 */
export const GetClusterRadarRequest = proto3.makeMessageType(
  "stats.GetClusterRadarRequest",
  [],
);

/**
 * @generated from message stats.GetClusterRadarResponse
 */
export const GetClusterRadarResponse = proto3.makeMessageType(
  "stats.GetClusterRadarResponse",
  () => [
    { no: 1, name: "radar", kind: "message", T: ClusterRadar },
  ],
);

/**
 * @generated from message stats.GetKoorClusterRequest
 */
export const GetKoorClusterRequest = proto3.makeMessageType(
  "stats.GetKoorClusterRequest",
  [],
);

/**
 * @generated from message stats.GetKoorClusterResponse
 */
export const GetKoorClusterResponse = proto3.makeMessageType(
  "stats.GetKoorClusterResponse",
  () => [
    { no: 1, name: "koor_cluster", kind: "message", T: KoorCluster },
  ],
);

