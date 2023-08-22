// @generated by protoc-gen-es v1.3.0
// @generated from file stats.proto (package services.cluster, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message as Message$1, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message services.cluster.Message
 */
export declare class Message extends Message$1<Message> {
  /**
   * @generated from field: string body = 1;
   */
  body: string;

  constructor(data?: PartialMessage<Message>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "services.cluster.Message";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Message;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Message;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Message;

  static equals(a: Message | PlainMessage<Message> | undefined, b: Message | PlainMessage<Message> | undefined): boolean;
}

