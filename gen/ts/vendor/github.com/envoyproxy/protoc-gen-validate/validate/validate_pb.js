// @generated by protoc-gen-es v1.3.0
// @generated from file vendor/github.com/envoyproxy/protoc-gen-validate/validate/validate.proto (package validate, syntax proto2)
/* eslint-disable */
// @ts-nocheck

import { Duration, proto2, Timestamp } from "@bufbuild/protobuf";

/**
 * WellKnownRegex contain some well-known patterns.
 *
 * @generated from enum validate.KnownRegex
 */
export const KnownRegex = proto2.makeEnum(
  "validate.KnownRegex",
  [
    {no: 0, name: "UNKNOWN"},
    {no: 1, name: "HTTP_HEADER_NAME"},
    {no: 2, name: "HTTP_HEADER_VALUE"},
  ],
);

/**
 * FieldRules encapsulates the rules for each type of field. Depending on the
 * field, the correct set should be used to ensure proper validations.
 *
 * @generated from message validate.FieldRules
 */
export const FieldRules = proto2.makeMessageType(
  "validate.FieldRules",
  () => [
    { no: 17, name: "message", kind: "message", T: MessageRules, opt: true },
    { no: 1, name: "float", kind: "message", T: FloatRules, oneof: "type" },
    { no: 2, name: "double", kind: "message", T: DoubleRules, oneof: "type" },
    { no: 3, name: "int32", kind: "message", T: Int32Rules, oneof: "type" },
    { no: 4, name: "int64", kind: "message", T: Int64Rules, oneof: "type" },
    { no: 5, name: "uint32", kind: "message", T: UInt32Rules, oneof: "type" },
    { no: 6, name: "uint64", kind: "message", T: UInt64Rules, oneof: "type" },
    { no: 7, name: "sint32", kind: "message", T: SInt32Rules, oneof: "type" },
    { no: 8, name: "sint64", kind: "message", T: SInt64Rules, oneof: "type" },
    { no: 9, name: "fixed32", kind: "message", T: Fixed32Rules, oneof: "type" },
    { no: 10, name: "fixed64", kind: "message", T: Fixed64Rules, oneof: "type" },
    { no: 11, name: "sfixed32", kind: "message", T: SFixed32Rules, oneof: "type" },
    { no: 12, name: "sfixed64", kind: "message", T: SFixed64Rules, oneof: "type" },
    { no: 13, name: "bool", kind: "message", T: BoolRules, oneof: "type" },
    { no: 14, name: "string", kind: "message", T: StringRules, oneof: "type" },
    { no: 15, name: "bytes", kind: "message", T: BytesRules, oneof: "type" },
    { no: 16, name: "enum", kind: "message", T: EnumRules, oneof: "type" },
    { no: 18, name: "repeated", kind: "message", T: RepeatedRules, oneof: "type" },
    { no: 19, name: "map", kind: "message", T: MapRules, oneof: "type" },
    { no: 20, name: "any", kind: "message", T: AnyRules, oneof: "type" },
    { no: 21, name: "duration", kind: "message", T: DurationRules, oneof: "type" },
    { no: 22, name: "timestamp", kind: "message", T: TimestampRules, oneof: "type" },
  ],
);

/**
 * FloatRules describes the constraints applied to `float` values
 *
 * @generated from message validate.FloatRules
 */
