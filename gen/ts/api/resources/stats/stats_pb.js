// @generated by protoc-gen-es v1.3.1
// @generated from file api/resources/stats/stats.proto (package resources.stats, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum resources.stats.ResourceStatus
 */
export const ResourceStatus = proto3.makeEnum(
  "resources.stats.ResourceStatus",
  [
    {no: 0, name: "RESOURCE_UNKNOWN"},
    {no: 1, name: "RESOURCE_READY"},
    {no: 2, name: "RESOURCE_NOT_READY"},
  ],
);

/**
 * @generated from message resources.stats.MonService
 */
export const MonService = proto3.makeMessageType(
  "resources.stats.MonService",
  () => [
    { no: 1, name: "daemon_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "quorum", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created_at", kind: "message", T: Timestamp },
    { no: 4, name: "updated_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message resources.stats.MgrService
 */
export const MgrService = proto3.makeMessageType(
  "resources.stats.MgrService",
  () => [
    { no: 1, name: "active", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "standbys", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "updated_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message resources.stats.MdsService
 */
export const MdsService = proto3.makeMessageType(
  "resources.stats.MdsService",
  () => [
    { no: 1, name: "daemons_up", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "hot_standby_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message resources.stats.OsdService
 */
export const OsdService = proto3.makeMessageType(
  "resources.stats.OsdService",
  () => [
    { no: 1, name: "osd_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "osd_up", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "osd_in", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "osd_up_updated_at", kind: "message", T: Timestamp },
    { no: 5, name: "osd_in_updated_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message resources.stats.RgwService
 */
export const RgwService = proto3.makeMessageType(
  "resources.stats.RgwService",
  () => [
    { no: 1, name: "active_daemon", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "host_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "zone_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message resources.stats.Services
 */
export const Services = proto3.makeMessageType(
  "resources.stats.Services",
  () => [
    { no: 1, name: "mon", kind: "message", T: MonService },
    { no: 2, name: "mgr", kind: "message", T: MgrService },
    { no: 3, name: "mds", kind: "message", T: MdsService },
    { no: 4, name: "osd", kind: "message", T: OsdService },
    { no: 5, name: "rgw", kind: "message", T: RgwService },
  ],
);

/**
 * @generated from message resources.stats.Pgs
 */
export const Pgs = proto3.makeMessageType(
  "resources.stats.Pgs",
  () => [
    { no: 1, name: "active_clean", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message resources.stats.Pools
 */
export const Pools = proto3.makeMessageType(
  "resources.stats.Pools",
  () => [
    { no: 1, name: "pools", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "pgs", kind: "message", T: Pgs },
  ],
);

/**
 * @generated from message resources.stats.Objects
 */
export const Objects = proto3.makeMessageType(
  "resources.stats.Objects",
  () => [
    { no: 1, name: "object_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "size", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message resources.stats.Usage
 */
export const Usage = proto3.makeMessageType(
  "resources.stats.Usage",
  () => [
    { no: 1, name: "used", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "available", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "total", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message resources.stats.Data
 */
export const Data = proto3.makeMessageType(
  "resources.stats.Data",
  () => [
    { no: 1, name: "volumes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "pools", kind: "message", T: Pools },
    { no: 3, name: "objects", kind: "message", T: Objects },
    { no: 4, name: "usage", kind: "message", T: Usage },
  ],
);

/**
 * @generated from message resources.stats.Io
 */
export const Io = proto3.makeMessageType(
  "resources.stats.Io",
  () => [
    { no: 1, name: "client_read", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "client_read_ops", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "client_write_ops", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message resources.stats.Crash
 */
export const Crash = proto3.makeMessageType(
  "resources.stats.Crash",
  () => [
    { no: 1, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message resources.stats.ClusterStats
 */
export const ClusterStats = proto3.makeMessageType(
  "resources.stats.ClusterStats",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "status", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "crashes", kind: "message", T: Crash, repeated: true },
    { no: 4, name: "services", kind: "message", T: Services },
    { no: 5, name: "data", kind: "message", T: Data },
    { no: 6, name: "test", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "io", kind: "message", T: Io },
  ],
);

/**
 * @generated from message resources.stats.ResourceInfo
 */
export const ResourceInfo = proto3.makeMessageType(
  "resources.stats.ResourceInfo",
  () => [
    { no: 1, name: "apiversion", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "kind", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "status", kind: "enum", T: proto3.getEnumType(ResourceStatus) },
  ],
);

/**
 * @generated from message resources.stats.NodeInfo
 */
export const NodeInfo = proto3.makeMessageType(
  "resources.stats.NodeInfo",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(ResourceStatus) },
    { no: 3, name: "roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "internal_ip", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "external_ip", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "age", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message resources.stats.ClusterRadar
 */
export const ClusterRadar = proto3.makeMessageType(
  "resources.stats.ClusterRadar",
  () => [
    { no: 1, name: "cluster_health", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "nodes_health", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 3, name: "capacity_used", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "stability", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "reliability", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);
