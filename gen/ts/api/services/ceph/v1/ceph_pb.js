// @generated by protoc-gen-es v1.4.1
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
    { no: 1, name: "cephUsers", kind: "message", T: User, repeated: true },
  ],
);

/**
 * @generated from message api.services.ceph.v1.CreatCephUsersRequest
 */
export const CreatCephUsersRequest = proto3.makeMessageType(
  "api.services.ceph.v1.CreatCephUsersRequest",
  () => [
    { no: 1, name: "cephUser", kind: "message", T: User },
  ],
);

/**
 * @generated from message api.services.ceph.v1.CephUsersResponse
 */
export const CephUsersResponse = proto3.makeMessageType(
  "api.services.ceph.v1.CephUsersResponse",
  () => [
    { no: 1, name: "cephUser", kind: "message", T: User },
  ],
);

