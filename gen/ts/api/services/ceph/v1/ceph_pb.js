// @generated by protoc-gen-es v1.6.0
// @generated from file api/services/ceph/v1/ceph.proto (package api.services.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { User } from "../../../resources/ceph/v1/ceph_pb.js";

/**
 * @generated from message api.services.ceph.v1.GetCephUsersRequest
 */
export const GetCephUsersRequest = proto3.makeMessageType(
  "api.services.ceph.v1.GetCephUsersRequest",
  [],
);

/**
 * @generated from message api.services.ceph.v1.GetCephUsersResponse
 */
export const GetCephUsersResponse = proto3.makeMessageType(
  "api.services.ceph.v1.GetCephUsersResponse",
  () => [
    { no: 1, name: "ceph_users", kind: "message", T: User, repeated: true },
  ],
);

/**
 * @generated from message api.services.ceph.v1.CreateCephUserRequest
 */
export const CreateCephUserRequest = proto3.makeMessageType(
  "api.services.ceph.v1.CreateCephUserRequest",
  () => [
    { no: 1, name: "ceph_user", kind: "message", T: User },
  ],
);

/**
 * @generated from message api.services.ceph.v1.CreateCephUserResponse
 */
export const CreateCephUserResponse = proto3.makeMessageType(
  "api.services.ceph.v1.CreateCephUserResponse",
  () => [
    { no: 1, name: "ceph_user", kind: "message", T: User },
  ],
);

/**
 * @generated from message api.services.ceph.v1.DeleteCephUserRequest
 */
export const DeleteCephUserRequest = proto3.makeMessageType(
  "api.services.ceph.v1.DeleteCephUserRequest",
  () => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.services.ceph.v1.DeleteCephUserResponse
 */
export const DeleteCephUserResponse = proto3.makeMessageType(
  "api.services.ceph.v1.DeleteCephUserResponse",
  () => [
    { no: 1, name: "status", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

