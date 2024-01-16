<script lang="ts" setup>
import type { ConnectError } from '@connectrpc/connect';
import { Listbox, ListboxButton, ListboxOption, ListboxOptions } from '@headlessui/vue';
import { CheckIcon, ChevronDownIcon } from 'mdi-vue3';
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
    try {
        const response = await $grpc.getControlsOSDsClient().getScrubbingSchedule({});

        return response.osdScrubbingSchedule!;
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});

async function setScrubbingSchedule() {
    if (!scrubbingSchedule.value) {
        return;
    }

    try {
        const response = await $grpc.getControlsOSDsClient().setScrubbingSchedule({
            osdScrubbingSchedule: scrubbingSchedule.value,
        });

        return response.osdScrubbingSchedule!;
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}

const weekDays = [
    { name: 'Monday', value: 0 },
    { name: 'Tuesday', value: 1 },
    { name: 'Wednesday', value: 2 },
    { name: 'Thursday', value: 3 },
    { name: 'Friday', value: 4 },
    { name: 'Saturday', value: 5 },
    { name: 'Sunday', value: 6 },
];

const canSubmit = ref(true);
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
                    <form @submit.prevent="setScrubbingSchedule">
                        <dl class="sm:divide-y sm:divide-base-400">
                            <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                                <dt class="text-sm font-medium">Maximumg Scrubbing Operations</dt>
                                <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                    <!-- TODO -->
                                    <VeeField
                                        v-model="scrubbingSchedule.maxScrubOps"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                        type="number"
                                        name="maxScrubOps"
                                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                        min="1"
                                        max="999999"
                                    />
                                </dd>
                            </div>
                        </dl>
                        <dl class="sm:divide-y sm:divide-base-400">
                            <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                                <dt class="text-sm font-medium">Run Hours</dt>
                                <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                    <label for="beginHour"> Begin Hour </label>
                                    <Listbox
                                        v-model="scrubbingSchedule.beginHour"
                                        as="div"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                    >
                                        <div class="relative">
                                            <ListboxButton
                                                class="block pl-3 text-left w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                            >
                                                <span class="block truncate">
                                                    {{ scrubbingSchedule.beginHour }}
                                                </span>
                                                <span
                                                    class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none"
                                                >
                                                    <ChevronDownIcon class="w-5 h-5 text-gray-400" aria-hidden="true" />
                                                </span>
                                            </ListboxButton>

                                            <transition
                                                leave-active-class="transition duration-100 ease-in"
                                                leave-from-class="opacity-100"
                                                leave-to-class="opacity-0"
                                            >
                                                <ListboxOptions
                                                    class="absolute z-10 w-full py-1 mt-1 overflow-auto text-base rounded-md bg-base-700 max-h-44 sm:text-sm"
                                                >
                                                    <ListboxOption
                                                        v-for="hour in 24"
                                                        :key="hour"
                                                        v-slot="{ active, selected }"
                                                        as="template"
                                                        :value="hour - 1"
                                                    >
                                                        <li
                                                            :class="[
                                                                active ? 'bg-primary-500' : '',
                                                                'text-neutral relative cursor-default select-none py-2 pl-8 pr-4',
                                                            ]"
                                                        >
                                                            <span
                                                                :class="[
                                                                    selected ? 'font-semibold' : 'font-normal',
                                                                    'block truncate',
                                                                ]"
                                                            >
                                                                {{ hour }}
                                                            </span>

                                                            <span
                                                                v-if="selected"
                                                                :class="[
                                                                    active ? 'text-neutral' : 'text-primary-500',
                                                                    'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                                                ]"
                                                            >
                                                                <CheckIcon class="w-5 h-5" aria-hidden="true" />
                                                            </span>
                                                        </li>
                                                    </ListboxOption>
                                                </ListboxOptions>
                                            </transition>
                                        </div>
                                    </Listbox>

                                    <label for="endHour"> End Hour </label>
                                    <Listbox
                                        v-model="scrubbingSchedule.endHour"
                                        as="div"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                    >
                                        <div class="relative">
                                            <ListboxButton
                                                class="block pl-3 text-left w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                            >
                                                <span class="block truncate">
                                                    {{ scrubbingSchedule.endHour }}
                                                </span>
                                                <span
                                                    class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none"
                                                >
                                                    <ChevronDownIcon class="w-5 h-5 text-gray-400" aria-hidden="true" />
                                                </span>
                                            </ListboxButton>

                                            <transition
                                                leave-active-class="transition duration-100 ease-in"
                                                leave-from-class="opacity-100"
                                                leave-to-class="opacity-0"
                                            >
                                                <ListboxOptions
                                                    class="absolute z-10 w-full py-1 mt-1 overflow-auto text-base rounded-md bg-base-700 max-h-44 sm:text-sm"
                                                >
                                                    <ListboxOption
                                                        v-for="hour in 24"
                                                        :key="hour"
                                                        v-slot="{ active, selected }"
                                                        as="template"
                                                        :value="hour - 1"
                                                    >
                                                        <li
                                                            :class="[
                                                                active ? 'bg-primary-500' : '',
                                                                'text-neutral relative cursor-default select-none py-2 pl-8 pr-4',
                                                            ]"
                                                        >
                                                            <span
                                                                :class="[
                                                                    selected ? 'font-semibold' : 'font-normal',
                                                                    'block truncate',
                                                                ]"
                                                            >
                                                                {{ hour }}
                                                            </span>

                                                            <span
                                                                v-if="selected"
                                                                :class="[
                                                                    active ? 'text-neutral' : 'text-primary-500',
                                                                    'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                                                ]"
                                                            >
                                                                <CheckIcon class="w-5 h-5" aria-hidden="true" />
                                                            </span>
                                                        </li>
                                                    </ListboxOption>
                                                </ListboxOptions>
                                            </transition>
                                        </div>
                                    </Listbox>
                                </dd>
                            </div>
                        </dl>
                        <dl class="sm:divide-y sm:divide-base-400">
                            <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                                <dt class="text-sm font-medium">Run Week Days</dt>
                                <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                    <label for="beginWeekDay"> Begin Week Day </label>
                                    <Listbox
                                        v-model="scrubbingSchedule.beginWeekDay"
                                        as="div"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                    >
                                        <div class="relative">
                                            <ListboxButton
                                                class="block pl-3 text-left w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                            >
                                                <span class="block truncate">
                                                    {{
                                                        weekDays.find(
                                                            (wd) =>
                                                                wd.value === parseInt(scrubbingSchedule?.beginWeekDay ?? '0'),
                                                        )?.name
                                                    }}
                                                </span>
                                                <span
                                                    class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none"
                                                >
                                                    <ChevronDownIcon class="w-5 h-5 text-gray-400" aria-hidden="true" />
                                                </span>
                                            </ListboxButton>

                                            <transition
                                                leave-active-class="transition duration-100 ease-in"
                                                leave-from-class="opacity-100"
                                                leave-to-class="opacity-0"
                                            >
                                                <ListboxOptions
                                                    class="absolute z-10 w-full py-1 mt-1 overflow-auto text-base rounded-md bg-base-700 max-h-44 sm:text-sm"
                                                >
                                                    <ListboxOption
                                                        v-for="day in weekDays"
                                                        :key="day.value"
                                                        v-slot="{ active, selected }"
                                                        as="template"
                                                        :value="day.value"
                                                    >
                                                        <li
                                                            :class="[
                                                                active ? 'bg-primary-500' : '',
                                                                'text-neutral relative cursor-default select-none py-2 pl-8 pr-4',
                                                            ]"
                                                        >
                                                            <span
                                                                :class="[
                                                                    selected ? 'font-semibold' : 'font-normal',
                                                                    'block truncate',
                                                                ]"
                                                            >
                                                                {{ day.name }}
                                                            </span>

                                                            <span
                                                                v-if="selected"
                                                                :class="[
                                                                    active ? 'text-neutral' : 'text-primary-500',
                                                                    'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                                                ]"
                                                            >
                                                                <CheckIcon class="w-5 h-5" aria-hidden="true" />
                                                            </span>
                                                        </li>
                                                    </ListboxOption>
                                                </ListboxOptions>
                                            </transition>
                                        </div>
                                    </Listbox>

                                    <label for="endWeekDay"> End Week Day </label>
                                    <Listbox
                                        v-model="scrubbingSchedule.endWeekDay"
                                        as="div"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                    >
                                        <div class="relative">
                                            <ListboxButton
                                                class="block pl-3 text-left w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                            >
                                                <span class="block truncate">
                                                    {{
                                                        weekDays.find(
                                                            (wd) => wd.value === parseInt(scrubbingSchedule?.endWeekDay ?? '0'),
                                                        )?.name
                                                    }}
                                                </span>
                                                <span
                                                    class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none"
                                                >
                                                    <ChevronDownIcon class="w-5 h-5 text-gray-400" aria-hidden="true" />
                                                </span>
                                            </ListboxButton>

                                            <transition
                                                leave-active-class="transition duration-100 ease-in"
                                                leave-from-class="opacity-100"
                                                leave-to-class="opacity-0"
                                            >
                                                <ListboxOptions
                                                    class="absolute z-10 w-full py-1 mt-1 overflow-auto text-base rounded-md bg-base-700 max-h-44 sm:text-sm"
                                                >
                                                    <ListboxOption
                                                        v-for="day in weekDays"
                                                        :key="day.value"
                                                        v-slot="{ active, selected }"
                                                        as="template"
                                                        :value="day.value"
                                                    >
                                                        <li
                                                            :class="[
                                                                active ? 'bg-primary-500' : '',
                                                                'text-neutral relative cursor-default select-none py-2 pl-8 pr-4',
                                                            ]"
                                                        >
                                                            <span
                                                                :class="[
                                                                    selected ? 'font-semibold' : 'font-normal',
                                                                    'block truncate',
                                                                ]"
                                                            >
                                                                {{ day.name }}
                                                            </span>

                                                            <span
                                                                v-if="selected"
                                                                :class="[
                                                                    active ? 'text-neutral' : 'text-primary-500',
                                                                    'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                                                ]"
                                                            >
                                                                <CheckIcon class="w-5 h-5" aria-hidden="true" />
                                                            </span>
                                                        </li>
                                                    </ListboxOption>
                                                </ListboxOptions>
                                            </transition>
                                        </div>
                                    </Listbox>
                                </dd>
                            </div>
                        </dl>
                        <dl class="sm:divide-y sm:divide-base-400">
                            <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                                <dt class="text-sm font-medium">Scrub Intervals</dt>
                                <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                    <!-- TODO -->
                                    <label for="minScrubInterval"> Min Scrub Interval </label>
                                    <VeeField
                                        v-model="scrubbingSchedule.minScrubInterval"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                        type="number"
                                        name="minScrubInterval"
                                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                    />

                                    <!-- TODO -->
                                    <label for="maxScrubInterval"> Max Scrub Interval </label>
                                    <VeeField
                                        v-model="scrubbingSchedule.maxScrubInterval"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                        type="number"
                                        name="maxScrubInterval"
                                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                    />

                                    <!-- TODO -->
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
                                    <!-- TODO -->
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
                        <dl class="sm:divide-y sm:divide-base-400">
                            <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                                <dt class="text-sm font-medium">Save Changes</dt>
                                <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                    <button
                                        type="submit"
                                        class="gap-x-1.5 rounded-md px-3 py-2 text-sm font-semibold text-neutral bg-success-500/10 text-success-400 ring-success-500/20"
                                    >
                                        Save Changes
                                    </button>
                                </dd>
                            </div>
                        </dl>
                    </form>
                </template>
            </div>
        </div>
    </div>
</template>
