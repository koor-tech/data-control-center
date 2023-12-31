// @generated by protoc-gen-es v1.6.0
// @generated from file api/resources/ceph/v1/config.proto (package api.resources.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * Osd Scrubbing schedule config
 *
 * @generated from message api.resources.ceph.v1.OSDScrubbingSchedule
 */
export const OSDScrubbingSchedule = proto3.makeMessageType(
  "api.resources.ceph.v1.OSDScrubbingSchedule",
  () => [
    { no: 1, name: "apply_schedule", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "max_scrub_ops", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 3, name: "begin_hour", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 4, name: "end_hour", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 5, name: "begin_week_day", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 6, name: "end_week_day", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 7, name: "min_scrub_interval", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 8, name: "max_scrub_interval", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 9, name: "deep_scrub_interval", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 10, name: "scrub_sleep_seconds", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

