// @generated by protoc-gen-connect-es v0.13.1
// @generated from file api/services/stats/v1/stats.proto (package api.services.stats.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetClusterNodesRequest, GetClusterNodesResponse, GetClusterRadarRequest, GetClusterRadarResponse, GetClusterResourcesRequest, GetClusterResourcesResponse, GetClusterStatsRequest, GetClusterStatsResponse } from "./stats_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.services.stats.v1.StatsService
 */
export declare const StatsService: {
  readonly typeName: "api.services.stats.v1.StatsService",
  readonly methods: {
    /**
     * @generated from rpc api.services.stats.v1.StatsService.GetClusterStats
     */
    readonly getClusterStats: {
      readonly name: "GetClusterStats",
      readonly I: typeof GetClusterStatsRequest,
      readonly O: typeof GetClusterStatsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.stats.v1.StatsService.GetClusterResources
     */
    readonly getClusterResources: {
      readonly name: "GetClusterResources",
      readonly I: typeof GetClusterResourcesRequest,
      readonly O: typeof GetClusterResourcesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.stats.v1.StatsService.GetClusterNodes
     */
    readonly getClusterNodes: {
      readonly name: "GetClusterNodes",
      readonly I: typeof GetClusterNodesRequest,
      readonly O: typeof GetClusterNodesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.stats.v1.StatsService.GetClusterRadar
     */
    readonly getClusterRadar: {
      readonly name: "GetClusterRadar",
      readonly I: typeof GetClusterRadarRequest,
      readonly O: typeof GetClusterRadarResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

