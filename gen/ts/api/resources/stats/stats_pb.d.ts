// @generated by protoc-gen-es v1.3.1
// @generated from file api/resources/stats/stats.proto (package resources.stats, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum resources.stats.ResourceStatus
 */
export declare enum ResourceStatus {
  /**
   * @generated from enum value: RESOURCE_UNKNOWN = 0;
   */
  RESOURCE_UNKNOWN = 0,

  /**
   * @generated from enum value: RESOURCE_READY = 1;
   */
  RESOURCE_READY = 1,

  /**
   * @generated from enum value: RESOURCE_NOT_READY = 2;
   */
  RESOURCE_NOT_READY = 2,
}

/**
 * @generated from message resources.stats.MonService
 */
export declare class MonService extends Message<MonService> {
  /**
   * @generated from field: int32 daemon_count = 1;
   */
  daemonCount: number;

  /**
   * @generated from field: string quorum = 2;
   */
  quorum: string;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 3;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 4;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<MonService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.MonService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MonService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MonService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MonService;

  static equals(a: MonService | PlainMessage<MonService> | undefined, b: MonService | PlainMessage<MonService> | undefined): boolean;
}

/**
 * @generated from message resources.stats.MgrService
 */
export declare class MgrService extends Message<MgrService> {
  /**
   * @generated from field: string active = 1;
   */
  active: string;

  /**
   * @generated from field: repeated string standbys = 2;
   */
  standbys: string[];

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 4;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<MgrService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.MgrService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MgrService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MgrService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MgrService;

  static equals(a: MgrService | PlainMessage<MgrService> | undefined, b: MgrService | PlainMessage<MgrService> | undefined): boolean;
}

/**
 * @generated from message resources.stats.MdsService
 */
export declare class MdsService extends Message<MdsService> {
  /**
   * @generated from field: int32 daemons_up = 1;
   */
  daemonsUp: number;

  /**
   * @generated from field: int32 hot_standby_count = 2;
   */
  hotStandbyCount: number;

  constructor(data?: PartialMessage<MdsService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.MdsService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MdsService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MdsService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MdsService;

  static equals(a: MdsService | PlainMessage<MdsService> | undefined, b: MdsService | PlainMessage<MdsService> | undefined): boolean;
}

/**
 * @generated from message resources.stats.OsdService
 */
export declare class OsdService extends Message<OsdService> {
  /**
   * @generated from field: int32 osd_count = 1;
   */
  osdCount: number;

  /**
   * @generated from field: int32 osd_up = 2;
   */
  osdUp: number;

  /**
   * @generated from field: int32 osd_in = 3;
   */
  osdIn: number;

  /**
   * @generated from field: google.protobuf.Timestamp osd_up_updated_at = 4;
   */
  osdUpUpdatedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp osd_in_updated_at = 5;
   */
  osdInUpdatedAt?: Timestamp;

  constructor(data?: PartialMessage<OsdService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.OsdService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OsdService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OsdService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OsdService;

  static equals(a: OsdService | PlainMessage<OsdService> | undefined, b: OsdService | PlainMessage<OsdService> | undefined): boolean;
}

/**
 * @generated from message resources.stats.RgwService
 */
export declare class RgwService extends Message<RgwService> {
  /**
   * @generated from field: int32 active_daemon = 1;
   */
  activeDaemon: number;

  /**
   * @generated from field: int32 host_count = 2;
   */
  hostCount: number;

  /**
   * @generated from field: int32 zone_count = 3;
   */
  zoneCount: number;

