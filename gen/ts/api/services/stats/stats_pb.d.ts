// @generated by protoc-gen-es v1.3.0
// @generated from file api/services/stats/stats.proto (package stats, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum stats.HealthStatus
 */
export declare enum HealthStatus {
  /**
   * @generated from enum value: HEALTH_OK = 0;
   */
  HEALTH_OK = 0,

  /**
   * @generated from enum value: HEALTH_WARN = 1;
   */
  HEALTH_WARN = 1,

  /**
   * @generated from enum value: HEALTH_ERR = 2;
   */
  HEALTH_ERR = 2,
}

/**
 * @generated from message stats.DaemonCrash
 */
export declare class DaemonCrash extends Message<DaemonCrash> {
  /**
   * @generated from field: string description = 1;
   */
  description: string;

  constructor(data?: PartialMessage<DaemonCrash>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.DaemonCrash";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DaemonCrash;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DaemonCrash;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DaemonCrash;

  static equals(a: DaemonCrash | PlainMessage<DaemonCrash> | undefined, b: DaemonCrash | PlainMessage<DaemonCrash> | undefined): boolean;
}

/**
 * @generated from message stats.Daemon
 */
export declare class Daemon extends Message<Daemon> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string status = 2;
   */
  status: string;

  /**
   * @generated from field: string description = 3;
   */
  description: string;

  constructor(data?: PartialMessage<Daemon>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.Daemon";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Daemon;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Daemon;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Daemon;

  static equals(a: Daemon | PlainMessage<Daemon> | undefined, b: Daemon | PlainMessage<Daemon> | undefined): boolean;
}

/**
 * @generated from message stats.MonService
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
   * @generated from field: string age = 3;
   */
  age: string;

  constructor(data?: PartialMessage<MonService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.MonService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MonService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MonService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MonService;

  static equals(a: MonService | PlainMessage<MonService> | undefined, b: MonService | PlainMessage<MonService> | undefined): boolean;
}

/**
 * @generated from message stats.MgrService
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
   * @generated from field: string since = 3;
   */
  since: string;

  constructor(data?: PartialMessage<MgrService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.MgrService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MgrService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MgrService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MgrService;

  static equals(a: MgrService | PlainMessage<MgrService> | undefined, b: MgrService | PlainMessage<MgrService> | undefined): boolean;
}

/**
 * @generated from message stats.MdsService
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
  static readonly typeName = "stats.MdsService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MdsService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MdsService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MdsService;

  static equals(a: MdsService | PlainMessage<MdsService> | undefined, b: MdsService | PlainMessage<MdsService> | undefined): boolean;
}

/**
 * @generated from message stats.OsdService
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
   * @generated from field: string since = 4;
   */
  since: string;

  constructor(data?: PartialMessage<OsdService>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.OsdService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OsdService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OsdService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OsdService;

  static equals(a: OsdService | PlainMessage<OsdService> | undefined, b: OsdService | PlainMessage<OsdService> | undefined): boolean;
}

/**
 * @generated from message stats.RgwService
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
  static readonly typeName = "stats.RgwService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RgwService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RgwService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RgwService;

  static equals(a: RgwService | PlainMessage<RgwService> | undefined, b: RgwService | PlainMessage<RgwService> | undefined): boolean;
}

/**
 * @generated from message stats.VolumeStatus
 */
export declare class VolumeStatus extends Message<VolumeStatus> {
  /**
   * @generated from field: string description = 1;
   */
  description: string;

  constructor(data?: PartialMessage<VolumeStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.VolumeStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VolumeStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VolumeStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VolumeStatus;

  static equals(a: VolumeStatus | PlainMessage<VolumeStatus> | undefined, b: VolumeStatus | PlainMessage<VolumeStatus> | undefined): boolean;
}

/**
 * @generated from message stats.PoolStatus
 */
export declare class PoolStatus extends Message<PoolStatus> {
  /**
   * @generated from field: int32 pool_count = 1;
   */
  poolCount: number;

  /**
   * @generated from field: int32 pgs = 2;
   */
  pgs: number;

  constructor(data?: PartialMessage<PoolStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.PoolStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PoolStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PoolStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PoolStatus;

  static equals(a: PoolStatus | PlainMessage<PoolStatus> | undefined, b: PoolStatus | PlainMessage<PoolStatus> | undefined): boolean;
}

/**
 * @generated from message stats.ObjectStatus
 */
export declare class ObjectStatus extends Message<ObjectStatus> {
  /**
   * @generated from field: string object_count = 1;
   */
  objectCount: string;

  /**
   * @generated from field: string size = 2;
   */
  size: string;

