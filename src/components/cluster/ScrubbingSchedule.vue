<script lang="ts" setup>
import { Switch, SwitchGroup, SwitchLabel } from '@headlessui/vue';
import { OSDScrubbingSchedule } from '~~/gen/ts/api/resources/koor/v1/koor_pb';

const osdScrubbingSchedule = ref<OSDScrubbingSchedule>(
    new OSDScrubbingSchedule({
        applySchedule: false,
    }),
);
</script>

<template>
    <div>
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
                            <SwitchGroup as="div" class="flex items-center">
                                <Switch
                                    v-model="osdScrubbingSchedule.applySchedule"
                                    :class="[
                                        osdScrubbingSchedule.applySchedule ? 'bg-indigo-600' : 'bg-gray-200',
                                        'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-600 focus:ring-offset-2',
                                    ]"
                                >
                                    <span
                                        aria-hidden="true"
                                        :class="[
                                            osdScrubbingSchedule.applySchedule ? 'translate-x-5' : 'translate-x-0',
                                            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
                                        ]"
                                    />
                                </Switch>
                                <SwitchLabel as="span" class="ml-3 text-sm">
                                    <span class="font-medium text-gray-300">Enabled</span>
                                </SwitchLabel>
                            </SwitchGroup>
                        </dd>
                    </div>
                </dl>
                <template v-if="osdScrubbingSchedule.applySchedule">
                    <dl class="sm:divide-y sm:divide-base-400">
                        <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                            <dt class="text-sm font-medium">Time Span</dt>
                            <dd class="mt-1 text-sm sm:col-span-2 sm:mt-0">
                                <!-- TODO add calendar for selecting each time span -->
                            </dd>
                        </div>
                    </dl>
                </template>
            </div>
        </div>
    </div>
</template>
