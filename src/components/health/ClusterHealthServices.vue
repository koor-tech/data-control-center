<script setup lang="ts">
import { ClusterStats } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import Time from '../partials/elements/Time.vue';

defineProps<{
    stats: ClusterStats;
}>();
</script>

<template>
    <div class="bg-white rounded-lg shadow dark:bg-gray-800 px-2 py-2">
        <dl class="mx-auto grid grid-cols-1 gap-px bg-gray-900/6 sm:grid-cols-2 lg:grid-cols-6">
            <div
                v-if="stats.crashes"
                class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8"
            >
                <dt class="text-sm font-medium leading-6 text-or">Crashes</dt>
                <dd :class="'text-rose-500'" v-for="crash in stats.crashes" :key="crash.description">
                    {{ crash.description }}
                </dd>
            </div>
            <template v-if="stats.services">
                <div
                    class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8"
                    v-if="stats.services.mon"
                >
                    <dt class="text-sm font-medium leading-6 text-or">MON</dt>
                    <ul role="list" class="divide-y divide-gray-100">
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Daemon Count: <span class="font-bold text-gray-900">{{ stats.services.mon.daemonCount }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Quorum: <span class="font-bold text-gray-900">{{ stats.services.mon.quorum }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Since:
                                <span class="font-bold text-gray-900"><Time :value="stats.services.mon.updatedAt" /></span>
                            </p>
                        </li>
                    </ul>
                </div>

                <div
                    class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8"
                    v-if="stats.services.mgr"
                >
                    <dt class="text-sm font-medium leading-6 text-or">MGR</dt>
                    <ul role="list" class="divide-y divide-gray-100">
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Standbys: <span class="font-bold text-gray-900">{{ stats.services.mgr.standbys }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Active: <span class="font-bold text-gray-900">{{ stats.services.mgr.active }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Since:
                                <span class="font-bold text-gray-900"><Time :value="stats.services.mgr.updatedAt" /></span>
                            </p>
                        </li>
                    </ul>
                </div>

                <div
                    class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8"
                    v-if="stats.services.mds"
                >
                    <dt class="text-sm font-medium leading-6 text-or">MDS</dt>
                    <ul role="list" class="divide-y divide-gray-100">
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Daemons Up: <span class="font-bold text-gray-900">{{ stats.services.mds.daemonsUp }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Hot Standby Count:
                                <span class="font-bold text-gray-900">{{ stats.services.mds.hotStandbyCount }}</span>
                            </p>
                        </li>
                    </ul>
                </div>

                <div
                    class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8"
                    v-if="stats.services.osd"
                >
                    <dt class="text-sm font-medium leading-6 text-or">OSD</dt>
                    <ul role="list" class="divide-y divide-gray-100">
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                OSD Count: <span class="font-bold text-gray-900">{{ stats.services.osd.osdCount }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Up: <span class="font-bold text-gray-900">{{ stats.services.osd.osdUp }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                In: <span class="font-bold text-gray-900">{{ stats.services.osd.osdIn }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Up Since:
                                <span class="font-bold text-gray-900"><Time :value="stats.services.osd.osdUpUpdatedAt" /></span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                In since:
                                <span class="font-bold text-gray-900"><Time :value="stats.services.osd.osdInUpdatedAt" /></span>
                            </p>
                        </li>
                    </ul>
                </div>

                <div
                    class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8"
                    v-if="stats.services.rgw"
                >
                    <dt class="text-sm font-medium leading-6 text-or">RGW</dt>
                    <ul role="list" class="divide-y divide-gray-100">
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Active Daemons:
                                <span class="font-bold text-gray-900">{{ stats.services.rgw.activeDaemon }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Host Count: <span class="font-bold text-gray-900">{{ stats.services.rgw.hostCount }}</span>
                            </p>
                        </li>
                        <li class="flex justify-between gap-x-4">
                            <p class="text-sm leading-6 text-gray-900">
                                Zone Count: <span class="font-bold text-gray-900">{{ stats.services.rgw.zoneCount }}</span>
                            </p>
                        </li>
                    </ul>
                </div>
            </template>
        </dl>
    </div>
</template>
