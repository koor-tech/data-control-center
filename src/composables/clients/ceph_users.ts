import { ConnectError } from '@connectrpc/connect';
import { User } from '~~/gen/ts/api/resources/ceph/v1/users_pb';
import {
    type CreateCephUserRequest,
    type CreateCephUserResponse,
    type DeleteCephUserRequest,
    type DeleteCephUserResponse,
} from '~~/gen/ts/api/services/ceph/v1/users_pb';

export async function listCephUsers(): Promise<User[]> {
    const { $grpc } = useNuxtApp();
    try {
        const resp = await $grpc.getCephUsersClient().listCephUsers({});

        return resp.cephUsers;
    } catch (e) {
        if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
        throw e;
    }
}

export async function createCephUser(
    username: string,
    name: string,
    email: string,
    password: string,
    rol: string,
): Promise<CreateCephUserResponse> {
    const { $grpc } = useNuxtApp();
    try {
        const createUser = {
            cephUser: {
                username,
                name,
                email,
                password,
                roles: [rol],
            },
        } as CreateCephUserRequest;
        const resp = await $grpc.getCephUsersClient().createCephUser(createUser);

        return resp;
    } catch (e) {
        if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
        throw e;
    }
}

export async function deleteCephUser(username: string): Promise<DeleteCephUserResponse> {
    const { $grpc } = useNuxtApp();
    try {
        const deleteUser = {
            username,
        } as DeleteCephUserRequest;
        return await $grpc.getCephUsersClient().deleteCephUser(deleteUser);
    } catch (e) {
        if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
        throw e;
    }
}
