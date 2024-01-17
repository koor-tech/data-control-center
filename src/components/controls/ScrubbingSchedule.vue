<script lang="ts" setup>
import type { ConnectError } from '@connectrpc/connect';
import { Listbox, ListboxButton, ListboxOption, ListboxOptions } from '@headlessui/vue';
import { CheckIcon, ChevronDownIcon } from 'mdi-vue3';
import { vMaska } from 'maska';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';
import { useNotificationsStore } from '~/store/notifications';

const { $grpc } = useNuxtApp();

const appConfig = useAppConfig();

const notifications = useNotificationsStore();

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

        notifications.dispatchNotification({
            title: 'Scrubbing Schedule saved successfully!',
            content: 'Your scrubbing schedule has been saved successfully.',
            type: 'success',
        });

        return response.osdScrubbingSchedule!;
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}

const weekDays = [
    { name: 'Sunday', value: 0 },
    { name: 'Monday', value: 1 },
    { name: 'Tuesday', value: 2 },
    { name: 'Wednesday', value: 3 },
    { name: 'Thursday', value: 4 },
    { name: 'Friday', value: 5 },
    { name: 'Saturday', value: 6 },
];

const canSubmit = ref(true);
</script>

<template>
    <div class="mx-auto max-w-7xl">
        <div class="overflow-hidden bg-base-800 shadow sm:rounded-lg text-neutral">
            <div class="px-4 py-5 sm:px-6">
                <h3 class="text-base font-semibold leading-6">Custom OSD Scrubbing Schedule</h3>
                <p class="mt-1 max-w-2xl text-sm">
                    Set a custom OSD scrubbing schedule here. For more information on scrubbing schedules checkout
                    <NuxtLink
                        :external="true"
                        to="https://docs.ceph.com/en/latest/rados/configuration/osd-config-ref/#scrubbing"
                        target="_blank"
                        class="underline"
                    >
                        Ceph's documentation here</NuxtLink
                    >.
                </p>
            </div>
            <div class="border-t border-b border-base-400 px-4 py-5 sm:p-0">
                <DataErrorBlock v-if="error" :retry="refresh" :message="error.message" />
                <DataPendingBlock v-else-if="pending" message="Loading Scrubbing Schedule" />
                <DataNoDataBlock v-else-if="scrubbingSchedule === null || scrubbingSchedule === undefined" />

                <template v-else>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-semibold">Maximumg Scrubbing Operations</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <VeeField
                                    v-model="scrubbingSchedule.maxScrubOps"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="maxScrubOps"
                                    class="block w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                    min="0"
                                    max="9999999"
                                />

                                <p>The maximum number of simultaneous scrub operations for a Ceph OSD Daemon.</p>
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-semibold">Run Hours</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <label for="beginHour" class="font-semibold"> Begin Hour </label>
                                <Listbox
                                    v-model="scrubbingSchedule.beginHour"
                                    as="div"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                >
                                    <div class="relative">
                                        <ListboxButton
                                            class="block pl-3 text-left w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                        >
                                            <span class="block truncate"> {{ scrubbingSchedule.beginHour }}:00 </span>
                                            <span class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
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
                                                    v-for="(_, hour) in 24"
                                                    :key="hour"
                                                    v-slot="{ active, selected }"
                                                    as="template"
                                                    :value="hour"
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
                                                            {{ hour }}:00
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

                                <label for="endHour" class="font-semibold"> End Hour </label>
                                <Listbox
                                    v-model="scrubbingSchedule.endHour"
                                    as="div"
                                    :disabled="!canSubmit || appConfig.readOnly"
                                >
                                    <div class="relative">
                                        <ListboxButton
                                            class="block pl-3 text-left w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                        >
                                            <span class="block truncate"> {{ scrubbingSchedule.endHour }}:00 </span>
                                            <span class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
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
                                                    v-for="(_, hour) in 24"
                                                    :key="hour"
                                                    v-slot="{ active, selected }"
                                                    as="template"
                                                    :value="hour"
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
                                                            {{ hour }}:00
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

                                <p>
                                    This restricts scrubbing to this hour of the day or later. Together these options define the
                                    time frame, in which the scrubs can happen. But a scrub will be performed no matter whether
                                    the time window allows or not, as long as the placement group's scrub interval exceeds the
                                    OSD Max scrub interval below.
                                </p>
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-semibold">Run Week Days</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <label for="beginWeekDay" class="font-semibold"> Begin Week Day </label>
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
                                                        (wd) => wd.value === parseInt(scrubbingSchedule?.beginWeekDay ?? '0'),
                                                    )?.name
                                                }}
                                            </span>
                                            <span class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
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

                                <label for="endWeekDay" class="font-semibold"> End Week Day </label>
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
                                            <span class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
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

                                <p>Setting both to "Sunday" will allow scrubbing the whole week.</p>
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-semibold">Scrub Intervals (in Seconds)</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <label for="minScrubInterval" class="font-semibold"> Min Scrub Interval </label>
                                <div class="flex items-center gap-2">
                                    <VeeField
                                        v-model="scrubbingSchedule.minScrubInterval"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                        type="number"
                                        name="minScrubInterval"
                                        class="block w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                    />
                                    <span class="shrink-0 min-w-60 max-w-60">
                                        {{
                                            fromSecondsToFormattedDuration(parseInt(scrubbingSchedule.minScrubInterval ?? '0'))
                                        }}
                                    </span>
                                </div>
                                <p>Minimum interval between scrubbing operations during low cluster loads.</p>

                                <label for="maxScrubInterval" class="font-semibold"> Max Scrub Interval </label>
                                <div class="flex items-center gap-2">
                                    <VeeField
                                        v-model="scrubbingSchedule.maxScrubInterval"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                        type="number"
                                        name="maxScrubInterval"
                                        class="block w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                    />
                                    <span class="shrink-0 min-w-60 max-w-60">
                                        {{
                                            fromSecondsToFormattedDuration(parseInt(scrubbingSchedule.maxScrubInterval ?? '0'))
                                        }}
                                    </span>
                                </div>
                                <p>Maximum interval to wait for scrubbing operations no matter the cluster load.</p>

                                <label for="deepScrubInterval" class="font-semibold"> Deep Scrub Interval </label>
                                <div class="flex items-center gap-2">
                                    <VeeField
                                        v-model="scrubbingSchedule.deepScrubInterval"
                                        :disabled="!canSubmit || appConfig.readOnly"
                                        type="number"
                                        name="deepScrubInterval"
                                        class="block w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                    />
                                    <span class="shrink-0 min-w-60 max-w-60">
                                        {{
                                            fromSecondsToFormattedDuration(parseInt(scrubbingSchedule.deepScrubInterval ?? '0'))
                                        }}
                                    </span>
                                </div>
                                <p>
                                    The interval for "deep" scrubbing (fully reading all data). This ignores any scrubbing
                                    cluster load threshold settings.
                                </p>
                            </dd>
                        </div>
                    </dl>
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-semibold">Scrub Sleep (in Seconds)</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <VeeField
                                    v-model="scrubbingSchedule.scrubSleepSeconds"
                                    v-maska
                                    :disabled="!canSubmit || appConfig.readOnly"
                                    type="number"
                                    name="scrubSleepSeconds"
                                    class="block w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                    step=".1"
                                    data-maska="0.99"
                                    data-maska-tokens="0:\d:multiple|9:\d:optional"
                                />
                                <p>
                                    Sleep between scrubbing the next chunks. A general point of recommendation here is to not
                                    make it too big, but keep it sensible, e.g., for a medium sized cluster a sleep of
                                    <code>0.1</code>, can help reduce the impact of scrubbing on clients.
                                </p>
                            </dd>
                        </div>
                    </dl>
                </template>
            </div>
            <div class="border-t border-b border-base-400 px-4 py-5 sm:p-0">
                <dl class="sm:divide-y sm:divide-base-400">
                    <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                        <dt class="text-sm font-semibold"></dt>
                        <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                            <button
                                type="button"
                                class="gap-x-1.5 rounded-md px-3 py-2 text-sm font-semibold text-neutral"
                                :disabled="!canSubmit"
                                :class="[
                                    !canSubmit
                                        ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                                        : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                                ]"
                                @click="setScrubbingSchedule()"
                            >
                                Save Changes
                            </button>
                        </dd>
                    </div>
                </dl>
            </div>
        </div>
    </div>
</template>
