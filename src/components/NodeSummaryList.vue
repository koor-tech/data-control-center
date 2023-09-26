<script setup lang="ts">
import NodeSummaryItem from './NodeSummaryItem.vue';
import Container from './partials/Container.vue';

const items = [
    {
        name: 'rack1-slots-1-3',
        description: '3 drives, 9 OSDs, 10 GiB Network',
        capacity: '180 GB',
        health: 0,
        lastPing: '5 min',
    },
    {
        name: 'rack1-slots-4-7',
        description: '4 drives, 12 OSDs, 10 GiB Network',
        capacity: '1 TB',
        health: 1,
        lastPing: '5 min',
    },
];
</script>

<template>
    <Container>
        <template v-slot:title>Storage Nodes</template>
        <template v-slot:default>
            <table class="min-w-full divide-y divide-gray-300">
                <thead>
                    <tr>
                        <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">Name</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Description</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Capacity</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Health</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Last Ping</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-200">
                    <NodeSummaryItem v-for="item in items" :key="item.name" :node="item" />
                </tbody>
            </table>
        </template>
    </Container>
</template>