  constructor(data?: PartialMessage<ObjectStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.ObjectStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ObjectStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ObjectStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ObjectStatus;

  static equals(a: ObjectStatus | PlainMessage<ObjectStatus> | undefined, b: ObjectStatus | PlainMessage<ObjectStatus> | undefined): boolean;
}

/**
 * @generated from message stats.UsageStatus
 */
export declare class UsageStatus extends Message<UsageStatus> {
  /**
   * @generated from field: float used = 1;
   */
  used: number;

  /**
   * @generated from field: float available = 2;
   */
  available: number;

  /**
   * @generated from field: float total = 3;
   */
  total: number;

  constructor(data?: PartialMessage<UsageStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.UsageStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UsageStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UsageStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UsageStatus;

  static equals(a: UsageStatus | PlainMessage<UsageStatus> | undefined, b: UsageStatus | PlainMessage<UsageStatus> | undefined): boolean;
}

/**
 * @generated from message stats.IoStatus
 */
export declare class IoStatus extends Message<IoStatus> {
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

  constructor(data?: PartialMessage<IoStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.IoStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IoStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IoStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IoStatus;

  static equals(a: IoStatus | PlainMessage<IoStatus> | undefined, b: IoStatus | PlainMessage<IoStatus> | undefined): boolean;
}

/**
 * @generated from message stats.PGs
 */
export declare class PGs extends Message<PGs> {
  /**
   * @generated from field: int32 active_clean = 1;
   */
  activeClean: number;

  constructor(data?: PartialMessage<PGs>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.PGs";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PGs;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PGs;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PGs;

  static equals(a: PGs | PlainMessage<PGs> | undefined, b: PGs | PlainMessage<PGs> | undefined): boolean;
}

/**
 * @generated from message stats.Services
 */
export declare class Services extends Message<Services> {
  /**
   * @generated from field: stats.MonService mon = 1;
   */
  mon?: MonService;

  /**
   * @generated from field: stats.MgrService mgr = 2;
   */
  mgr?: MgrService;

  /**
   * @generated from field: stats.MdsService mds = 3;
   */
  mds?: MdsService;

  /**
   * @generated from field: stats.OsdService osd = 4;
   */
  osd?: OsdService;

  /**
   * @generated from field: stats.RgwService rgw = 5;
   */
  rgw?: RgwService;

  constructor(data?: PartialMessage<Services>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.Services";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Services;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Services;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Services;

  static equals(a: Services | PlainMessage<Services> | undefined, b: Services | PlainMessage<Services> | undefined): boolean;
}

/**
 * @generated from message stats.Data
 */
export declare class Data extends Message<Data> {
  /**
   * @generated from field: stats.VolumeStatus volumes = 1;
   */
  volumes?: VolumeStatus;

  /**
   * @generated from field: stats.PoolStatus pools = 2;
   */
  pools?: PoolStatus;

  /**
   * @generated from field: stats.ObjectStatus objects = 3;
   */
  objects?: ObjectStatus;

  /**
   * @generated from field: stats.UsageStatus usage = 4;
   */
  usage?: UsageStatus;

  /**
   * @generated from field: stats.PGs pgs = 5;
   */
  pgs?: PGs;

  constructor(data?: PartialMessage<Data>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.Data";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Data;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Data;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Data;

  static equals(a: Data | PlainMessage<Data> | undefined, b: Data | PlainMessage<Data> | undefined): boolean;
}

/**
 * @generated from message stats.ClusterStatusResponse
 */
export declare class ClusterStatusResponse extends Message<ClusterStatusResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: stats.HealthStatus health = 2;
   */
  health: HealthStatus;

  /**
   * @generated from field: repeated stats.DaemonCrash daemon_crashes = 3;
   */
  daemonCrashes: DaemonCrash[];

  /**
   * @generated from field: stats.Services services = 4;
   */
  services?: Services;

  /**
   * @generated from field: stats.Data data = 5;
   */
  data?: Data;

  /**
   * @generated from field: stats.IoStatus io = 6;
   */
  io?: IoStatus;

  constructor(data?: PartialMessage<ClusterStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.ClusterStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStatusResponse;

  static equals(a: ClusterStatusResponse | PlainMessage<ClusterStatusResponse> | undefined, b: ClusterStatusResponse | PlainMessage<ClusterStatusResponse> | undefined): boolean;
}

/**
 * @generated from message stats.EmptyRequest
 */
export declare class EmptyRequest extends Message<EmptyRequest> {
  constructor(data?: PartialMessage<EmptyRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "stats.EmptyRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EmptyRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EmptyRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EmptyRequest;

  static equals(a: EmptyRequest | PlainMessage<EmptyRequest> | undefined, b: EmptyRequest | PlainMessage<EmptyRequest> | undefined): boolean;
}

