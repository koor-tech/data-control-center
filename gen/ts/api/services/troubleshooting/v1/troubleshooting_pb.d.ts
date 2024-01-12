// @generated by protoc-gen-es v1.6.0
// @generated from file api/services/troubleshooting/v1/troubleshooting.proto (package api.services.troubleshooting.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum api.services.troubleshooting.v1.NetworkTestOutputFormat
 */
export declare enum NetworkTestOutputFormat {
  /**
   * @generated from enum value: NETWORK_TEST_OUTPUT_FORMAT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: NETWORK_TEST_OUTPUT_FORMAT_CSV = 1;
   */
  CSV = 1,

  /**
   * @generated from enum value: NETWORK_TEST_OUTPUT_FORMAT_EXCELIZE = 2;
   */
  EXCELIZE = 2,
}

/**
 * @generated from message api.services.troubleshooting.v1.GetTroubleshootReportRequest
 */
export declare class GetTroubleshootReportRequest extends Message<GetTroubleshootReportRequest> {
  constructor(data?: PartialMessage<GetTroubleshootReportRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.GetTroubleshootReportRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTroubleshootReportRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTroubleshootReportRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTroubleshootReportRequest;

  static equals(a: GetTroubleshootReportRequest | PlainMessage<GetTroubleshootReportRequest> | undefined, b: GetTroubleshootReportRequest | PlainMessage<GetTroubleshootReportRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.GetTroubleshootReportResponse
 */
export declare class GetTroubleshootReportResponse extends Message<GetTroubleshootReportResponse> {
  /**
   * @generated from field: string report = 1;
   */
  report: string;

  constructor(data?: PartialMessage<GetTroubleshootReportResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.GetTroubleshootReportResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTroubleshootReportResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTroubleshootReportResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTroubleshootReportResponse;

  static equals(a: GetTroubleshootReportResponse | PlainMessage<GetTroubleshootReportResponse> | undefined, b: GetTroubleshootReportResponse | PlainMessage<GetTroubleshootReportResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.GetNetworkTestStatusRequest
 */
export declare class GetNetworkTestStatusRequest extends Message<GetNetworkTestStatusRequest> {
  constructor(data?: PartialMessage<GetNetworkTestStatusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.GetNetworkTestStatusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNetworkTestStatusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNetworkTestStatusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNetworkTestStatusRequest;

  static equals(a: GetNetworkTestStatusRequest | PlainMessage<GetNetworkTestStatusRequest> | undefined, b: GetNetworkTestStatusRequest | PlainMessage<GetNetworkTestStatusRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.GetNetworkTestStatusResponse
 */
export declare class GetNetworkTestStatusResponse extends Message<GetNetworkTestStatusResponse> {
  /**
   * @generated from field: bool running = 1;
   */
  running: boolean;

  /**
   * @generated from field: bool complete = 2;
   */
  complete: boolean;

  /**
   * @generated from field: string logs = 3;
   */
  logs: string;

  constructor(data?: PartialMessage<GetNetworkTestStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.GetNetworkTestStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNetworkTestStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNetworkTestStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNetworkTestStatusResponse;

  static equals(a: GetNetworkTestStatusResponse | PlainMessage<GetNetworkTestStatusResponse> | undefined, b: GetNetworkTestStatusResponse | PlainMessage<GetNetworkTestStatusResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.StartNetworkTestRequest
 */
export declare class StartNetworkTestRequest extends Message<StartNetworkTestRequest> {
  /**
   * @generated from field: bool host_network = 1;
   */
  hostNetwork: boolean;

  /**
   * @generated from field: api.services.troubleshooting.v1.NetworkTestOutputFormat output_format = 2;
   */
  outputFormat: NetworkTestOutputFormat;

  constructor(data?: PartialMessage<StartNetworkTestRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.StartNetworkTestRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartNetworkTestRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartNetworkTestRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartNetworkTestRequest;

  static equals(a: StartNetworkTestRequest | PlainMessage<StartNetworkTestRequest> | undefined, b: StartNetworkTestRequest | PlainMessage<StartNetworkTestRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.StartNetworkTestResponse
 */
export declare class StartNetworkTestResponse extends Message<StartNetworkTestResponse> {
  /**
   * @generated from field: bool already_running = 1;
   */
  alreadyRunning: boolean;

  constructor(data?: PartialMessage<StartNetworkTestResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.StartNetworkTestResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartNetworkTestResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartNetworkTestResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartNetworkTestResponse;

  static equals(a: StartNetworkTestResponse | PlainMessage<StartNetworkTestResponse> | undefined, b: StartNetworkTestResponse | PlainMessage<StartNetworkTestResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.CancelNetworkTestRequest
 */
export declare class CancelNetworkTestRequest extends Message<CancelNetworkTestRequest> {
  constructor(data?: PartialMessage<CancelNetworkTestRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.CancelNetworkTestRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CancelNetworkTestRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CancelNetworkTestRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CancelNetworkTestRequest;

  static equals(a: CancelNetworkTestRequest | PlainMessage<CancelNetworkTestRequest> | undefined, b: CancelNetworkTestRequest | PlainMessage<CancelNetworkTestRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.CancelNetworkTestResponse
 */
export declare class CancelNetworkTestResponse extends Message<CancelNetworkTestResponse> {
  /**
   * @generated from field: bool not_running = 1;
   */
  notRunning: boolean;

  constructor(data?: PartialMessage<CancelNetworkTestResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.CancelNetworkTestResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CancelNetworkTestResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CancelNetworkTestResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CancelNetworkTestResponse;

  static equals(a: CancelNetworkTestResponse | PlainMessage<CancelNetworkTestResponse> | undefined, b: CancelNetworkTestResponse | PlainMessage<CancelNetworkTestResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.GetNetworkTestResultsRequest
 */
export declare class GetNetworkTestResultsRequest extends Message<GetNetworkTestResultsRequest> {
  constructor(data?: PartialMessage<GetNetworkTestResultsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.GetNetworkTestResultsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNetworkTestResultsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNetworkTestResultsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNetworkTestResultsRequest;

  static equals(a: GetNetworkTestResultsRequest | PlainMessage<GetNetworkTestResultsRequest> | undefined, b: GetNetworkTestResultsRequest | PlainMessage<GetNetworkTestResultsRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.troubleshooting.v1.GetNetworkTestResultsResponse
 */
export declare class GetNetworkTestResultsResponse extends Message<GetNetworkTestResultsResponse> {
  /**
   * @generated from field: string file_name = 1;
   */
  fileName: string;

  /**
   * @generated from field: string file_type = 2;
   */
  fileType: string;

  /**
   * @generated from field: bytes file_contents = 3;
   */
  fileContents: Uint8Array;

  constructor(data?: PartialMessage<GetNetworkTestResultsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.troubleshooting.v1.GetNetworkTestResultsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNetworkTestResultsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNetworkTestResultsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNetworkTestResultsResponse;

  static equals(a: GetNetworkTestResultsResponse | PlainMessage<GetNetworkTestResultsResponse> | undefined, b: GetNetworkTestResultsResponse | PlainMessage<GetNetworkTestResultsResponse> | undefined): boolean;
}

