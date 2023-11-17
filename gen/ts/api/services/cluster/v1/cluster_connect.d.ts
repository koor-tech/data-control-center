// @generated by protoc-gen-connect-es v0.13.1
// @generated from file api/services/cluster/v1/cluster.proto (package api.services.cluster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CancelNetworkTestRequest, CancelNetworkTestResponse, GetKoorClusterRequest, GetKoorClusterResponse, GetNetworkTestResultsRequest, GetNetworkTestResultsResponse, GetNetworkTestStatusRequest, GetNetworkTestStatusResponse, GetResourcesRequest, GetResourcesResponse, GetTroubleshootReportRequest, GetTroubleshootReportResponse, SaveResourcesRequest, SaveResourcesResponse, SetScrubbingScheduleRequest, SetScrubbingScheduleResponse, StartNetworkTestRequest, StartNetworkTestResponse } from "./cluster_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.services.cluster.v1.ClusterService
 */
export declare const ClusterService: {
  readonly typeName: "api.services.cluster.v1.ClusterService",
  readonly methods: {
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.GetKoorCluster
     */
    readonly getKoorCluster: {
      readonly name: "GetKoorCluster",
      readonly I: typeof GetKoorClusterRequest,
      readonly O: typeof GetKoorClusterResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.GetTroubleshootReport
     */
    readonly getTroubleshootReport: {
      readonly name: "GetTroubleshootReport",
      readonly I: typeof GetTroubleshootReportRequest,
      readonly O: typeof GetTroubleshootReportResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.GetNetworkTestStatus
     */
    readonly getNetworkTestStatus: {
      readonly name: "GetNetworkTestStatus",
      readonly I: typeof GetNetworkTestStatusRequest,
      readonly O: typeof GetNetworkTestStatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.StartNetworkTest
     */
    readonly startNetworkTest: {
      readonly name: "StartNetworkTest",
      readonly I: typeof StartNetworkTestRequest,
      readonly O: typeof StartNetworkTestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.CancelNetworkTest
     */
    readonly cancelNetworkTest: {
      readonly name: "CancelNetworkTest",
      readonly I: typeof CancelNetworkTestRequest,
      readonly O: typeof CancelNetworkTestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.GetNetworkTestResults
     */
    readonly getNetworkTestResults: {
      readonly name: "GetNetworkTestResults",
      readonly I: typeof GetNetworkTestResultsRequest,
      readonly O: typeof GetNetworkTestResultsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.SetScrubbingSchedule
     */
    readonly setScrubbingSchedule: {
      readonly name: "SetScrubbingSchedule",
      readonly I: typeof SetScrubbingScheduleRequest,
      readonly O: typeof SetScrubbingScheduleResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.GetResources
     */
    readonly getResources: {
      readonly name: "GetResources",
      readonly I: typeof GetResourcesRequest,
      readonly O: typeof GetResourcesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.SaveResources
     */
    readonly saveResources: {
      readonly name: "SaveResources",
      readonly I: typeof SaveResourcesRequest,
      readonly O: typeof SaveResourcesResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

