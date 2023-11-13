<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';

import { ref } from 'vue';
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';
import { CloseIcon } from 'mdi-vue3';
import { useCephStore } from '~/store/ceph';
import GenericBadge from '~/components/partials/GenericBadge.vue';

useHead({
    title: 'MGR Dashboard Users',
});
definePageMeta({
    title: 'MGR Dashboard Users',
    requiresAuth: true,
});

const modalWindowOpen = ref(false);
const username = ref();

const { $grpc } = useNuxtApp();
const cephStore = useCephStore();

const { data: cephUsers } = useLazyAsyncData('cephUsers', async () => {
    try {
        return await cephStore.getCephUsers();
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});

const displayConfirmation = function (Username: string) {
    modalWindowOpen.value = true;
    username.value = Username;
};

const deleteCephUser = async function (): Promise<void> {
    try {
        await cephStore.deleteUser(username.value);
        console.log('removing username.value', username.value);
        modalWindowOpen.value = false;
        cephUsers.value = await cephStore.getCephUsers();
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
};
</script>
<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Ceph Users</h2>
        <div class="px-4 sm:px-6 lg:px-8">
            <div class="sm:flex sm:items-center">
                <div class="sm:flex-auto">
                    <h1 class="text-base font-semibold leading-6 text-gray-900">Users</h1>
                    <p class="mt-2 text-sm text-gray-700">A list of all the ceph users</p>
                </div>
                <div class="mt-4 sm:ml-16 sm:mt-0 sm:flex-none">
                    <NuxtLink :to="{ name: 'controls-cephusers-add' }">
                        <button
                            type="button"
                            class="block rounded-md bg-indigo-600 px-3 py-2 text-center text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                        >
                            Add user
                        </button>
                    </NuxtLink>
                </div>
            </div>
        </div>
        <div class="mt-8 flow-root">
            <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                    <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg">
                        <table class="min-w-full divide-y divide-gray-300">
                            <thead class="bg-gray-50">
                                <tr>
                                    <th
                                        scope="col"
                                        class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6"
                                    >
                                        Username
                                    </th>
                                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Roles</th>
                                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                                        Enabled
                                    </th>
                                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                                        Delete
                                    </th>
                                </tr>
                            </thead>
                            <tbody class="divide-y divide-gray-200 bg-white">
                                <tr v-for="{ Enabled, Roles, Username } in cephUsers" :key="Username">
                                    <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">
                                        {{ Username }}
                                    </td>
                                    <td
                                        v-for="role in Roles"
                                        :key="role"
                                        class="whitespace-nowrap px-3 py-4 text-sm text-gray-500"
                                    >
                                        <GenericBadge>{{ role }}</GenericBadge>
                                    </td>
                                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                                        {{ Enabled }}
                                    </td>
                                    <td class="relative whitespace-nowrap py-4 pl-3 text-sm font-medium sm:pr-6">
                                        <a
                                            href="#"
                                            class="text-indigo-600 hover:text-indigo-900"
                                            @click="displayConfirmation(Username)"
                                            >Delete<span class="sr-only"></span
                                        ></a>
                                    </td>
                                </tr>
                            </tbody>
                        </table>

                        <TransitionRoot as="template" :show="modalWindowOpen">
                            <Dialog as="div" class="relative z-10" @close="modalWindowOpen = false">
                                <TransitionChild
                                    as="template"
                                    enter="ease-out duration-300"
                                    enter-from="opacity-0"
                                    enter-to="opacity-100"
                                    leave="ease-in duration-200"
                                    leave-from="opacity-100"
                                    leave-to="opacity-0"
                                >
                                    <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
                                </TransitionChild>

                                <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
                                    <div
                                        class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"
                                    >
                                        <TransitionChild
                                            as="template"
                                            enter="ease-out duration-300"
                                            enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                                            enter-to="opacity-100 translate-y-0 sm:scale-100"
                                            leave="ease-in duration-200"
                                            leave-from="opacity-100 translate-y-0 sm:scale-100"
                                            leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                                        >
                                            <DialogPanel
                                                class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6"
                                            >
                                                <div>
                                                    <div
                                                        class="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-red-100"
                                                    >
                                                        <CloseIcon class="h-6 w-6 text-red-600" aria-hidden="true" />
                                                    </div>
                                                    <div class="mt-3 text-center sm:mt-5">
                                                        <DialogTitle
                                                            as="h3"
                                                            class="text-base font-semibold leading-6 text-gray-900"
                                                            >Delete Ceph Dashboard User</DialogTitle
                                                        >
                                                        <div class="mt-2">
                                                            <p class="text-sm text-gray-500">
                                                                Are you sure to delete the user from ceph dashboard?
                                                            </p>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div
                                                    class="mt-5 sm:mt-6 sm:grid sm:grid-flow-row-dense sm:grid-cols-2 sm:gap-3"
                                                >
                                                    <button
                                                        type="button"
                                                        class="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 sm:col-start-2"
                                                        @click="deleteCephUser()"
                                                    >
                                                        Delete
                                                    </button>
                                                    <button
                                                        ref="cancelButtonRef"
                                                        type="button"
                                                        class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:col-start-1 sm:mt-0"
                                                        @click="modalWindowOpen = false"
                                                    >
                                                        Cancel
                                                    </button>
                                                </div>
                                            </DialogPanel>
                                        </TransitionChild>
                                    </div>
                                </div>
                            </Dialog>
                        </TransitionRoot>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
