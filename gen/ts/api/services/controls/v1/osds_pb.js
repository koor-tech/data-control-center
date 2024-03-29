// @generated by protoc-gen-es v1.7.2
// @generated from file api/services/controls/v1/osds.proto (package api.services.controls.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { OSDScrubbingSchedule } from "../../../resources/ceph/v1/controls_pb.js";

/**
 * @generated from message api.services.controls.v1.GetScrubbingScheduleRequest
 */
export const GetScrubbingScheduleRequest = proto3.makeMessageType(
  "api.services.controls.v1.GetScrubbingScheduleRequest",
  [],
);

/**
 * @generated from message api.services.controls.v1.GetScrubbingScheduleResponse
 */
export const GetScrubbingScheduleResponse = proto3.makeMessageType(
  "api.services.controls.v1.GetScrubbingScheduleResponse",
  () => [
    { no: 1, name: "osd_scrubbing_schedule", kind: "message", T: OSDScrubbingSchedule },
  ],
);

/**
 * @generated from message api.services.controls.v1.SetScrubbingScheduleRequest
 */
export const SetScrubbingScheduleRequest = proto3.makeMessageType(
  "api.services.controls.v1.SetScrubbingScheduleRequest",
  () => [
    { no: 1, name: "osd_scrubbing_schedule", kind: "message", T: OSDScrubbingSchedule },
  ],
);

/**
 * @generated from message api.services.controls.v1.SetScrubbingScheduleResponse
 */
export const SetScrubbingScheduleResponse = proto3.makeMessageType(
  "api.services.controls.v1.SetScrubbingScheduleResponse",
  () => [
    { no: 1, name: "osd_scrubbing_schedule", kind: "message", T: OSDScrubbingSchedule },
  ],
);