export const FloatRules = proto2.makeMessageType(
  "validate.FloatRules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * DoubleRules describes the constraints applied to `double` values
 *
 * @generated from message validate.DoubleRules
 */
export const DoubleRules = proto2.makeMessageType(
  "validate.DoubleRules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * Int32Rules describes the constraints applied to `int32` values
 *
 * @generated from message validate.Int32Rules
 */
export const Int32Rules = proto2.makeMessageType(
  "validate.Int32Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * Int64Rules describes the constraints applied to `int64` values
 *
 * @generated from message validate.Int64Rules
 */
export const Int64Rules = proto2.makeMessageType(
  "validate.Int64Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 3 /* ScalarType.INT64 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 3 /* ScalarType.INT64 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * UInt32Rules describes the constraints applied to `uint32` values
 *
 * @generated from message validate.UInt32Rules
 */
export const UInt32Rules = proto2.makeMessageType(
  "validate.UInt32Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 13 /* ScalarType.UINT32 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 13 /* ScalarType.UINT32 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * UInt64Rules describes the constraints applied to `uint64` values
 *
 * @generated from message validate.UInt64Rules
 */
export const UInt64Rules = proto2.makeMessageType(
  "validate.UInt64Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 4 /* ScalarType.UINT64 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 4 /* ScalarType.UINT64 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * SInt32Rules describes the constraints applied to `sint32` values
 *
 * @generated from message validate.SInt32Rules
 */
export const SInt32Rules = proto2.makeMessageType(
  "validate.SInt32Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 17 /* ScalarType.SINT32 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 17 /* ScalarType.SINT32 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 17 /* ScalarType.SINT32 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 17 /* ScalarType.SINT32 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 17 /* ScalarType.SINT32 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 17 /* ScalarType.SINT32 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 17 /* ScalarType.SINT32 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * SInt64Rules describes the constraints applied to `sint64` values
 *
 * @generated from message validate.SInt64Rules
 */
export const SInt64Rules = proto2.makeMessageType(
  "validate.SInt64Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 18 /* ScalarType.SINT64 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 18 /* ScalarType.SINT64 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 18 /* ScalarType.SINT64 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 18 /* ScalarType.SINT64 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 18 /* ScalarType.SINT64 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 18 /* ScalarType.SINT64 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 18 /* ScalarType.SINT64 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * Fixed32Rules describes the constraints applied to `fixed32` values
 *
 * @generated from message validate.Fixed32Rules
 */
export const Fixed32Rules = proto2.makeMessageType(
  "validate.Fixed32Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 7 /* ScalarType.FIXED32 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 7 /* ScalarType.FIXED32 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 7 /* ScalarType.FIXED32 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 7 /* ScalarType.FIXED32 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 7 /* ScalarType.FIXED32 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 7 /* ScalarType.FIXED32 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 7 /* ScalarType.FIXED32 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * Fixed64Rules describes the constraints applied to `fixed64` values
 *
 * @generated from message validate.Fixed64Rules
 */
export const Fixed64Rules = proto2.makeMessageType(
  "validate.Fixed64Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 6 /* ScalarType.FIXED64 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 6 /* ScalarType.FIXED64 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 6 /* ScalarType.FIXED64 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 6 /* ScalarType.FIXED64 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 6 /* ScalarType.FIXED64 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 6 /* ScalarType.FIXED64 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 6 /* ScalarType.FIXED64 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * SFixed32Rules describes the constraints applied to `sfixed32` values
 *
 * @generated from message validate.SFixed32Rules
 */
export const SFixed32Rules = proto2.makeMessageType(
  "validate.SFixed32Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 15 /* ScalarType.SFIXED32 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 15 /* ScalarType.SFIXED32 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 15 /* ScalarType.SFIXED32 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 15 /* ScalarType.SFIXED32 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 15 /* ScalarType.SFIXED32 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 15 /* ScalarType.SFIXED32 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 15 /* ScalarType.SFIXED32 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * SFixed64Rules describes the constraints applied to `sfixed64` values
 *
 * @generated from message validate.SFixed64Rules
 */
export const SFixed64Rules = proto2.makeMessageType(
  "validate.SFixed64Rules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 16 /* ScalarType.SFIXED64 */, opt: true },
    { no: 2, name: "lt", kind: "scalar", T: 16 /* ScalarType.SFIXED64 */, opt: true },
    { no: 3, name: "lte", kind: "scalar", T: 16 /* ScalarType.SFIXED64 */, opt: true },
    { no: 4, name: "gt", kind: "scalar", T: 16 /* ScalarType.SFIXED64 */, opt: true },
    { no: 5, name: "gte", kind: "scalar", T: 16 /* ScalarType.SFIXED64 */, opt: true },
    { no: 6, name: "in", kind: "scalar", T: 16 /* ScalarType.SFIXED64 */, repeated: true },
    { no: 7, name: "not_in", kind: "scalar", T: 16 /* ScalarType.SFIXED64 */, repeated: true },
    { no: 8, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * BoolRules describes the constraints applied to `bool` values
 *
 * @generated from message validate.BoolRules
 */
export const BoolRules = proto2.makeMessageType(
  "validate.BoolRules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * StringRules describe the constraints applied to `string` values
 *
 * @generated from message validate.StringRules
 */
export const StringRules = proto2.makeMessageType(
  "validate.StringRules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 19, name: "len", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 2, name: "min_len", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 3, name: "max_len", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 20, name: "len_bytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 4, name: "min_bytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 5, name: "max_bytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 6, name: "pattern", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "prefix", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 8, name: "suffix", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 9, name: "contains", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 23, name: "not_contains", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 10, name: "in", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 11, name: "not_in", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 12, name: "email", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 13, name: "hostname", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 14, name: "ip", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 15, name: "ipv4", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 16, name: "ipv6", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 17, name: "uri", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 18, name: "uri_ref", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 21, name: "address", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 22, name: "uuid", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 24, name: "well_known_regex", kind: "enum", T: proto2.getEnumType(KnownRegex), oneof: "well_known" },
    { no: 25, name: "strict", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true, default: true },
    { no: 26, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * BytesRules describe the constraints applied to `bytes` values
 *
 * @generated from message validate.BytesRules
 */
export const BytesRules = proto2.makeMessageType(
  "validate.BytesRules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
    { no: 13, name: "len", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 2, name: "min_len", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 3, name: "max_len", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 4, name: "pattern", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "prefix", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
    { no: 6, name: "suffix", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
    { no: 7, name: "contains", kind: "scalar", T: 12 /* ScalarType.BYTES */, opt: true },
    { no: 8, name: "in", kind: "scalar", T: 12 /* ScalarType.BYTES */, repeated: true },
    { no: 9, name: "not_in", kind: "scalar", T: 12 /* ScalarType.BYTES */, repeated: true },
    { no: 10, name: "ip", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 11, name: "ipv4", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 12, name: "ipv6", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "well_known" },
    { no: 14, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * EnumRules describe the constraints applied to enum values
 *
 * @generated from message validate.EnumRules
 */
export const EnumRules = proto2.makeMessageType(
  "validate.EnumRules",
  () => [
    { no: 1, name: "const", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 2, name: "defined_only", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 3, name: "in", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 4, name: "not_in", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * MessageRules describe the constraints applied to embedded message values.
 * For message-type fields, validation is performed recursively.
 *
 * @generated from message validate.MessageRules
 */
export const MessageRules = proto2.makeMessageType(
  "validate.MessageRules",
  () => [
    { no: 1, name: "skip", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 2, name: "required", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * RepeatedRules describe the constraints applied to `repeated` values
 *
 * @generated from message validate.RepeatedRules
 */
export const RepeatedRules = proto2.makeMessageType(
  "validate.RepeatedRules",
  () => [
    { no: 1, name: "min_items", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 2, name: "max_items", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 3, name: "unique", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 4, name: "items", kind: "message", T: FieldRules, opt: true },
    { no: 5, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * MapRules describe the constraints applied to `map` values
 *
 * @generated from message validate.MapRules
 */
export const MapRules = proto2.makeMessageType(
  "validate.MapRules",
  () => [
    { no: 1, name: "min_pairs", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 2, name: "max_pairs", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 3, name: "no_sparse", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 4, name: "keys", kind: "message", T: FieldRules, opt: true },
    { no: 5, name: "values", kind: "message", T: FieldRules, opt: true },
    { no: 6, name: "ignore_empty", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * AnyRules describe constraints applied exclusively to the
 * `google.protobuf.Any` well-known type
 *
 * @generated from message validate.AnyRules
 */
export const AnyRules = proto2.makeMessageType(
  "validate.AnyRules",
  () => [
    { no: 1, name: "required", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 2, name: "in", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "not_in", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * DurationRules describe the constraints applied exclusively to the
 * `google.protobuf.Duration` well-known type
 *
 * @generated from message validate.DurationRules
 */
export const DurationRules = proto2.makeMessageType(
  "validate.DurationRules",
  () => [
    { no: 1, name: "required", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 2, name: "const", kind: "message", T: Duration, opt: true },
    { no: 3, name: "lt", kind: "message", T: Duration, opt: true },
    { no: 4, name: "lte", kind: "message", T: Duration, opt: true },
    { no: 5, name: "gt", kind: "message", T: Duration, opt: true },
    { no: 6, name: "gte", kind: "message", T: Duration, opt: true },
    { no: 7, name: "in", kind: "message", T: Duration, repeated: true },
    { no: 8, name: "not_in", kind: "message", T: Duration, repeated: true },
  ],
);

/**
 * TimestampRules describe the constraints applied exclusively to the
 * `google.protobuf.Timestamp` well-known type
 *
 * @generated from message validate.TimestampRules
 */
export const TimestampRules = proto2.makeMessageType(
  "validate.TimestampRules",
  () => [
    { no: 1, name: "required", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 2, name: "const", kind: "message", T: Timestamp, opt: true },
    { no: 3, name: "lt", kind: "message", T: Timestamp, opt: true },
    { no: 4, name: "lte", kind: "message", T: Timestamp, opt: true },
    { no: 5, name: "gt", kind: "message", T: Timestamp, opt: true },
    { no: 6, name: "gte", kind: "message", T: Timestamp, opt: true },
    { no: 7, name: "lt_now", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 8, name: "gt_now", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 9, name: "within", kind: "message", T: Duration, opt: true },
  ],
);

