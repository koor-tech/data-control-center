import { ConnectError } from '@connectrpc/connect';
import { defineStore, type StoreDefinition } from 'pinia';
import { User } from '~~/gen/ts/api/resources/ceph/v1/ceph_pb';
import {
    DeleteCephUserRequest,
    DeleteCephUserResponse,
    type CreateCephUserRequest,
    type CreateCephUserResponse,
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
        async listCephUsers(): Promise<User[]> {
            const { $grpc } = useNuxtApp();
            try {
                const resp = await $grpc.listCephUsers().listCephUsers({});

                return resp.cephUsers;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        createCephUser: async function (
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
                const resp = await $grpc.listCephUsers().createCephUser(createUser);

                return resp;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        async deleteCephUser(username: string): Promise<DeleteCephUserResponse> {
            const { $grpc } = useNuxtApp();
            try {
                const deleteUser = {
                    username,
                } as DeleteCephUserRequest;
                return await $grpc.listCephUsers().deleteCephUser(deleteUser);
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
