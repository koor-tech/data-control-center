// @generated by protoc-gen-es v1.3.1
// @generated from file api/resources/koor/v1/koor.proto (package api.resources.koor.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { BoolValue, proto3 } from "@bufbuild/protobuf";

/**
 * The mode of the upgrade
 *
 * @generated from enum api.resources.koor.v1.UpgradeMode
 */
export const UpgradeMode = proto3.makeEnum(
  "api.resources.koor.v1.UpgradeMode",
  [
    {no: 0, name: "UPGRADE_MODE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "UPGRADE_MODE_DISABLED", localName: "DISABLED"},
    {no: 2, name: "UPGRADE_MODE_NOTIFY", localName: "NOTIFY"},
    {no: 3, name: "UPGRADE_MODE_UPGRADE", localName: "UPGRADE"},
  ],
);

/**
 * Represents a map of products to version strings.
 *
 * @generated from message api.resources.koor.v1.ProductVersions
 */
export const ProductVersions = proto3.makeMessageType(
  "api.resources.koor.v1.ProductVersions",
  () => [
    { no: 1, name: "kube", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "koor_operator", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "ksd", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "ceph", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Defines a detailed version of a product, which includes a container image or a helm chart.
 *
 * @generated from message api.resources.koor.v1.DetailedVersion
 */
export const DetailedVersion = proto3.makeMessageType(
  "api.resources.koor.v1.DetailedVersion",
  () => [
    { no: 1, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "image_uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "image_hash", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "helm_repository", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "helm_chart", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Represents a map of products to detailed versions, which include images or helm charts.
 *
 * @generated from message api.resources.koor.v1.DetailedProductVersions
 */
export const DetailedProductVersions = proto3.makeMessageType(
  "api.resources.koor.v1.DetailedProductVersions",
  () => [
    { no: 1, name: "koor_operator", kind: "message", T: DetailedVersion },
    { no: 2, name: "ksd", kind: "message", T: DetailedVersion },
    { no: 3, name: "ceph", kind: "message", T: DetailedVersion },
  ],
);

/**
 * @generated from message api.resources.koor.v1.ClusterResources
 */
export const ClusterResources = proto3.makeMessageType(
  "api.resources.koor.v1.ClusterResources",
  () => [
    { no: 1, name: "nodes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "storage", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "cpu", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "memory", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Represents the status of the KoorCluster CRD
 *
 * @generated from message api.resources.koor.v1.KoorClusterStatus
 */
export const KoorClusterStatus = proto3.makeMessageType(
  "api.resources.koor.v1.KoorClusterStatus",
  () => [
    { no: 1, name: "total_resources", kind: "message", T: ClusterResources },
    { no: 2, name: "meets_minimum_resources", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "current_versions", kind: "message", T: ProductVersions },
    { no: 4, name: "latest_versions", kind: "message", T: DetailedProductVersions },
  ],
);

/**
 * @generated from message api.resources.koor.v1.UpgradeOptions
 */
export const UpgradeOptions = proto3.makeMessageType(
  "api.resources.koor.v1.UpgradeOptions",
  () => [
    { no: 1, name: "mode", kind: "enum", T: proto3.getEnumType(UpgradeMode) },
    { no: 2, name: "endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "schedule", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Represents the state of KoorCluster
 *
 * @generated from message api.resources.koor.v1.KoorClusterSpec
 */
export const KoorClusterSpec = proto3.makeMessageType(
  "api.resources.koor.v1.KoorClusterSpec",
  () => [
    { no: 1, name: "use_all_devices", kind: "message", T: BoolValue },
    { no: 2, name: "monitoring_enabled", kind: "message", T: BoolValue },
    { no: 3, name: "dashboard_enabled", kind: "message", T: BoolValue },
    { no: 4, name: "toolbox_enabled", kind: "message", T: BoolValue },
    { no: 5, name: "upgrade_options", kind: "message", T: UpgradeOptions },
    { no: 6, name: "ksd_release_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "ksd_cluster_release_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.resources.koor.v1.KoorCluster
 */
export const KoorCluster = proto3.makeMessageType(
  "api.resources.koor.v1.KoorCluster",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "spec", kind: "message", T: KoorClusterSpec },
    { no: 4, name: "status", kind: "message", T: KoorClusterStatus },
  ],
);

