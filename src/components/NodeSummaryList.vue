<script setup lang="ts">
import { NodeInfo } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import NodeSummaryItem from './NodeSummaryItem.vue';
import Container from './partials/Container.vue';

defineProps<{
    nodes: NodeInfo[];
}>();
</script>

<template>
    <Container>
        <template v-slot:title>Storage Nodes</template>
        <template v-slot:default>
            <table class="min-w-full divide-y divide-gray-300">
                <thead>
                    <tr>
                        <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0">Name</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Status</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Roles</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Internal IP</th>
                        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">External IP</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-200">
                    <NodeSummaryItem v-for="item in nodes" :key="item.name" :node="item" />
                </tbody>
            </table>
        </template>
    </Container>
</template>
