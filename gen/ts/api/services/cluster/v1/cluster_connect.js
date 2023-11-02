// @generated by protoc-gen-connect-es v0.13.1
// @generated from file api/services/cluster/v1/cluster.proto (package api.services.cluster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetKoorClusterRequest, GetKoorClusterResponse, GetTroubleshootReportRequest, GetTroubleshootReportResponse, SetScrubbingScheduleRequest, SetScrubbingScheduleResponse } from "./cluster_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.services.cluster.v1.ClusterService
 */
export const ClusterService = {
  typeName: "api.services.cluster.v1.ClusterService",
  methods: {
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.GetKoorCluster
     */
    getKoorCluster: {
      name: "GetKoorCluster",
      I: GetKoorClusterRequest,
      O: GetKoorClusterResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.GetTroubleshootReport
     */
    getTroubleshootReport: {
      name: "GetTroubleshootReport",
      I: GetTroubleshootReportRequest,
      O: GetTroubleshootReportResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.cluster.v1.ClusterService.SetScrubbingSchedule
     */
    setScrubbingSchedule: {
      name: "SetScrubbingSchedule",
      I: SetScrubbingScheduleRequest,
      O: SetScrubbingScheduleResponse,
      kind: MethodKind.Unary,
    },
  }
};
