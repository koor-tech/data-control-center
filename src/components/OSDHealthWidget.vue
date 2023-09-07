<template>
  <div class="px-4 sm:px-6 lg:px-8">
    <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h1 class="text-base font-semibold leading-6 text-gray-900">Pod Status</h1>
        <p class="mt-2 text-sm text-gray-700">A list of the Rook Ceph Pods.</p>
      </div>
    </div>
    <div class="mt-8 flow-root">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <table class="min-w-full divide-y divide-gray-300">
            <thead>
              <tr>
                <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">Name</th>
                <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Ready</th>
                <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Status</th>
                <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Restarts</th>
                <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Age</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="row in rows" :key="row.name">
                <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">{{ row.name }}
                </td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ row.ready }}</td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ row.status }}</td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ row.restarts }}</td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ row.age }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const rows = [
  {
    name: 'rook-ceph-mon1-dw8wr',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '17h',
  },
  {
    name: 'rook-ceph-mon1-kio3k',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '17h',
  },
  {
    name: 'rook-ceph-osd-cdefg',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '12h',
  },
  {
    name: 'rook-ceph-osd-bcdef',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '12h',
  },
  {
    name: 'rook-ceph-osd-abcde',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '8h',
  },
]
</script>
