// @generated by protoc-gen-es v1.3.1
// @generated from file api/resources/stats/v1/stats.proto (package api.resources.stats.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum api.resources.stats.v1.ResourceStatus
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
 * @generated from enum api.resources.stats.v1.ReliabilityScore
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
 * The mode of the upgrade
 *
 * @generated from enum api.resources.stats.v1.UpgradeMode
 */
export declare enum UpgradeMode {
  /**
   * @generated from enum value: UPGRADE_MODE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * Disable upgrades
   *
   * @generated from enum value: UPGRADE_MODE_DISABLED = 1;
   */
  DISABLED = 1,

  /**
   * Notify about new upgrades but do not apply them
   *
   * @generated from enum value: UPGRADE_MODE_NOTIFY = 2;
   */
  NOTIFY = 2,

  /**
   * Notify about new upgrades and apply them
   *
   * @generated from enum value: UPGRADE_MODE_UPGRADE = 3;
   */
  UPGRADE = 3,
}

/**
 * @generated from message api.resources.stats.v1.MonService
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
  static readonly typeName = "api.resources.stats.v1.MonService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MonService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MonService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MonService;

  static equals(a: MonService | PlainMessage<MonService> | undefined, b: MonService | PlainMessage<MonService> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.MgrService
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
  static readonly typeName = "api.resources.stats.v1.MgrService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MgrService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MgrService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MgrService;

  static equals(a: MgrService | PlainMessage<MgrService> | undefined, b: MgrService | PlainMessage<MgrService> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.MdsService
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
  static readonly typeName = "api.resources.stats.v1.MdsService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MdsService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MdsService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MdsService;

  static equals(a: MdsService | PlainMessage<MdsService> | undefined, b: MdsService | PlainMessage<MdsService> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.OsdService
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
  static readonly typeName = "api.resources.stats.v1.OsdService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OsdService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OsdService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OsdService;

  static equals(a: OsdService | PlainMessage<OsdService> | undefined, b: OsdService | PlainMessage<OsdService> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.RgwService
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
  static readonly typeName = "api.resources.stats.v1.RgwService";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RgwService;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RgwService;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RgwService;

  static equals(a: RgwService | PlainMessage<RgwService> | undefined, b: RgwService | PlainMessage<RgwService> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.Services
 */
export declare class Services extends Message<Services> {
  /**
   * @generated from field: api.resources.stats.v1.MonService mon = 1;
   */
  mon?: MonService;

  /**
   * @generated from field: api.resources.stats.v1.MgrService mgr = 2;
   */
  mgr?: MgrService;

  /**
   * @generated from field: api.resources.stats.v1.MdsService mds = 3;
   */
  mds?: MdsService;

  /**
   * @generated from field: api.resources.stats.v1.OsdService osd = 4;
   */
  osd?: OsdService;

  /**
   * @generated from field: api.resources.stats.v1.RgwService rgw = 5;
   */
  rgw?: RgwService;

