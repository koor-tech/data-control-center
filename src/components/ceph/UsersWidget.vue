<script setup lang="ts">
import { HumanMaleIcon } from 'mdi-vue3';
</script>

<template>
    <div class="overflow-hidden rounded-lg bg-white shadow">
        <div class="bg-white p-3">
            <div class="sm:flex sm:items-center sm:justify-between">
                <div class="sm:flex sm:space-x-1">
                    <div class="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left">
                        <p class="text-xl font-bold text-gray-900 sm:text-2xl"></p>
                        <div class="flex justify-between gap-x-4 py-4">
                            <dt class="text-gray-500">
                                <HumanMaleIcon class="h-9 w-9" aria-hidden="true" />
                            </dt>
                            <dd class="text-gray-900 pt-2">User management</dd>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="grid divide-gray-200 border-t border-gray-200 bg-gray-50">
            <div class="px-2 py-5 text-center text-sm font-medium">
                <span class="isolate inline-flex rounded-md shadow-sm">
                    <router-link to="/ceph/users/add">
                        <button
                            type="button"
                            class="relative inline-flex items-center rounded-l-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-10"
                        >
                            Add user
                        </button>
                    </router-link>

                    <router-link to="/ceph/users/">
                        <button
                            type="button"
                            class="relative -ml-px inline-flex items-center bg-white px-3 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-10"
                        >
                            View users
                        </button>
                    </router-link>
                </span>
            </div>
        </div>
    </div>
</template>
