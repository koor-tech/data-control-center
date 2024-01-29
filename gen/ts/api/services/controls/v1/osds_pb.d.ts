// @generated by protoc-gen-es v1.6.0
// @generated from file api/services/controls/v1/osds.proto (package api.services.controls.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { OSDScrubbingSchedule } from "../../../resources/ceph/v1/controls_pb.js";

/**
 * @generated from message api.services.controls.v1.GetScrubbingScheduleRequest
 */
export declare class GetScrubbingScheduleRequest extends Message<GetScrubbingScheduleRequest> {
  constructor(data?: PartialMessage<GetScrubbingScheduleRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.controls.v1.GetScrubbingScheduleRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetScrubbingScheduleRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetScrubbingScheduleRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetScrubbingScheduleRequest;

  static equals(a: GetScrubbingScheduleRequest | PlainMessage<GetScrubbingScheduleRequest> | undefined, b: GetScrubbingScheduleRequest | PlainMessage<GetScrubbingScheduleRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.controls.v1.GetScrubbingScheduleResponse
 */
export declare class GetScrubbingScheduleResponse extends Message<GetScrubbingScheduleResponse> {
  /**
   * @generated from field: api.resources.ceph.v1.OSDScrubbingSchedule osd_scrubbing_schedule = 1;
   */
  osdScrubbingSchedule?: OSDScrubbingSchedule;

  constructor(data?: PartialMessage<GetScrubbingScheduleResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.controls.v1.GetScrubbingScheduleResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetScrubbingScheduleResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetScrubbingScheduleResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetScrubbingScheduleResponse;

  static equals(a: GetScrubbingScheduleResponse | PlainMessage<GetScrubbingScheduleResponse> | undefined, b: GetScrubbingScheduleResponse | PlainMessage<GetScrubbingScheduleResponse> | undefined): boolean;
}

/**
 * @generated from message api.services.controls.v1.SetScrubbingScheduleRequest
 */
export declare class SetScrubbingScheduleRequest extends Message<SetScrubbingScheduleRequest> {
  /**
   * @generated from field: api.resources.ceph.v1.OSDScrubbingSchedule osd_scrubbing_schedule = 1;
   */
  osdScrubbingSchedule?: OSDScrubbingSchedule;

  constructor(data?: PartialMessage<SetScrubbingScheduleRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.controls.v1.SetScrubbingScheduleRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetScrubbingScheduleRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetScrubbingScheduleRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetScrubbingScheduleRequest;

  static equals(a: SetScrubbingScheduleRequest | PlainMessage<SetScrubbingScheduleRequest> | undefined, b: SetScrubbingScheduleRequest | PlainMessage<SetScrubbingScheduleRequest> | undefined): boolean;
}

/**
 * @generated from message api.services.controls.v1.SetScrubbingScheduleResponse
 */
export declare class SetScrubbingScheduleResponse extends Message<SetScrubbingScheduleResponse> {
  /**
   * @generated from field: api.resources.ceph.v1.OSDScrubbingSchedule osd_scrubbing_schedule = 1;
   */
  osdScrubbingSchedule?: OSDScrubbingSchedule;

  constructor(data?: PartialMessage<SetScrubbingScheduleResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.services.controls.v1.SetScrubbingScheduleResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetScrubbingScheduleResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetScrubbingScheduleResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetScrubbingScheduleResponse;

  static equals(a: SetScrubbingScheduleResponse | PlainMessage<SetScrubbingScheduleResponse> | undefined, b: SetScrubbingScheduleResponse | PlainMessage<SetScrubbingScheduleResponse> | undefined): boolean;
}
