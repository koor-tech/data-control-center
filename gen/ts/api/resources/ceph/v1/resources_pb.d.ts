// @generated by protoc-gen-es v1.5.1
// @generated from file api/resources/ceph/v1/resources.proto (package api.resources.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message api.resources.ceph.v1.Resource
 */
export declare class Resource extends Message<Resource> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string content = 2;
   */
  content: string;

  /**
   * @generated from field: string namespace = 3;
   */
  namespace: string;

  /**
   * @generated from field: string kind = 4;
   */
  kind: string;

  /**
   * @generated from field: string object = 5;
   */
  object: string;

  constructor(data?: PartialMessage<Resource>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.Resource";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Resource;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Resource;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Resource;

  static equals(a: Resource | PlainMessage<Resource> | undefined, b: Resource | PlainMessage<Resource> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.Resources
 */
export declare class Resources extends Message<Resources> {
  /**
   * @generated from field: repeated api.resources.ceph.v1.Resource resources = 1;
   */
  resources: Resource[];

  constructor(data?: PartialMessage<Resources>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.Resources";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Resources;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Resources;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Resources;

  static equals(a: Resources | PlainMessage<Resources> | undefined, b: Resources | PlainMessage<Resources> | undefined): boolean;
}

