import { ConnectError} from "@connectrpc/connect";
import { defineStore } from "pinia";
import {CreatCephUsersRequest} from "~~/gen/ts/api/services/ceph/v1/ceph_pb";
import {User} from "~~/gen/ts/api/resources/ceph/v1/ceph_pb";

export interface CephState {
    addError: null | string;
}

export const useCephStore = defineStore('ceph', {
    state: () => ({
        addError: null,
    }) as CephState,
    persist: false,
    actions: {
        async getCephUsers(): Promise<User[]> {
            return new Promise(async (res, rej) => {
                const { $grpc } = useNuxtApp();
                try {
                    const resp = await $grpc.getCephUsers().getCephUsers({});

                    return res(resp.cephUsers);
                } catch (e) {
                    if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                    return rej(e);
                }
            })
        },
        addCephUser: async function (username: string, name: string, email: string, password: string, rol: string): Promise<void> {
            return new Promise(async (res, rej) => {

                const {$grpc} = useNuxtApp();
                try {
                    const createUser = {
                        cephUser: {
                            Username: username,
                            Name: name,
                            Email: email,
                            Password: password,
                            Roles: [rol],
                        }
                    } as CreatCephUsersRequest;
                    const resp = await $grpc.getCephUsers().createCephUsers(createUser);
                    //await navigateTo({ name: 'ceph/users/' });

                    return res();

                } catch (e) {
                    if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                    return rej(e);
                }
            })
        }
    }
})
