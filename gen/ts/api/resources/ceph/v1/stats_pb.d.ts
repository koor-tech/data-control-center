// @generated by protoc-gen-es v1.7.2
// @generated from file api/resources/ceph/v1/stats.proto (package api.resources.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum api.resources.ceph.v1.ClusterHealth
 */
export declare enum ClusterHealth {
  /**
   * @generated from enum value: CLUSTER_HEALTH_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CLUSTER_HEALTH_OFFLINE = 1;
   */
  OFFLINE = 1,

  /**
   * @generated from enum value: CLUSTER_HEALTH_OK = 2;
   */
  OK = 2,

  /**
   * @generated from enum value: CLUSTER_HEALTH_WARN = 3;
   */
  WARN = 3,

  /**
   * @generated from enum value: CLUSTER_HEALTH_ERR = 4;
   */
  ERR = 4,
}

/**
 * @generated from message api.resources.ceph.v1.MonService
 */
export declare class MonService extends Message<MonService> {
  /**
   * @generated from field: int32 daemon_count = 1;
   */
  daemonCount: number;

  /**
   * @generated from field: repeated string quorum = 2;
   */
  quorum: string[];

  /**
   * @generated from field: google.protobuf.Timestamp created_since = 5;
   */
  createdSince?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_since = 6;
   */
  updatedSince?: Timestamp;