  constructor(data?: PartialMessage<RgwService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.RgwService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RgwService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RgwService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RgwService;

  static equals(a: RgwService | PlainMessage<RgwService> | undefined, b: RgwService | PlainMessage<RgwService> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Services
 */
export declare class Services extends Message<Services> {
  /**
   * @generated from field: resources.stats.MonService mon = 1;
   */
  mon?: MonService;

  /**
   * @generated from field: resources.stats.MgrService mgr = 2;
   */
  mgr?: MgrService;

  /**
   * @generated from field: resources.stats.MdsService mds = 3;
   */
  mds?: MdsService;

  /**
   * @generated from field: resources.stats.OsdService osd = 4;
   */
  osd?: OsdService;

  /**
   * @generated from field: resources.stats.RgwService rgw = 5;
   */
  rgw?: RgwService;

  constructor(data?: PartialMessage<Services>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Services";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Services;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Services;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Services;

  static equals(a: Services | PlainMessage<Services> | undefined, b: Services | PlainMessage<Services> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Pgs
 */
export declare class Pgs extends Message<Pgs> {
  /**
   * @generated from field: int32 active_clean = 1;
   */
  activeClean: number;

  constructor(data?: PartialMessage<Pgs>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Pgs";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pgs;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pgs;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pgs;

  static equals(a: Pgs | PlainMessage<Pgs> | undefined, b: Pgs | PlainMessage<Pgs> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Pools
 */
export declare class Pools extends Message<Pools> {
  /**
   * @generated from field: int32 pools = 1;
   */
  pools: number;

  /**
   * @generated from field: resources.stats.Pgs pgs = 2;
   */
  pgs?: Pgs;

  constructor(data?: PartialMessage<Pools>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Pools";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pools;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pools;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pools;

  static equals(a: Pools | PlainMessage<Pools> | undefined, b: Pools | PlainMessage<Pools> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Objects
 */
export declare class Objects extends Message<Objects> {
  /**
   * @generated from field: int32 object_count = 1;
   */
  objectCount: number;

  /**
   * @generated from field: string size = 2;
   */
  size: string;

  constructor(data?: PartialMessage<Objects>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Objects";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Objects;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Objects;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Objects;

  static equals(a: Objects | PlainMessage<Objects> | undefined, b: Objects | PlainMessage<Objects> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Usage
 */
export declare class Usage extends Message<Usage> {
  /**
   * @generated from field: int64 used = 1;
   */
  used: bigint;

  /**
   * @generated from field: int64 available = 2;
   */
  available: bigint;

  /**
   * @generated from field: int64 total = 3;
   */
  total: bigint;

  constructor(data?: PartialMessage<Usage>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Usage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Usage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Usage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Usage;

  static equals(a: Usage | PlainMessage<Usage> | undefined, b: Usage | PlainMessage<Usage> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Data
 */
export declare class Data extends Message<Data> {
  /**
   * @generated from field: int32 volumes = 1;
   */
  volumes: number;

  /**
   * @generated from field: resources.stats.Pools pools = 2;
   */
  pools?: Pools;

  /**
   * @generated from field: resources.stats.Objects objects = 3;
   */
  objects?: Objects;

  /**
   * @generated from field: resources.stats.Usage usage = 4;
   */
  usage?: Usage;

  constructor(data?: PartialMessage<Data>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Data";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Data;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Data;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Data;

  static equals(a: Data | PlainMessage<Data> | undefined, b: Data | PlainMessage<Data> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Io
 */
export declare class Io extends Message<Io> {
  /**
   * @generated from field: string client_read = 1;
   */
  clientRead: string;

  /**
   * @generated from field: string client_read_ops = 2;
   */
  clientReadOps: string;

  /**
   * @generated from field: string client_write_ops = 3;
   */
  clientWriteOps: string;

  constructor(data?: PartialMessage<Io>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Io";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Io;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Io;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Io;

  static equals(a: Io | PlainMessage<Io> | undefined, b: Io | PlainMessage<Io> | undefined): boolean;
}

/**
 * @generated from message resources.stats.Crash
 */
export declare class Crash extends Message<Crash> {
  /**
   * @generated from field: string description = 1;
   */
  description: string;

  constructor(data?: PartialMessage<Crash>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.Crash";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Crash;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Crash;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Crash;

  static equals(a: Crash | PlainMessage<Crash> | undefined, b: Crash | PlainMessage<Crash> | undefined): boolean;
}

/**
 * @generated from message resources.stats.ClusterStats
 */
export declare class ClusterStats extends Message<ClusterStats> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string status = 2;
   */
  status: string;

  /**
   * @generated from field: repeated resources.stats.Crash crashes = 3;
   */
  crashes: Crash[];

  /**
   * @generated from field: resources.stats.Services services = 4;
   */
  services?: Services;

  /**
   * @generated from field: resources.stats.Data data = 5;
   */
  data?: Data;

  /**
   * @generated from field: string test = 6;
   */
  test: string;

  /**
   * @generated from field: resources.stats.Io io = 7;
   */
  io?: Io;

  constructor(data?: PartialMessage<ClusterStats>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.ClusterStats";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStats;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStats;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStats;

  static equals(a: ClusterStats | PlainMessage<ClusterStats> | undefined, b: ClusterStats | PlainMessage<ClusterStats> | undefined): boolean;
}

/**
 * @generated from message resources.stats.ResourceInfo
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
   * @generated from field: resources.stats.ResourceStatus status = 5;
   */
  status: ResourceStatus;

  constructor(data?: PartialMessage<ResourceInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.ResourceInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResourceInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResourceInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResourceInfo;

  static equals(a: ResourceInfo | PlainMessage<ResourceInfo> | undefined, b: ResourceInfo | PlainMessage<ResourceInfo> | undefined): boolean;
}

/**
 * @generated from message resources.stats.NodeInfo
 */
export declare class NodeInfo extends Message<NodeInfo> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: resources.stats.ResourceStatus status = 2;
   */
  status: ResourceStatus;

  /**
   * @generated from field: repeated string roles = 3;
   */
  roles: string[];

  /**
   * @generated from field: string internal_ip = 4;
   */
  internalIp: string;

  /**
   * @generated from field: string external_ip = 5;
   */
  externalIp: string;

  /**
   * @generated from field: google.protobuf.Timestamp age = 6;
   */
  age?: Timestamp;

  constructor(data?: PartialMessage<NodeInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.NodeInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeInfo;

  static equals(a: NodeInfo | PlainMessage<NodeInfo> | undefined, b: NodeInfo | PlainMessage<NodeInfo> | undefined): boolean;
}

/**
 * @generated from message resources.stats.ClusterRadar
 */
export declare class ClusterRadar extends Message<ClusterRadar> {
  /**
   * @generated from field: float cluster_health = 1;
   */
  clusterHealth: number;

  /**
   * @generated from field: float nodes_health = 2;
   */
  nodesHealth: number;

  /**
   * @generated from field: float capacity_used = 3;
   */
  capacityUsed: number;

  /**
   * @generated from field: float stability = 4;
   */
  stability: number;

  /**
   * @generated from field: float reliability = 5;
   */
  reliability: number;

  constructor(data?: PartialMessage<ClusterRadar>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "resources.stats.ClusterRadar";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterRadar;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterRadar;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterRadar;

  static equals(a: ClusterRadar | PlainMessage<ClusterRadar> | undefined, b: ClusterRadar | PlainMessage<ClusterRadar> | undefined): boolean;
}
