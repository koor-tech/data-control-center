// @generated by protoc-gen-es v1.6.0
// @generated from file api/resources/k8s/v1/resources.proto (package api.resources.k8s.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum api.resources.k8s.v1.ReliabilityScore
 */
export declare enum ReliabilityScore {
  /**
   * @generated from enum value: RELIABILITY_SCORE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RELIABILITY_SCORE_UNKNOWN = 1;
   */
  UNKNOWN = 1,

  /**
   * @generated from enum value: RELIABILITY_SCORE_NONE = 2;
   */
  NONE = 2,

  /**
   * @generated from enum value: RELIABILITY_SCORE_DEGRADED = 3;
   */
  DEGRADED = 3,

  /**
   * @generated from enum value: RELIABILITY_SCORE_OK = 4;
   */
  OK = 4,
}

/**
 * @generated from enum api.resources.k8s.v1.ResourceStatus
 */
export declare enum ResourceStatus {
  /**
   * @generated from enum value: RESOURCE_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RESOURCE_STATUS_UNKNOWN = 1;
   */
  UNKNOWN = 1,

  /**
   * @generated from enum value: RESOURCE_STATUS_READY = 2;
   */
  READY = 2,

  /**
   * @generated from enum value: RESOURCE_STATUS_PROGRESSING = 3;
   */
  PROGRESSING = 3,

  /**
   * @generated from enum value: RESOURCE_STATUS_NOT_READY = 4;
   */
  NOT_READY = 4,
}

/**
 * @generated from message api.resources.k8s.v1.Resource
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
  static readonly typeName = "api.resources.k8s.v1.Resource";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Resource;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Resource;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Resource;

  static equals(a: Resource | PlainMessage<Resource> | undefined, b: Resource | PlainMessage<Resource> | undefined): boolean;
}

/**
 * @generated from message api.resources.k8s.v1.Resources
 */
export declare class Resources extends Message<Resources> {
  /**
   * @generated from field: repeated api.resources.k8s.v1.Resource resources = 1;
   */
  resources: Resource[];

  constructor(data?: PartialMessage<Resources>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.k8s.v1.Resources";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Resources;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Resources;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Resources;

  static equals(a: Resources | PlainMessage<Resources> | undefined, b: Resources | PlainMessage<Resources> | undefined): boolean;
}

/**
 * @generated from message api.resources.k8s.v1.ResourceInfo
 */
export declare class ResourceInfo extends Message<ResourceInfo> {
  /**
   * @generated from field: string apiversion = 1;
   */
  apiversion: string;

  /**
   * @generated from field: string kind = 2;
   */
  kind: string;

  /**
   * @generated from field: string namespace = 3;
   */
  namespace: string;

  /**
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * @generated from field: api.resources.k8s.v1.ResourceStatus status = 5;
   */
  status: ResourceStatus;

  /**
   * @generated from field: int32 replicas = 6;
   */
  replicas: number;

  /**
   * @generated from field: api.resources.k8s.v1.ReliabilityScore reliability = 7;
   */
  reliability: ReliabilityScore;

  /**
   * @generated from field: optional string version = 8;
   */
  version?: string;

  constructor(data?: PartialMessage<ResourceInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.k8s.v1.ResourceInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResourceInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResourceInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResourceInfo;

  static equals(a: ResourceInfo | PlainMessage<ResourceInfo> | undefined, b: ResourceInfo | PlainMessage<ResourceInfo> | undefined): boolean;
}

/**
 * @generated from message api.resources.k8s.v1.NodeInfo
 */
export declare class NodeInfo extends Message<NodeInfo> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: api.resources.k8s.v1.ResourceStatus status = 2;
   */
  status: ResourceStatus;

  /**
   * @generated from field: repeated string roles = 3;
   */
  roles: string[];

  /**
   * @generated from field: optional string internal_ip = 4;
   */
  internalIp?: string;

  /**
   * @generated from field: optional string external_ip = 5;
   */
  externalIp?: string;

  /**
   * @generated from field: optional google.protobuf.Timestamp age = 6;
   */
  age?: Timestamp;

  constructor(data?: PartialMessage<NodeInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.k8s.v1.NodeInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeInfo;

  static equals(a: NodeInfo | PlainMessage<NodeInfo> | undefined, b: NodeInfo | PlainMessage<NodeInfo> | undefined): boolean;
}

