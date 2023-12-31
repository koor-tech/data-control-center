// @generated by protoc-gen-es v1.3.0
// @generated from file vendor/google.golang.org/appengine/internal/remote_api/remote_api.proto (package remote_api, syntax proto2)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto2 } from "@bufbuild/protobuf";

/**
 * @generated from message remote_api.Request
 */
export declare class Request extends Message<Request> {
  /**
   * @generated from field: required string service_name = 2;
   */
  serviceName: string;

  /**
   * @generated from field: required string method = 3;
   */
  method: string;

  /**
   * @generated from field: required bytes request = 4;
   */
  request: Uint8Array;

  /**
   * @generated from field: optional string request_id = 5;
   */
  requestId?: string;

  constructor(data?: PartialMessage<Request>);

  static readonly runtime: typeof proto2;
  static readonly typeName = "remote_api.Request";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Request;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Request;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Request;

  static equals(a: Request | PlainMessage<Request> | undefined, b: Request | PlainMessage<Request> | undefined): boolean;
}

/**
 * @generated from message remote_api.ApplicationError
 */
export declare class ApplicationError extends Message<ApplicationError> {
  /**
   * @generated from field: required int32 code = 1;
   */
  code: number;

  /**
   * @generated from field: required string detail = 2;
   */
  detail: string;

  constructor(data?: PartialMessage<ApplicationError>);

  static readonly runtime: typeof proto2;
  static readonly typeName = "remote_api.ApplicationError";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApplicationError;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApplicationError;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApplicationError;

  static equals(a: ApplicationError | PlainMessage<ApplicationError> | undefined, b: ApplicationError | PlainMessage<ApplicationError> | undefined): boolean;
}

/**
 * @generated from message remote_api.RpcError
 */
export declare class RpcError extends Message<RpcError> {
  /**
   * @generated from field: required int32 code = 1;
   */
  code: number;

  /**
   * @generated from field: optional string detail = 2;
   */
  detail?: string;

  constructor(data?: PartialMessage<RpcError>);

  static readonly runtime: typeof proto2;
  static readonly typeName = "remote_api.RpcError";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RpcError;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RpcError;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RpcError;

  static equals(a: RpcError | PlainMessage<RpcError> | undefined, b: RpcError | PlainMessage<RpcError> | undefined): boolean;
}

/**
 * @generated from enum remote_api.RpcError.ErrorCode
 */
export declare enum RpcError_ErrorCode {
  /**
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * @generated from enum value: CALL_NOT_FOUND = 1;
   */
  CALL_NOT_FOUND = 1,

  /**
   * @generated from enum value: PARSE_ERROR = 2;
   */
  PARSE_ERROR = 2,

  /**
   * @generated from enum value: SECURITY_VIOLATION = 3;
   */
  SECURITY_VIOLATION = 3,

  /**
   * @generated from enum value: OVER_QUOTA = 4;
   */
  OVER_QUOTA = 4,

  /**
   * @generated from enum value: REQUEST_TOO_LARGE = 5;
   */
  REQUEST_TOO_LARGE = 5,

  /**
   * @generated from enum value: CAPABILITY_DISABLED = 6;
   */
  CAPABILITY_DISABLED = 6,

  /**
   * @generated from enum value: FEATURE_DISABLED = 7;
   */
  FEATURE_DISABLED = 7,

  /**
   * @generated from enum value: BAD_REQUEST = 8;
   */
  BAD_REQUEST = 8,

  /**
   * @generated from enum value: RESPONSE_TOO_LARGE = 9;
   */
  RESPONSE_TOO_LARGE = 9,

  /**
   * @generated from enum value: CANCELLED = 10;
   */
  CANCELLED = 10,

  /**
   * @generated from enum value: REPLAY_ERROR = 11;
   */
  REPLAY_ERROR = 11,

  /**
   * @generated from enum value: DEADLINE_EXCEEDED = 12;
   */
  DEADLINE_EXCEEDED = 12,
}

/**
 * @generated from message remote_api.Response
 */
export declare class Response extends Message<Response> {
  /**
   * @generated from field: optional bytes response = 1;
   */
  response?: Uint8Array;

  /**
   * @generated from field: optional bytes exception = 2;
   */
  exception?: Uint8Array;

  /**
   * @generated from field: optional remote_api.ApplicationError application_error = 3;
   */
  applicationError?: ApplicationError;

  /**
   * @generated from field: optional bytes java_exception = 4;
   */
  javaException?: Uint8Array;

  /**
   * @generated from field: optional remote_api.RpcError rpc_error = 5;
   */
  rpcError?: RpcError;

  constructor(data?: PartialMessage<Response>);

  static readonly runtime: typeof proto2;
  static readonly typeName = "remote_api.Response";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Response;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Response;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Response;

  static equals(a: Response | PlainMessage<Response> | undefined, b: Response | PlainMessage<Response> | undefined): boolean;
}

