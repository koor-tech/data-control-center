// @generated by protoc-gen-es v1.6.0
// @generated from file api/services/ceph/v1/users.proto (package api.services.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { User } from "../../../resources/ceph/v1/users_pb.js";

/**
 * @generated from message api.services.ceph.v1.ListCephUsersRequest
 */
export declare class ListCephUsersRequest extends Message<ListCephUsersRequest> {
  constructor(data?: PartialMessage<ListCephUsersRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.ListCephUsersRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCephUsersRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCephUsersRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCephUsersRequest;

  static equals(a: ListCephUsersRequest | PlainMessage<ListCephUsersRequest> | undefined, b: ListCephUsersRequest | PlainMessage<ListCephUsersRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.ListCephUsersResponse
 */
export declare class ListCephUsersResponse extends Message<ListCephUsersResponse> {
  /**
   * @generated from field: repeated api.resources.ceph.v1.User ceph_users = 1;
   */
  cephUsers: User[];

  constructor(data?: PartialMessage<ListCephUsersResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.ListCephUsersResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCephUsersResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCephUsersResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCephUsersResponse;

  static equals(a: ListCephUsersResponse | PlainMessage<ListCephUsersResponse> | undefined, b: ListCephUsersResponse | PlainMessage<ListCephUsersResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.CreateCephUserRequest
 */
export declare class CreateCephUserRequest extends Message<CreateCephUserRequest> {
  /**
   * @generated from field: api.resources.ceph.v1.User ceph_user = 1;
   */
  cephUser?: User;

  constructor(data?: PartialMessage<CreateCephUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.CreateCephUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateCephUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateCephUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateCephUserRequest;

  static equals(a: CreateCephUserRequest | PlainMessage<CreateCephUserRequest> | undefined, b: CreateCephUserRequest | PlainMessage<CreateCephUserRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.CreateCephUserResponse
 */
export declare class CreateCephUserResponse extends Message<CreateCephUserResponse> {
  /**
   * @generated from field: api.resources.ceph.v1.User ceph_user = 1;
   */
  cephUser?: User;

  constructor(data?: PartialMessage<CreateCephUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.CreateCephUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateCephUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateCephUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateCephUserResponse;

  static equals(a: CreateCephUserResponse | PlainMessage<CreateCephUserResponse> | undefined, b: CreateCephUserResponse | PlainMessage<CreateCephUserResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.DeleteCephUserRequest
 */
export declare class DeleteCephUserRequest extends Message<DeleteCephUserRequest> {
  /**
   * @generated from field: string username = 1;
   */
  username: string;

  constructor(data?: PartialMessage<DeleteCephUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.DeleteCephUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteCephUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteCephUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteCephUserRequest;

  static equals(a: DeleteCephUserRequest | PlainMessage<DeleteCephUserRequest> | undefined, b: DeleteCephUserRequest | PlainMessage<DeleteCephUserRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.ceph.v1.DeleteCephUserResponse
 */
export declare class DeleteCephUserResponse extends Message<DeleteCephUserResponse> {
  /**
   * @generated from field: string status = 1;
   */
  status: string;

  constructor(data?: PartialMessage<DeleteCephUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.ceph.v1.DeleteCephUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteCephUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteCephUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteCephUserResponse;

  static equals(a: DeleteCephUserResponse | PlainMessage<DeleteCephUserResponse> | undefined, b: DeleteCephUserResponse | PlainMessage<DeleteCephUserResponse> | undefined): boolean;
}
