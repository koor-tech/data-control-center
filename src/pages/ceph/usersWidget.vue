<script setup lang="ts">
import {UsersIcon} from "@heroicons/vue/20/solid";
</script>
<template>
    <li class="overflow-hidden rounded-xl border border-gray-200">
        <div class="flex items-center gap-x-4 border-b border-gray-900/5 bg-gray-50 p-6">
            <UsersIcon aria-hidden="true" class="h-12 w-12 flex-none rounded-lg bg-white object-cover ring-1 ring-gray-900/10" />
            <div class="text-sm font-medium leading-6 text-gray-900">Users Management</div>
        </div>
        <dl class="-my-3 divide-y divide-gray-100 px-6 py-4 text-sm leading-6">
            <div class="flex gap-x-4 py-3">
                <dd class="text-gray-700">
                    <router-link to="/ceph/users/">
                        <button type="button" class="rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
                            Ceph Users
                        </button>
                    </router-link>
                </dd>
                <dd class="text-gray-700">
                    <router-link to="/ceph/users/add">
                        <button type="button" class="rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
                            Add Ceph User
                        </button>
                    </router-link>
                </dd>
            </div>

        </dl>
    </li>
</template>