  constructor(data?: PartialMessage<Services>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.Services";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Services;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Services;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Services;

  static equals(a: Services | PlainMessage<Services> | undefined, b: Services | PlainMessage<Services> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.PGs
 */
export declare class PGs extends Message<PGs> {
  /**
   * @generated from field: int32 active_clean = 1;
   */
  activeClean: number;

  constructor(data?: PartialMessage<PGs>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.PGs";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PGs;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PGs;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PGs;

  static equals(a: PGs | PlainMessage<PGs> | undefined, b: PGs | PlainMessage<PGs> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.Pools
 */
export declare class Pools extends Message<Pools> {
  /**
   * @generated from field: int32 pools = 1;
   */
  pools: number;

  /**
   * @generated from field: api.resources.stats.v1.PGs pgs = 2;
   */
  pgs?: PGs;

  constructor(data?: PartialMessage<Pools>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.Pools";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pools;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pools;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pools;

  static equals(a: Pools | PlainMessage<Pools> | undefined, b: Pools | PlainMessage<Pools> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.Objects
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
  static readonly typeName = "api.resources.stats.v1.Objects";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Objects;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Objects;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Objects;

  static equals(a: Objects | PlainMessage<Objects> | undefined, b: Objects | PlainMessage<Objects> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.Usage
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
  static readonly typeName = "api.resources.stats.v1.Usage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Usage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Usage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Usage;

  static equals(a: Usage | PlainMessage<Usage> | undefined, b: Usage | PlainMessage<Usage> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.Data
 */
export declare class Data extends Message<Data> {
  /**
   * @generated from field: int32 volumes = 1;
   */
  volumes: number;

  /**
   * @generated from field: api.resources.stats.v1.Pools pools = 2;
   */
  pools?: Pools;

  /**
   * @generated from field: api.resources.stats.v1.Objects objects = 3;
   */
  objects?: Objects;

  /**
   * @generated from field: api.resources.stats.v1.Usage usage = 4;
   */
  usage?: Usage;

  constructor(data?: PartialMessage<Data>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.Data";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Data;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Data;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Data;

  static equals(a: Data | PlainMessage<Data> | undefined, b: Data | PlainMessage<Data> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.IOPS
 */
export declare class IOPS extends Message<IOPS> {
  /**
   * @generated from field: string client_read = 1;
   */
  clientRead: string;

  /**
   * @generated from field: string client_write = 2;
   */
  clientWrite: string;

  /**
   * @generated from field: string client_read_ops = 3;
   */
  clientReadOps: string;

  /**
   * @generated from field: string client_write_ops = 4;
   */
  clientWriteOps: string;

  constructor(data?: PartialMessage<IOPS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.IOPS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IOPS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IOPS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IOPS;

  static equals(a: IOPS | PlainMessage<IOPS> | undefined, b: IOPS | PlainMessage<IOPS> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.Crash
 */
export declare class Crash extends Message<Crash> {
  /**
   * @generated from field: string description = 1;
   */
  description: string;

  constructor(data?: PartialMessage<Crash>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.Crash";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Crash;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Crash;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Crash;

  static equals(a: Crash | PlainMessage<Crash> | undefined, b: Crash | PlainMessage<Crash> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.ClusterStats
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
   * @generated from field: repeated api.resources.stats.v1.Crash crashes = 3;
   */
  crashes: Crash[];

  /**
   * @generated from field: api.resources.stats.v1.Services services = 4;
   */
  services?: Services;

  /**
   * @generated from field: api.resources.stats.v1.Data data = 5;
   */
  data?: Data;

  /**
   * @generated from field: string test = 6;
   */
  test: string;

  /**
   * @generated from field: api.resources.stats.v1.IOPS iops = 7;
   */
  iops?: IOPS;

  constructor(data?: PartialMessage<ClusterStats>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.ClusterStats";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterStats;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterStats;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterStats;

  static equals(a: ClusterStats | PlainMessage<ClusterStats> | undefined, b: ClusterStats | PlainMessage<ClusterStats> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.ResourceInfo
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
   * @generated from field: api.resources.stats.v1.ResourceStatus status = 5;
   */
  status: ResourceStatus;

  /**
   * @generated from field: int32 replicas = 6;
   */
  replicas: number;

  /**
   * @generated from field: api.resources.stats.v1.ReliabilityScore reliability = 7;
   */
  reliability: ReliabilityScore;

  constructor(data?: PartialMessage<ResourceInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.ResourceInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResourceInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResourceInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResourceInfo;

  static equals(a: ResourceInfo | PlainMessage<ResourceInfo> | undefined, b: ResourceInfo | PlainMessage<ResourceInfo> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.NodeInfo
 */
export declare class NodeInfo extends Message<NodeInfo> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: api.resources.stats.v1.ResourceStatus status = 2;
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
  static readonly typeName = "api.resources.stats.v1.NodeInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeInfo;

  static equals(a: NodeInfo | PlainMessage<NodeInfo> | undefined, b: NodeInfo | PlainMessage<NodeInfo> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.ClusterRadar
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
  static readonly typeName = "api.resources.stats.v1.ClusterRadar";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterRadar;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterRadar;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterRadar;

  static equals(a: ClusterRadar | PlainMessage<ClusterRadar> | undefined, b: ClusterRadar | PlainMessage<ClusterRadar> | undefined): boolean;
}

/**
 * Represents a map of products to version strings.
 *
 * @generated from message api.resources.stats.v1.ProductVersions
 */
export declare class ProductVersions extends Message<ProductVersions> {
  /**
   * Kubernetes version, must be a valid semver.
   *
   * @generated from field: string kube = 1;
   */
  kube: string;

  /**
   * Koor Operator version, must be a valid semver.
   *
   * @generated from field: string koor_operator = 2;
   */
  koorOperator: string;

  /**
   * Koor Storage Distribution version, must be a valid semver.
   *
   * @generated from field: string ksd = 3;
   */
  ksd: string;

  /**
   * Ceph version, must be a valid semver.
   *
   * @generated from field: string ceph = 4;
   */
  ceph: string;

  constructor(data?: PartialMessage<ProductVersions>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.ProductVersions";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProductVersions;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProductVersions;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProductVersions;

  static equals(a: ProductVersions | PlainMessage<ProductVersions> | undefined, b: ProductVersions | PlainMessage<ProductVersions> | undefined): boolean;
}

/**
 * Defines a detailed version of a product, which includes a container image or a helm chart.
 *
 * @generated from message api.resources.stats.v1.DetailedVersion
 */
export declare class DetailedVersion extends Message<DetailedVersion> {
  /**
   * The product version, must be a valid semver.
   *
   * @generated from field: string version = 1;
   */
  version: string;

  /**
   * The URI of the container image.
   *
   * @generated from field: string image_uri = 2;
   */
  imageUri: string;

  /**
   * The hash of the container image.
   *
   * @generated from field: string image_hash = 3;
   */
  imageHash: string;

  /**
   * The URI of the helm repository.
   *
   * @generated from field: string helm_repository = 4;
   */
  helmRepository: string;

  /**
   * The name of the helm chart in the repository.
   *
   * @generated from field: string helm_chart = 5;
   */
  helmChart: string;

  constructor(data?: PartialMessage<DetailedVersion>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.DetailedVersion";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DetailedVersion;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DetailedVersion;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DetailedVersion;

  static equals(a: DetailedVersion | PlainMessage<DetailedVersion> | undefined, b: DetailedVersion | PlainMessage<DetailedVersion> | undefined): boolean;
}

/**
 * Represents a map of products to detailed versions, which include images or helm charts.
 *
 * @generated from message api.resources.stats.v1.DetailedProductVersions
 */
export declare class DetailedProductVersions extends Message<DetailedProductVersions> {
  /**
   * The detailed Koor Operator version.
   *
   * @generated from field: api.resources.stats.v1.DetailedVersion koor_operator = 1;
   */
  koorOperator?: DetailedVersion;

  /**
   * The detailed Koor Storage Distribution version.
   *
   * @generated from field: api.resources.stats.v1.DetailedVersion ksd = 2;
   */
  ksd?: DetailedVersion;

  /**
   * The detailed Ceph version.
   *
   * @generated from field: api.resources.stats.v1.DetailedVersion ceph = 3;
   */
  ceph?: DetailedVersion;

  constructor(data?: PartialMessage<DetailedProductVersions>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.DetailedProductVersions";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DetailedProductVersions;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DetailedProductVersions;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DetailedProductVersions;

  static equals(a: DetailedProductVersions | PlainMessage<DetailedProductVersions> | undefined, b: DetailedProductVersions | PlainMessage<DetailedProductVersions> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.ClusterResources
 */
export declare class ClusterResources extends Message<ClusterResources> {
  /**
   * The number of nodes in the cluster
   *
   * @generated from field: string nodes = 1;
   */
  nodes: string;

  /**
   * Ephemeral Storage available
   *
   * @generated from field: string storage = 2;
   */
  storage: string;

  /**
   * CPU cores available
   *
   * @generated from field: string cpu = 3;
   */
  cpu: string;

  /**
   * Memory available
   *
   * @generated from field: string memory = 4;
   */
  memory: string;

  constructor(data?: PartialMessage<ClusterResources>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.ClusterResources";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClusterResources;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClusterResources;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClusterResources;

  static equals(a: ClusterResources | PlainMessage<ClusterResources> | undefined, b: ClusterResources | PlainMessage<ClusterResources> | undefined): boolean;
}

/**
 * Represents the status of the KoorCluster CRD
 *
 * @generated from message api.resources.stats.v1.KoorClusterStatus
 */
export declare class KoorClusterStatus extends Message<KoorClusterStatus> {
  /**
   * The total resources available in the cluster nodes
   *
   * @generated from field: api.resources.stats.v1.ClusterResources total_resources = 1;
   */
  totalResources?: ClusterResources;

  /**
   * Does the cluster meet the minimum recommended resources
   *
   * @generated from field: bool meets_minimum_resources = 2;
   */
  meetsMinimumResources: boolean;

  /**
   * The current versions of rook and ceph
   *
   * @generated from field: api.resources.stats.v1.ProductVersions current_versions = 3;
   */
  currentVersions?: ProductVersions;

  /**
   * The latest versions of rook and ceph
   *
   * @generated from field: api.resources.stats.v1.DetailedProductVersions latest_versions = 4;
   */
  latestVersions?: DetailedProductVersions;

  constructor(data?: PartialMessage<KoorClusterStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.KoorClusterStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KoorClusterStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KoorClusterStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KoorClusterStatus;

  static equals(a: KoorClusterStatus | PlainMessage<KoorClusterStatus> | undefined, b: KoorClusterStatus | PlainMessage<KoorClusterStatus> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.UpgradeOptions
 */
export declare class UpgradeOptions extends Message<UpgradeOptions> {
  /**
   * Upgrade mode
   *
   * @generated from field: api.resources.stats.v1.UpgradeMode mode = 1;
   */
  mode: UpgradeMode;

  /**
   * The api endpoint used to find the ceph latest version
   *
   * @generated from field: string endpoint = 2;
   */
  endpoint: string;

  /**
   *
   * The schedule to check for new versions. Uses CRON format as specified by https://github.com/robfig/cron/tree/v3.
   * Defaults to everyday at midnight in the local timezone.
   * To change the timezone, prefix the schedule with CRON_TZ=<Timezone>.
   * For example: "CRON_TZ=UTC 0 0 * * *" is midnight UTC.
   *
   * @generated from field: string schedule = 3;
   */
  schedule: string;

  constructor(data?: PartialMessage<UpgradeOptions>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.UpgradeOptions";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpgradeOptions;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpgradeOptions;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpgradeOptions;

  static equals(a: UpgradeOptions | PlainMessage<UpgradeOptions> | undefined, b: UpgradeOptions | PlainMessage<UpgradeOptions> | undefined): boolean;
}

/**
 * Represents the state of KoorCluster
 *
 * @generated from message api.resources.stats.v1.KoorClusterSpec
 */
export declare class KoorClusterSpec extends Message<KoorClusterSpec> {
  /**
   * Use all devices on nodes
   *
   * @generated from field: google.protobuf.BoolValue use_all_devices = 1;
   */
  useAllDevices?: boolean;

  /**
   * Enable monitoring. Requires Prometheus to be pre-installed.
   *
   * @generated from field: google.protobuf.BoolValue monitoring_enabled = 2;
   */
  monitoringEnabled?: boolean;

  /**
   * Enable the ceph dashboard for viewing cluster status
   *
   * @generated from field: google.protobuf.BoolValue dashboard_enabled = 3;
   */
  dashboardEnabled?: boolean;

  /**
   * Installs a debugging toolbox deployment
   *
   * @generated from field: google.protobuf.BoolValue toolbox_enabled = 4;
   */
  toolboxEnabled?: boolean;

  /**
   * Specifies the upgrade options for new ceph versions
   *
   * @generated from field: api.resources.stats.v1.UpgradeOptions upgrade_options = 5;
   */
  upgradeOptions?: UpgradeOptions;

  /**
   * The name to use for KSD helm release.
   *
   * @generated from field: string ksd_release_name = 6;
   */
  ksdReleaseName: string;

  /**
   * The name to use for KSD cluster helm release.
   *
   * @generated from field: string ksd_cluster_release_name = 7;
   */
  ksdClusterReleaseName: string;

  constructor(data?: PartialMessage<KoorClusterSpec>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.KoorClusterSpec";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KoorClusterSpec;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KoorClusterSpec;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KoorClusterSpec;

  static equals(a: KoorClusterSpec | PlainMessage<KoorClusterSpec> | undefined, b: KoorClusterSpec | PlainMessage<KoorClusterSpec> | undefined): boolean;
}

/**
 * @generated from message api.resources.stats.v1.KoorCluster
 */
export declare class KoorCluster extends Message<KoorCluster> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string namespace = 2;
   */
  namespace: string;

  /**
   * @generated from field: api.resources.stats.v1.KoorClusterSpec spec = 3;
   */
  spec?: KoorClusterSpec;

  /**
   * @generated from field: api.resources.stats.v1.KoorClusterStatus status = 4;
   */
  status?: KoorClusterStatus;

  constructor(data?: PartialMessage<KoorCluster>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.resources.stats.v1.KoorCluster";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KoorCluster;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KoorCluster;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KoorCluster;

  static equals(a: KoorCluster | PlainMessage<KoorCluster> | undefined, b: KoorCluster | PlainMessage<KoorCluster> | undefined): boolean;
}

