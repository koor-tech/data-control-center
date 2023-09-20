<script setup lang="ts">
import { ArcElement, Chart as ChartJS, Legend, Tooltip } from 'chart.js';
import { ChartPieIcon } from 'mdi-vue3';
import { Doughnut } from 'vue-chartjs';

ChartJS.register(ArcElement, Tooltip, Legend);

const usage = {
    used: 3.3,
    available: 117,
    total: 120,
};

const usedPercentage = (usage.used / usage.total) * 100;
const availablePercentage = (usage.available / usage.total) * 100;
const data = {
    labels: ['used', 'available'],
    datasets: [
        {
            backgroundColor: ['#E46651', '#41B883'],
            data: [usedPercentage, availablePercentage],
        },
    ],
};

const options = {
    responsive: true,
    maintainAspectRatio: false,
};

const props = defineProps<{
    statsContainer: any;
}>();

const { statsContainer } = toRefs(props);

const init = async () => {
    console.log(await props);
};
init();
</script>

<template>
    <div class="relative overflow-hidden rounded-lg bg-white px-4 pb-12 pt-5 shadow sm:px-6 sm:pt-6">
        <dt>
            <div
                :class="[
                    statsContainer.color ? 'bg-' + statsContainer.color + '-500' : 'bg-blue-500',
                    'absolute rounded-md p-3',
                ]"
            >
                <ChartPieIcon class="h-6 w-6 text-white" aria-hidden="true" />
            </div>
            <p class="ml-16 truncate text-sm font-medium text-gray-500">Usage</p>
        </dt>
        <dd class="flex items-baseline pb-1 sm:pb-1 pt-10">
            <div>
                <Doughnut :data="data" :options="options" />
            </div>
            <div class="absolute inset-x-0 bottom-0 bg-gray-50 px-4 py-4 sm:px-6">
                <div class="text-xs text-center">Used 3.3 GiB, available 117 GiB, total 120 GiB</div>
            </div>
        </dd>
    </div>
</template>
