// @generated by protoc-gen-es v1.3.0
// @generated from file vendor/google.golang.org/appengine/internal/log/log_service.proto (package appengine, syntax proto2)
/* eslint-disable */
// @ts-nocheck

import { proto2 } from "@bufbuild/protobuf";

/**
 * @generated from message appengine.LogServiceError
 */
export const LogServiceError = proto2.makeMessageType(
  "appengine.LogServiceError",
  [],
);

/**
 * @generated from enum appengine.LogServiceError.ErrorCode
 */
export const LogServiceError_ErrorCode = proto2.makeEnum(
  "appengine.LogServiceError.ErrorCode",
  [
    {no: 0, name: "OK"},
    {no: 1, name: "INVALID_REQUEST"},
    {no: 2, name: "STORAGE_ERROR"},
  ],
);

/**
 * @generated from message appengine.UserAppLogLine
 */
export const UserAppLogLine = proto2.makeMessageType(
  "appengine.UserAppLogLine",
  () => [
    { no: 1, name: "timestamp_usec", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "level", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message appengine.UserAppLogGroup
 */
export const UserAppLogGroup = proto2.makeMessageType(
  "appengine.UserAppLogGroup",
  () => [
    { no: 2, name: "log_line", kind: "message", T: UserAppLogLine, repeated: true },
  ],
);

/**
 * @generated from message appengine.FlushRequest
 */
export const FlushRequest = proto2.makeMessageType(
  "appengine.FlushRequest",
  () => [
    { no: 1, name: "logs", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
  ],
);

/**
 * @generated from message appengine.SetStatusRequest
 */
export const SetStatusRequest = proto2.makeMessageType(
  "appengine.SetStatusRequest",
  () => [
    { no: 1, name: "status", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message appengine.LogOffset
 */
export const LogOffset = proto2.makeMessageType(
  "appengine.LogOffset",
  () => [
    { no: 1, name: "request_id", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
  ],
);

/**
 * @generated from message appengine.LogLine
 */
export const LogLine = proto2.makeMessageType(
  "appengine.LogLine",
  () => [
    { no: 1, name: "time", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "level", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "log_message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message appengine.RequestLog
 */
export const RequestLog = proto2.makeMessageType(
  "appengine.RequestLog",
  () => [
    { no: 1, name: "app_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 37, name: "module_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true, default: "default" },
    { no: 2, name: "version_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "request_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 35, name: "offset", kind: "message", T: LogOffset, opt: true },
    { no: 4, name: "ip", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "nickname", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "start_time", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 7, name: "end_time", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 8, name: "latency", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 9, name: "mcycles", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 10, name: "method", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "resource", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "http_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 13, name: "status", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 14, name: "response_size", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 15, name: "referrer", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 16, name: "user_agent", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 17, name: "url_map_entry", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 18, name: "combined", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 19, name: "api_mcycles", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 20, name: "host", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 21, name: "cost", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 22, name: "task_queue_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 23, name: "task_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 24, name: "was_loading_request", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 25, name: "pending_time", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 26, name: "replica_index", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true, default: -1 },
    { no: 27, name: "finished", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true, default: true },
    { no: 28, name: "clone_key", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
    { no: 29, name: "line", kind: "message", T: LogLine, repeated: true },
    { no: 36, name: "lines_incomplete", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 38, name: "app_engine_release", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
    { no: 30, name: "exit_reason", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 31, name: "was_throttled_for_time", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 32, name: "was_throttled_for_requests", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 33, name: "throttled_time", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 34, name: "server_name", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
  ],
);

/**
 * @generated from message appengine.LogModuleVersion
 */
export const LogModuleVersion = proto2.makeMessageType(
  "appengine.LogModuleVersion",
  () => [
    { no: 1, name: "module_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true, default: "default" },
    { no: 2, name: "version_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * @generated from message appengine.LogReadRequest
 */
export const LogReadRequest = proto2.makeMessageType(
  "appengine.LogReadRequest",
  () => [
    { no: 1, name: "app_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "version_id", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 19, name: "module_version", kind: "message", T: LogModuleVersion, repeated: true },
    { no: 3, name: "start_time", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 4, name: "end_time", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 5, name: "offset", kind: "message", T: LogOffset, opt: true },
    { no: 6, name: "request_id", kind: "scalar", T: 12 /* ScalarType.BYTES */, repeated: true },
    { no: 7, name: "minimum_log_level", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 8, name: "include_incomplete", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 9, name: "count", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 14, name: "combined_log_regex", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 15, name: "host_regex", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 16, name: "replica_index", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 10, name: "include_app_logs", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 17, name: "app_logs_per_request", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 11, name: "include_host", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 12, name: "include_all", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 13, name: "cache_iterator", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 18, name: "num_shards", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ],
);

/**
 * @generated from message appengine.LogReadResponse
 */
export const LogReadResponse = proto2.makeMessageType(
  "appengine.LogReadResponse",
  () => [
    { no: 1, name: "log", kind: "message", T: RequestLog, repeated: true },
    { no: 2, name: "offset", kind: "message", T: LogOffset, opt: true },
    { no: 3, name: "last_end_time", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
  ],
);

/**
 * @generated from message appengine.LogUsageRecord
 */
export const LogUsageRecord = proto2.makeMessageType(
  "appengine.LogUsageRecord",
  () => [
    { no: 1, name: "version_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "start_time", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 3, name: "end_time", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 4, name: "count", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 5, name: "total_size", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 6, name: "records", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ],
);

/**
 * @generated from message appengine.LogUsageRequest
 */
export const LogUsageRequest = proto2.makeMessageType(
  "appengine.LogUsageRequest",
  () => [
    { no: 1, name: "app_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "version_id", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "start_time", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 4, name: "end_time", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 5, name: "resolution_hours", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true, default: 1 },
    { no: 6, name: "combine_versions", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 7, name: "usage_version", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 8, name: "versions_only", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * @generated from message appengine.LogUsageResponse
 */
export const LogUsageResponse = proto2.makeMessageType(
  "appengine.LogUsageResponse",
  () => [
    { no: 1, name: "usage", kind: "message", T: LogUsageRecord, repeated: true },
    { no: 2, name: "summary", kind: "message", T: LogUsageRecord, opt: true },
  ],
);