  constructor(data?: PartialMessage<MonService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.MonService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MonService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MonService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MonService;

  static equals(a: MonService | PlainMessage<MonService> | undefined, b: MonService | PlainMessage<MonService> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.MgrService
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
   * @generated from field: google.protobuf.Timestamp updated_since = 6;
   */
  updatedSince?: Timestamp;

  constructor(data?: PartialMessage<MgrService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.MgrService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MgrService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MgrService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MgrService;

  static equals(a: MgrService | PlainMessage<MgrService> | undefined, b: MgrService | PlainMessage<MgrService> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.MdsService
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
  static readonly typeName = "api.resources.ceph.v1.MdsService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MdsService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MdsService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MdsService;

  static equals(a: MdsService | PlainMessage<MdsService> | undefined, b: MdsService | PlainMessage<MdsService> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.OsdService
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
   * @generated from field: google.protobuf.Timestamp osd_up_updated_since = 6;
   */
  osdUpUpdatedSince?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp osd_in_updated_since = 7;
   */
  osdInUpdatedSince?: Timestamp;

  constructor(data?: PartialMessage<OsdService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.OsdService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OsdService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OsdService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OsdService;

  static equals(a: OsdService | PlainMessage<OsdService> | undefined, b: OsdService | PlainMessage<OsdService> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.RgwService
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
  static readonly typeName = "api.resources.ceph.v1.RgwService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RgwService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RgwService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RgwService;

  static equals(a: RgwService | PlainMessage<RgwService> | undefined, b: RgwService | PlainMessage<RgwService> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.Services
 */
export declare class Services extends Message<Services> {
  /**
   * @generated from field: api.resources.ceph.v1.MonService mon = 1;
   */
  mon?: MonService;

  /**
   * @generated from field: api.resources.ceph.v1.MgrService mgr = 2;
   */
  mgr?: MgrService;

  /**
   * @generated from field: api.resources.ceph.v1.MdsService mds = 3;
   */
  mds?: MdsService;

  /**
   * @generated from field: api.resources.ceph.v1.OsdService osd = 4;
   */
  osd?: OsdService;

  /**
   * @generated from field: api.resources.ceph.v1.RgwService rgw = 5;
   */
  rgw?: RgwService;

  constructor(data?: PartialMessage<Services>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.Services";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Services;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Services;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Services;

  static equals(a: Services | PlainMessage<Services> | undefined, b: Services | PlainMessage<Services> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.PGs
 */
export declare class PGs extends Message<PGs> {
  /**
   * @generated from field: int32 active_clean = 1;
   */
  activeClean: number;

  constructor(data?: PartialMessage<PGs>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.PGs";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PGs;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PGs;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PGs;

  static equals(a: PGs | PlainMessage<PGs> | undefined, b: PGs | PlainMessage<PGs> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.Pools
 */
export declare class Pools extends Message<Pools> {
  /**
   * @generated from field: int32 pools = 1;
   */
  pools: number;

  /**
   * @generated from field: api.resources.ceph.v1.PGs pgs = 2;
   */
  pgs?: PGs;

  constructor(data?: PartialMessage<Pools>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.Pools";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pools;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pools;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pools;

  static equals(a: Pools | PlainMessage<Pools> | undefined, b: Pools | PlainMessage<Pools> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.Objects
 */
export declare class Objects extends Message<Objects> {
  /**
   * @generated from field: int32 object_count = 1;
   */
  objectCount: number;

  /**
   * @generated from field: int64 size = 2;
   */
  size: bigint;

  constructor(data?: PartialMessage<Objects>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.Objects";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Objects;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Objects;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Objects;

  static equals(a: Objects | PlainMessage<Objects> | undefined, b: Objects | PlainMessage<Objects> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.Usage
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
  static readonly typeName = "api.resources.ceph.v1.Usage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Usage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Usage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Usage;

  static equals(a: Usage | PlainMessage<Usage> | undefined, b: Usage | PlainMessage<Usage> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.Data
 */
export declare class Data extends Message<Data> {
  /**
   * @generated from field: int32 volumes = 1;
   */
  volumes: number;

  /**
   * @generated from field: api.resources.ceph.v1.Pools pools = 2;
   */
  pools?: Pools;

  /**
   * @generated from field: api.resources.ceph.v1.Objects objects = 3;
   */
  objects?: Objects;

  /**
   * @generated from field: api.resources.ceph.v1.Usage usage = 4;
   */
  usage?: Usage;

  constructor(data?: PartialMessage<Data>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.Data";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Data;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Data;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Data;

  static equals(a: Data | PlainMessage<Data> | undefined, b: Data | PlainMessage<Data> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.IOPS
 */
export declare class IOPS extends Message<IOPS> {
  /**
   * @generated from field: int64 client_read = 1;
   */
  clientRead: bigint;

  /**
   * @generated from field: int64 client_write = 2;
   */
  clientWrite: bigint;

  /**
   * @generated from field: int64 client_read_ops = 3;
   */
  clientReadOps: bigint;

  /**
   * @generated from field: int64 client_write_ops = 4;
   */
  clientWriteOps: bigint;

  constructor(data?: PartialMessage<IOPS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.IOPS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IOPS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IOPS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IOPS;

  static equals(a: IOPS | PlainMessage<IOPS> | undefined, b: IOPS | PlainMessage<IOPS> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.Crash
 */
export declare class Crash extends Message<Crash> {
  /**
   * @generated from field: string description = 1;
   */
  description: string;

  constructor(data?: PartialMessage<Crash>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.Crash";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Crash;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Crash;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Crash;

  static equals(a: Crash | PlainMessage<Crash> | undefined, b: Crash | PlainMessage<Crash> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.ClusterStats
 */
export declare class ClusterStats extends Message<ClusterStats> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: api.resources.ceph.v1.ClusterHealth status = 2;
   */
  status: ClusterHealth;

  /**
   * @generated from field: repeated api.resources.ceph.v1.Crash crashes = 3;
   */
  crashes: Crash[];

  /**
   * @generated from field: api.resources.ceph.v1.Services services = 4;
   */
  services?: Services;

  /**
   * @generated from field: api.resources.ceph.v1.Data data = 5;
   */
  data?: Data;

  /**
   * @generated from field: string test = 6;
   */
  test: string;

  /**
   * @generated from field: api.resources.ceph.v1.IOPS iops = 7;
   */
  iops?: IOPS;

  constructor(data?: PartialMessage<ClusterStats>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.ceph.v1.ClusterStats";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStats;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStats;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStats;

  static equals(a: ClusterStats | PlainMessage<ClusterStats> | undefined, b: ClusterStats | PlainMessage<ClusterStats> | undefined): boolean;
}

/**
 * @generated from message api.resources.ceph.v1.ClusterRadar
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
   * @generated from field: float capacity_available = 3;
   */
  capacityAvailable: number;

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
  static readonly typeName = "api.resources.ceph.v1.ClusterRadar";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterRadar;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterRadar;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterRadar;

  static equals(a: ClusterRadar | PlainMessage<ClusterRadar> | undefined, b: ClusterRadar | PlainMessage<ClusterRadar> | undefined): boolean;
}

