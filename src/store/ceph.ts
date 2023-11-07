import { ConnectError } from '@connectrpc/connect';
import { defineStore, type StoreDefinition } from 'pinia';
import { User } from '~~/gen/ts/api/resources/ceph/v1/ceph_pb';
import {
    DeleteCephUserRequest,
    DeleteCephUserResponse,
    type CreateCephUsersRequest,
    type CreateCephUsersResponse,
} from '~~/gen/ts/api/services/ceph/v1/ceph_pb';

export interface CephState {
    addError: null | string;
}

export const useCephStore = defineStore('ceph', {
    state: () =>
        ({
            addError: null,
        }) as CephState,
    persist: false,
    actions: {
        async getCephUsers(): Promise<User[]> {
            const { $grpc } = useNuxtApp();
            try {
                const resp = await $grpc.getCephUsers().getCephUsers({});

                return resp.cephUsers;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        addCephUser: async function (
            username: string,
            name: string,
            email: string,
            password: string,
            rol: string,
        ): Promise<CreateCephUsersResponse> {
            const { $grpc } = useNuxtApp();
            try {
                const createUser = {
                    cephUser: {
                        Username: username,
                        Name: name,
                        Email: email,
                        Password: password,
                        Roles: [rol],
                    },
                } as CreateCephUsersRequest;
                const resp = await $grpc.getCephUsers().createCephUsers(createUser);

                return resp;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        async deleteUser(username: string): Promise<DeleteCephUserResponse> {
            const { $grpc } = useNuxtApp();
            try {
                const deleteUser = {
                    username,
                } as DeleteCephUserRequest;
                return await $grpc.getCephUsers().deleteCephUser(deleteUser);
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useCephStore as unknown as StoreDefinition, import.meta.hot));
}
