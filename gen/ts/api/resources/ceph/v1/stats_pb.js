// @generated by protoc-gen-es v1.6.0
// @generated from file api/resources/ceph/v1/stats.proto (package api.resources.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum api.resources.ceph.v1.ClusterHealth
 */
export const ClusterHealth = proto3.makeEnum(
  "api.resources.ceph.v1.ClusterHealth",
  [
    {no: 0, name: "CLUSTER_HEALTH_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "CLUSTER_HEALTH_OFFLINE", localName: "OFFLINE"},
    {no: 2, name: "CLUSTER_HEALTH_OK", localName: "OK"},
    {no: 3, name: "CLUSTER_HEALTH_WARN", localName: "WARN"},
    {no: 4, name: "CLUSTER_HEALTH_ERR", localName: "ERR"},
  ],
);

/**
 * @generated from message api.resources.ceph.v1.MonService
 */
export const MonService = proto3.makeMessageType(
  "api.resources.ceph.v1.MonService",
  () => [
    { no: 1, name: "daemon_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "quorum", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "created_since", kind: "message", T: Timestamp },
    { no: 6, name: "updated_since", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.MgrService
 */
export const MgrService = proto3.makeMessageType(
  "api.resources.ceph.v1.MgrService",
  () => [
    { no: 1, name: "active", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "standbys", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "updated_since", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.MdsService
 */
export const MdsService = proto3.makeMessageType(
  "api.resources.ceph.v1.MdsService",
  () => [
    { no: 1, name: "daemons_up", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "hot_standby_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.OsdService
 */
export const OsdService = proto3.makeMessageType(
  "api.resources.ceph.v1.OsdService",
  () => [
    { no: 1, name: "osd_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "osd_up", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "osd_in", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "osd_up_updated_since", kind: "message", T: Timestamp },
    { no: 7, name: "osd_in_updated_since", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.RgwService
 */
export const RgwService = proto3.makeMessageType(
  "api.resources.ceph.v1.RgwService",
  () => [
    { no: 1, name: "active_daemon", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "host_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "zone_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.Services
 */
export const Services = proto3.makeMessageType(
  "api.resources.ceph.v1.Services",
  () => [
    { no: 1, name: "mon", kind: "message", T: MonService },
    { no: 2, name: "mgr", kind: "message", T: MgrService },
    { no: 3, name: "mds", kind: "message", T: MdsService },
    { no: 4, name: "osd", kind: "message", T: OsdService },
    { no: 5, name: "rgw", kind: "message", T: RgwService },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.PGs
 */
export const PGs = proto3.makeMessageType(
  "api.resources.ceph.v1.PGs",
  () => [
    { no: 1, name: "active_clean", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.Pools
 */
export const Pools = proto3.makeMessageType(
  "api.resources.ceph.v1.Pools",
  () => [
    { no: 1, name: "pools", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "pgs", kind: "message", T: PGs },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.Objects
 */
export const Objects = proto3.makeMessageType(
  "api.resources.ceph.v1.Objects",
  () => [
    { no: 1, name: "object_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "size", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.Usage
 */
export const Usage = proto3.makeMessageType(
  "api.resources.ceph.v1.Usage",
  () => [
    { no: 1, name: "used", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "available", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "total", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.Data
 */
export const Data = proto3.makeMessageType(
  "api.resources.ceph.v1.Data",
  () => [
    { no: 1, name: "volumes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "pools", kind: "message", T: Pools },
    { no: 3, name: "objects", kind: "message", T: Objects },
    { no: 4, name: "usage", kind: "message", T: Usage },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.IOPS
 */
export const IOPS = proto3.makeMessageType(
  "api.resources.ceph.v1.IOPS",
  () => [
    { no: 1, name: "client_read", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "client_write", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "client_read_ops", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "client_write_ops", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.Crash
 */
export const Crash = proto3.makeMessageType(
  "api.resources.ceph.v1.Crash",
  () => [
    { no: 1, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.ClusterStats
 */
export const ClusterStats = proto3.makeMessageType(
  "api.resources.ceph.v1.ClusterStats",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(ClusterHealth) },
    { no: 3, name: "crashes", kind: "message", T: Crash, repeated: true },
    { no: 4, name: "services", kind: "message", T: Services },
    { no: 5, name: "data", kind: "message", T: Data },
    { no: 6, name: "test", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "iops", kind: "message", T: IOPS },
  ],
);

/**
 * @generated from message api.resources.ceph.v1.ClusterRadar
 */
export const ClusterRadar = proto3.makeMessageType(
  "api.resources.ceph.v1.ClusterRadar",
  () => [
    { no: 1, name: "cluster_health", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "nodes_health", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 3, name: "capacity_available", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "stability", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "reliability", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);
