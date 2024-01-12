// @generated by protoc-gen-connect-es v0.13.1
// @generated from file api/services/ceph/v1/users.proto (package api.services.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateCephUserRequest, CreateCephUserResponse, DeleteCephUserRequest, DeleteCephUserResponse, ListCephUsersRequest, ListCephUsersResponse } from "./users_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.services.ceph.v1.CephUsersService
 */
export declare const CephUsersService: {
  readonly typeName: "api.services.ceph.v1.CephUsersService",
  readonly methods: {
    /**
     * @generated from rpc api.services.ceph.v1.CephUsersService.ListCephUsers
     */
    readonly listCephUsers: {
      readonly name: "ListCephUsers",
      readonly I: typeof ListCephUsersRequest,
      readonly O: typeof ListCephUsersResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.ceph.v1.CephUsersService.CreateCephUser
     */
    readonly createCephUser: {
      readonly name: "CreateCephUser",
      readonly I: typeof CreateCephUserRequest,
      readonly O: typeof CreateCephUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.ceph.v1.CephUsersService.DeleteCephUser
     */
    readonly deleteCephUser: {
      readonly name: "DeleteCephUser",
      readonly I: typeof DeleteCephUserRequest,
      readonly O: typeof DeleteCephUserResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};
