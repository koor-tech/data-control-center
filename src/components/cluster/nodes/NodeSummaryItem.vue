<script setup lang="ts">
import { resourceStatusToTextColor } from '~/components/helpers';
import { NodeInfo, ResourceStatus } from '~~/gen/ts/api/resources/k8s/v1/resources_pb';

defineProps<{
    node: NodeInfo;
}>();
</script>

<template>
    <tr>
        <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-normal text-gray-900 sm:pl-0">{{ node.name }}</td>
        <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500" :class="resourceStatusToTextColor(node.status)">
            {{ ResourceStatus[node.status] }}
        </td>
        <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ node.roles.join(', ') }}</td>
        <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ node.internalIp ?? 'N/A' }}</td>
        <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ node.externalIp ?? 'N/A' }}</td>
    </tr>
</template>
