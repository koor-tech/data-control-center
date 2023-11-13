<script lang="ts" setup>
import { useNotificationsStore } from '~/store/notifications';
import { OSDScrubbingSchedule } from '~~/gen/ts/api/resources/koor/v1/koor_pb';

// TODO we should just rely on the defaults from the CRD/API
const osdScrubbingSchedule = ref<OSDScrubbingSchedule>(
    new OSDScrubbingSchedule({
        applySchedule: false,
        maxScrubOps: 3n,
        beginHour: 0n,
        endHour: 3n,
        beginWeekDay: 1n,
        endWeekDay: 0n,
        minScrubInterval: '24h',
        maxScrubInterval: '168h',
        deepScrubInterval: '168h',
        scrubSleepSeconds: '100ms',
    }),
);

const notifications = useNotificationsStore();

function sendKSDNotification(): void {
    notifications.dispatchNotification({
        title: 'Koor Storage Distribution required!',
        content: 'Please contact Koor to learn how to activate this feature in your cluster.',
        type: 'info',
    });
}

const canSubmit = ref(false);

watch(osdScrubbingSchedule.value, () => (osdScrubbingSchedule.value.applySchedule = false));
</script>

<template>
    <div class="mx-auto">
        <div class="overflow-hidden bg-base-800 shadow sm:rounded-lg text-neutral">
            <div class="px-4 py-5 sm:px-6">
                <h3 class="text-base font-semibold leading-6">Custom OSD Scrubbing Schedule</h3>
                <p class="mt-1 max-w-2xl text-sm">Set a custom OSD scrubbing schedule here.</p>
            </div>
            <div class="border-t border-base-400 px-4 py-5 sm:p-0">
                <dl class="sm:divide-y sm:divide-base-400">
                    <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                        <dt class="text-sm font-medium">Enable custom Schedule</dt>
                        <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                            <input
                                v-model="osdScrubbingSchedule.applySchedule"
                                type="checkbox"
                                @click="sendKSDNotification()"
                            />
                        </dd>
                    </div>
                </dl>
                <dl class="sm:divide-y sm:divide-base-400">
                    <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                        <dt class="text-sm font-medium">Maximumg Scrubbing Operations</dt>
                        <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                            <VeeField
                                v-model="osdScrubbingSchedule.maxScrubOps"
                                :disabled="!canSubmit"
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
                                v-model="osdScrubbingSchedule.beginHour"
                                :disabled="!canSubmit"
                                type="number"
                                name="beginHour"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                min="1"
                                max="99999"
                            />
                            <label for="endHour"> End Hour </label>
                            <VeeField
                                v-model="osdScrubbingSchedule.endHour"
                                :disabled="!canSubmit"
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
                                v-model="osdScrubbingSchedule.beginWeekDay"
                                :disabled="!canSubmit"
                                type="number"
                                name="beginWeekDay"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                min="1"
                                max="99999"
                            />
                            <label for="endWeekDay"> End Week Day </label>
                            <VeeField
                                v-model="osdScrubbingSchedule.endWeekDay"
                                :disabled="!canSubmit"
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
                                v-model="osdScrubbingSchedule.minScrubInterval"
                                :disabled="!canSubmit"
                                type="text"
                                name="minScrubInterval"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                            <label for="maxScrubInterval"> Max Scrub Interval </label>
                            <VeeField
                                v-model="osdScrubbingSchedule.maxScrubInterval"
                                :disabled="!canSubmit"
                                type="text"
                                name="maxScrubInterval"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                            <label for="deepScrubInterval"> Deep Scrub Interval </label>
                            <VeeField
                                v-model="osdScrubbingSchedule.deepScrubInterval"
                                :disabled="!canSubmit"
                                type="text"
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
                                v-model="osdScrubbingSchedule.scrubSleepSeconds"
                                :disabled="!canSubmit"
                                type="text"
                                name="scrubSleepSeconds"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                        </dd>
                    </div>
                </dl>
            </div>
        </div>
    </div>
</template>
