// @generated by protoc-gen-es v1.4.1
// @generated from file api/resources/ceph/v1/ceph.proto (package api.resources.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message api.resources.ceph.v1.User
 */
export declare class User extends Message<User> {
  /**
   * @generated from field: string Username = 1;
   */
  Username: string;

  /**
   * @generated from field: string Name = 2;
   */
  Name: string;

  /**
   * @generated from field: string Email = 3;
   */
  Email: string;

  /**
   * @generated from field: string Password = 4;
   */
  Password: string;

  /**
   * @generated from field: bool Enabled = 5;
   */
  Enabled: boolean;

  /**
   * @generated from field: repeated string Roles = 6;
   */
  Roles: string[];

  constructor(data?: PartialMessage<User>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.User";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User;

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean;
}
