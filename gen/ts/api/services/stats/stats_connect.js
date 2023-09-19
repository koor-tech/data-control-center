// @generated by protoc-gen-connect-es v0.13.1
// @generated from file api/services/stats/stats.proto (package stats, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetClusterNodesRequest, GetClusterNodesResponse, GetClusterRadarRequest, GetClusterRadarResponse, GetClusterResourcesRequest, GetClusterResourcesResponse, GetClusterStatsRequest, GetClusterStatsResponse, GetKoorClusterRequest, GetKoorClusterResponse } from "./stats_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service stats.StatsService
 */
export const StatsService = {
  typeName: "stats.StatsService",
  methods: {
    /**
     * @generated from rpc stats.StatsService.GetClusterStats
     */
    getClusterStats: {
      name: "GetClusterStats",
      I: GetClusterStatsRequest,
      O: GetClusterStatsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stats.StatsService.GetClusterResources
     */
    getClusterResources: {
      name: "GetClusterResources",
      I: GetClusterResourcesRequest,
      O: GetClusterResourcesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stats.StatsService.GetClusterNodes
     */
    getClusterNodes: {
      name: "GetClusterNodes",
      I: GetClusterNodesRequest,
      O: GetClusterNodesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stats.StatsService.GetClusterRadar
     */
    getClusterRadar: {
      name: "GetClusterRadar",
      I: GetClusterRadarRequest,
      O: GetClusterRadarResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc stats.StatsService.GetKoorCluster
     */
    getKoorCluster: {
      name: "GetKoorCluster",
      I: GetKoorClusterRequest,
      O: GetKoorClusterResponse,
      kind: MethodKind.Unary,
    },
  }
};

