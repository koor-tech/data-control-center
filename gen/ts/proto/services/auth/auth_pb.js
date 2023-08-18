// @generated by protoc-gen-es v1.3.0
// @generated from file proto/services/auth/auth.proto (package services.auth, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Timestamp } from "../../resources/timestamp/timestamp_pb.js";

/**
 * @generated from message services.auth.LoginRequest
 */
export const LoginRequest = proto3.makeMessageType(
  "services.auth.LoginRequest",
  () => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message services.auth.LoginResponse
 */
export const LoginResponse = proto3.makeMessageType(
  "services.auth.LoginResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expires", kind: "message", T: Timestamp },
    { no: 3, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * @generated from message services.auth.LogoutRequest
 */
export const LogoutRequest = proto3.makeMessageType(
  "services.auth.LogoutRequest",
  [],
);

/**
 * @generated from message services.auth.LogoutResponse
 */
export const LogoutResponse = proto3.makeMessageType(
  "services.auth.LogoutResponse",
  () => [
    { no: 1, name: "success", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

