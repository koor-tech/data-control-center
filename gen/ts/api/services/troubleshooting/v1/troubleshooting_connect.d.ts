// @generated by protoc-gen-connect-es v1.3.0
// @generated from file api/services/troubleshooting/v1/troubleshooting.proto (package api.services.troubleshooting.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CancelNetworkTestRequest, CancelNetworkTestResponse, GetNetworkTestResultsRequest, GetNetworkTestResultsResponse, GetNetworkTestStatusRequest, GetNetworkTestStatusResponse, GetTroubleshootReportRequest, GetTroubleshootReportResponse, StartNetworkTestRequest, StartNetworkTestResponse } from "./troubleshooting_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.services.troubleshooting.v1.TroubleshootingService
 */
export declare const TroubleshootingService: {
  readonly typeName: "api.services.troubleshooting.v1.TroubleshootingService",
  readonly methods: {
    /**
     * @generated from rpc api.services.troubleshooting.v1.TroubleshootingService.GetTroubleshootReport
     */
    readonly getTroubleshootReport: {
      readonly name: "GetTroubleshootReport",
      readonly I: typeof GetTroubleshootReportRequest,
      readonly O: typeof GetTroubleshootReportResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.troubleshooting.v1.TroubleshootingService.GetNetworkTestStatus
     */
    readonly getNetworkTestStatus: {
      readonly name: "GetNetworkTestStatus",
      readonly I: typeof GetNetworkTestStatusRequest,
      readonly O: typeof GetNetworkTestStatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.troubleshooting.v1.TroubleshootingService.StartNetworkTest
     */
    readonly startNetworkTest: {
      readonly name: "StartNetworkTest",
      readonly I: typeof StartNetworkTestRequest,
      readonly O: typeof StartNetworkTestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.troubleshooting.v1.TroubleshootingService.CancelNetworkTest
     */
    readonly cancelNetworkTest: {
      readonly name: "CancelNetworkTest",
      readonly I: typeof CancelNetworkTestRequest,
      readonly O: typeof CancelNetworkTestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.troubleshooting.v1.TroubleshootingService.GetNetworkTestResults
     */
    readonly getNetworkTestResults: {
      readonly name: "GetNetworkTestResults",
      readonly I: typeof GetNetworkTestResultsRequest,
      readonly O: typeof GetNetworkTestResultsResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

