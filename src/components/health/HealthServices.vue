<template>
    <HealthUsage v-if="statsContainer.title == 'usage'" :statsContainer="statsContainer.description" />

    <div v-else class="relative overflow-hidden rounded-lg bg-white px-4 pb-12 pt-5 shadow sm:px-6 sm:pt-6">
        <dt>
            <div :class="[
                statsContainer.color ? 'bg-' + statsContainer.color + '-500' : 'bg-blue-500',
                'absolute rounded-md p-3']">

                <component :is="currentIcon" class="h-6 w-6 text-white"
                           aria-hidden="true"/>
            </div>
            <p class="ml-16 truncate text-sm font-medium text-gray-500">{{ statsContainer.title }}</p>
        </dt>
        <dd class="flex items-baseline pb-1 sm:pb-1 pt-10">
            <ul role="list" class="w-full divide-y divide-gray-100">
                <HealthLiItem v-for="healthStats in statsContainer.description" :healthStats="healthStats" />
            </ul>
            <div class="absolute inset-x-0 bottom-0 bg-gray-50 px-4 py-4 sm:px-6">
                <div class="text-sm text-center">
                    <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">View all</a>
                </div>
            </div>
            <div class="absolute inset-x-0 bottom-0 bg-gray-50 px-4 py-4 sm:px-6">
                <div class="text-sm text-center">
                    <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">View all</a>
                </div>
            </div>
        </dd>
    </div>
</template>
<script setup lang="ts">
import {
    ArchiveBoxArrowDownIcon,
    ArchiveBoxIcon,
    ArrowsUpDownIcon,
    CircleStackIcon,
    ClipboardIcon,
    ComputerDesktopIcon,
    CubeIcon,
    DocumentTextIcon,
    ShieldExclamationIcon,
    Square3Stack3DIcon,
    TableCellsIcon,
} from '@heroicons/vue/20/solid';
const props = defineProps<{
    statsContainer: any;
}>();


const currentIcon = computed(() => {
    const iconName = props.statsContainer.icon;
    console.log(props.statsContainer.description)
    return {
        'ArchiveBoxArrowDownIcon': ArchiveBoxArrowDownIcon,
        'ArchiveBoxIcon': ArchiveBoxIcon,
        'ArrowsUpDownIcon': ArrowsUpDownIcon,
        'CircleStackIcon': CircleStackIcon,
        'ClipboardIcon': ClipboardIcon,
        'ComputerDesktopIcon': ComputerDesktopIcon,
        'CubeIcon': CubeIcon,
        'DocumentTextIcon': DocumentTextIcon,
        'ShieldExclamationIcon': ShieldExclamationIcon,
        'Square3Stack3DIcon': Square3Stack3DIcon,
        'TableCellsIcon': TableCellsIcon,
    }[iconName] || null;
});
</script>
