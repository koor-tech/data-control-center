<script lang="ts" setup>
import type { ConnectError } from '@connectrpc/connect';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';

const { $grpc } = useNuxtApp();

const appConfig = useAppConfig();

const {
    data: scrubbingSchedule,
    error,
    refresh,
    pending,
} = useLazyAsyncData(`clusterRadar`, async () => {
    const response = await $grpc.getClusterClient().getScrubbingSchedule({});

    return response.osdScrubbingSchedule!;
});
watch(error, () => {
    if (error.value !== null) $grpc.handleError(error.value as ConnectError);
});

const canSubmit = ref(false);
</script>

<template>
    <div class="mx-auto">
        <div class="overflow-hidden bg-base-800 shadow sm:rounded-lg text-neutral">
            <div class="px-4 py-5 sm:px-6">
                <h3 class="text-base font-semibold leading-6">Custom OSD Scrubbing Schedule</h3>
                <p class="mt-1 max-w-2xl text-sm">Set a custom OSD scrubbing schedule here.</p>
            </div>
            <div class="border-t border-base-400 px-4 py-5 sm:p-0">
                <DataErrorBlock v-if="error" :retry="refresh" :message="error.message" />
                <DataPendingBlock v-else-if="pending" message="Loading Scrubbing Schedule" />
                <DataNoDataBlock v-else-if="scrubbingSchedule === null || scrubbingSchedule === undefined" />

                <template v-else>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-medium">Maximumg Scrubbing Operations</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <VeeField
                                    v-model="scrubbingSchedule.maxScrubOps"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="maxScrubOps"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                    min="1"
                                    max="99999"
                                />
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-medium">Run Hours</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <label for="beginHour"> Begin Hour </label>
                                <VeeField
                                    v-model="scrubbingSchedule.beginHour"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="beginHour"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                    min="1"
                                    max="99999"
                                />
                                <label for="endHour"> End Hour </label>
                                <VeeField
                                    v-model="scrubbingSchedule.endHour"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="endHour"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                    min="1"
                                    max="99999"
                                />
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-medium">Run Week Days</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <label for="beginWeekDay"> Begin Week Day </label>
                                <VeeField
                                    v-model="scrubbingSchedule.beginWeekDay"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="beginWeekDay"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                    min="1"
                                    max="99999"
                                />
                                <label for="endWeekDay"> End Week Day </label>
                                <VeeField
                                    v-model="scrubbingSchedule.endWeekDay"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="endWeekDay"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                    min="1"
                                    max="99999"
                                />
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-medium">Scrub Intervals</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <label for="minScrubInterval"> Min Scrub Interval </label>
                                <VeeField
                                    v-model="scrubbingSchedule.minScrubInterval"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="minScrubInterval"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                                <label for="maxScrubInterval"> Max Scrub Interval </label>
                                <VeeField
                                    v-model="scrubbingSchedule.maxScrubInterval"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="maxScrubInterval"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                                <label for="deepScrubInterval"> Deep Scrub Interval </label>
                                <VeeField
                                    v-model="scrubbingSchedule.deepScrubInterval"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="deepScrubInterval"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-medium">Scrub Sleep Seconds</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <VeeField
                                    v-model="scrubbingSchedule.scrubSleepSeconds"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="scrubSleepSeconds"
                                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </dd>
                        </div>
                    </dl>
                </template>
            </div>
        </div>
    </div>
</template>
