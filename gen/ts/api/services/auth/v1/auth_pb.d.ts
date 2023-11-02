// @generated by protoc-gen-es v1.4.1
// @generated from file api/services/auth/v1/auth.proto (package api.services.auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Timestamp } from "../../../resources/timestamp/v1/timestamp_pb.js";

/**
 * @generated from message api.services.auth.v1.LoginRequest
 */
export declare class LoginRequest extends Message<LoginRequest> {
  /**
   * @generated from field: string username = 1;
   */
  username: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;

  constructor(data?: PartialMessage<LoginRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.auth.v1.LoginRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginRequest;

  static equals(a: LoginRequest | PlainMessage<LoginRequest> | undefined, b: LoginRequest | PlainMessage<LoginRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.auth.v1.LoginResponse
 */
export declare class LoginResponse extends Message<LoginResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  /**
   * @generated from field: api.resources.timestamp.v1.Timestamp expires = 2;
   */
  expires?: Timestamp;

  /**
   * @generated from field: string account_id = 3;
   */
  accountId: string;

  constructor(data?: PartialMessage<LoginResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.auth.v1.LoginResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginResponse;

  static equals(a: LoginResponse | PlainMessage<LoginResponse> | undefined, b: LoginResponse | PlainMessage<LoginResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.auth.v1.LogoutRequest
 */
export declare class LogoutRequest extends Message<LogoutRequest> {
  constructor(data?: PartialMessage<LogoutRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.auth.v1.LogoutRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutRequest;

  static equals(a: LogoutRequest | PlainMessage<LogoutRequest> | undefined, b: LogoutRequest | PlainMessage<LogoutRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.auth.v1.LogoutResponse
 */
export declare class LogoutResponse extends Message<LogoutResponse> {
  /**
   * @generated from field: bool success = 1;
   */
  success: boolean;

  constructor(data?: PartialMessage<LogoutResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.auth.v1.LogoutResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutResponse;

  static equals(a: LogoutResponse | PlainMessage<LogoutResponse> | undefined, b: LogoutResponse | PlainMessage<LogoutResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.auth.v1.CheckTokenRequest
 */
export declare class CheckTokenRequest extends Message<CheckTokenRequest> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  constructor(data?: PartialMessage<CheckTokenRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.auth.v1.CheckTokenRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckTokenRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckTokenRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckTokenRequest;

  static equals(a: CheckTokenRequest | PlainMessage<CheckTokenRequest> | undefined, b: CheckTokenRequest | PlainMessage<CheckTokenRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.auth.v1.CheckTokenResponse
 */
export declare class CheckTokenResponse extends Message<CheckTokenResponse> {
  /**
   * @generated from field: bool success = 1;
   */
  success: boolean;

  constructor(data?: PartialMessage<CheckTokenResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.auth.v1.CheckTokenResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckTokenResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckTokenResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckTokenResponse;

  static equals(a: CheckTokenResponse | PlainMessage<CheckTokenResponse> | undefined, b: CheckTokenResponse | PlainMessage<CheckTokenResponse> | undefined): boolean;
}

