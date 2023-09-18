<script setup lang="ts">
import { ResourceStatus } from '~~/gen/ts/api/resources/stats/stats_pb';
import { ResourceInfo } from '~~/gen/ts/api/resources/stats/stats_pb';

withDefaults(defineProps<{
  rows: ResourceInfo[];
  hideAPI?: boolean;
  hideKind?: boolean;
  hideNamespace?: boolean;
}>(), {
  hideAPI: true,
  hideKind: false,
  hideNamespace: false,
});
</script>

<template>
  <div class="px-4 sm:px-6 lg:px-8">
    <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h2 class="text-base font-semibold leading-6 text-gray-900">Resource Info</h2>
        <p class="mt-2 text-sm text-gray-700">A list of the Storage Cluster relevant resources.</p>
      </div>
    </div>
    <div class="mt-8 flow-root">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <table class="min-w-full divide-y divide-gray-300">
            <thead>
              <tr>
                <th v-if="!hideAPI" scope="col"
                  class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">API</th>
                <th v-if="!hideKind" scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Kind
                </th>
                <th v-if="!hideNamespace" scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                  Namespace</th>
                <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Name</th>
                <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Status</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="row in rows" :key="row.name">
                <td v-if="!hideAPI" class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">{{
                  row.apiversion
                }}
                </td>
                <td v-if="!hideKind" class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{
                  row.kind }}</td>
                <td v-if="!hideNamespace" class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ row.namespace }}
                </td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ row.name }}</td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{
                  ResourceStatus[row.status].replace('RESOURCE_', '') }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
