// @generated by protoc-gen-es v1.4.1
// @generated from file api/services/ceph/v1/ceph.proto (package api.services.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { User } from "../../../resources/ceph/v1/ceph_pb.js";

/**
 * @generated from message api.services.ceph.v1.GetCephUsersRequest
 */
export declare class GetCephUsersRequest extends Message<GetCephUsersRequest> {
  constructor(data?: PartialMessage<GetCephUsersRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.GetCephUsersRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCephUsersRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCephUsersRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCephUsersRequest;

  static equals(a: GetCephUsersRequest | PlainMessage<GetCephUsersRequest> | undefined, b: GetCephUsersRequest | PlainMessage<GetCephUsersRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.GetCephUsersResponse
 */
export declare class GetCephUsersResponse extends Message<GetCephUsersResponse> {
  /**
   * @generated from field: repeated api.resources.ceph.v1.User cephUsers = 1;
   */
  cephUsers: User[];

  constructor(data?: PartialMessage<GetCephUsersResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.GetCephUsersResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCephUsersResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCephUsersResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCephUsersResponse;

  static equals(a: GetCephUsersResponse | PlainMessage<GetCephUsersResponse> | undefined, b: GetCephUsersResponse | PlainMessage<GetCephUsersResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.CreatCephUsersRequest
 */
export declare class CreatCephUsersRequest extends Message<CreatCephUsersRequest> {
  /**
   * @generated from field: api.resources.ceph.v1.User cephUser = 1;
   */
  cephUser?: User;

  constructor(data?: PartialMessage<CreatCephUsersRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.CreatCephUsersRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreatCephUsersRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreatCephUsersRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreatCephUsersRequest;

  static equals(a: CreatCephUsersRequest | PlainMessage<CreatCephUsersRequest> | undefined, b: CreatCephUsersRequest | PlainMessage<CreatCephUsersRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.CephUsersResponse
 */
export declare class CephUsersResponse extends Message<CephUsersResponse> {
  /**
   * @generated from field: api.resources.ceph.v1.User cephUser = 1;
   */
  cephUser?: User;

  constructor(data?: PartialMessage<CephUsersResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.CephUsersResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CephUsersResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CephUsersResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CephUsersResponse;

  static equals(a: CephUsersResponse | PlainMessage<CephUsersResponse> | undefined, b: CephUsersResponse | PlainMessage<CephUsersResponse> | undefined): boolean;
}

