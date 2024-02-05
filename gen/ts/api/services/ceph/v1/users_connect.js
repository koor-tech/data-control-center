// @generated by protoc-gen-connect-es v1.3.0
// @generated from file api/services/ceph/v1/users.proto (package api.services.ceph.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateCephUserRequest, CreateCephUserResponse, DeleteCephUserRequest, DeleteCephUserResponse, ListCephUsersRequest, ListCephUsersResponse } from "./users_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.services.ceph.v1.CephUsersService
 */
export const CephUsersService = {
  typeName: "api.services.ceph.v1.CephUsersService",
  methods: {
    /**
     * @generated from rpc api.services.ceph.v1.CephUsersService.ListCephUsers
     */
    listCephUsers: {
      name: "ListCephUsers",
      I: ListCephUsersRequest,
      O: ListCephUsersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.ceph.v1.CephUsersService.CreateCephUser
     */
    createCephUser: {
      name: "CreateCephUser",
      I: CreateCephUserRequest,
      O: CreateCephUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.services.ceph.v1.CephUsersService.DeleteCephUser
     */
    deleteCephUser: {
      name: "DeleteCephUser",
      I: DeleteCephUserRequest,
      O: DeleteCephUserResponse,
      kind: MethodKind.Unary,
    },
  }
};

