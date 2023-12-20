// @generated by protoc-gen-es v1.6.0
// @generated from file api/services/auth/v1/auth.proto (package api.services.auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message api.services.auth.v1.LoginRequest
 */
export const LoginRequest = proto3.makeMessageType(
  "api.services.auth.v1.LoginRequest",
  () => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.services.auth.v1.LoginResponse
 */
export const LoginResponse = proto3.makeMessageType(
  "api.services.auth.v1.LoginResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expires", kind: "message", T: Timestamp },
    { no: 3, name: "account_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.services.auth.v1.LogoutRequest
 */
export const LogoutRequest = proto3.makeMessageType(
  "api.services.auth.v1.LogoutRequest",
  [],
);

/**
 * @generated from message api.services.auth.v1.LogoutResponse
 */
export const LogoutResponse = proto3.makeMessageType(
  "api.services.auth.v1.LogoutResponse",
  () => [
    { no: 1, name: "success", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message api.services.auth.v1.CheckTokenRequest
 */
export const CheckTokenRequest = proto3.makeMessageType(
  "api.services.auth.v1.CheckTokenRequest",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.services.auth.v1.CheckTokenResponse
 */
export const CheckTokenResponse = proto3.makeMessageType(
  "api.services.auth.v1.CheckTokenResponse",
  () => [
    { no: 1, name: "success", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

