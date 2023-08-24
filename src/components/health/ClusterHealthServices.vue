<template>
    <dl class="mx-auto grid grid-cols-1 gap-px bg-gray-900/6 sm:grid-cols-2 lg:grid-cols-6">
        <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
            <dt class="text-sm font-medium leading-6 text-gray-500">Crashes</dt>
            <dd :class="'text-rose-500'" v-for="crash in data.daemon_crashes" :key="crash.description"> {{ crash.description }}</dd>
            <!--                    <dd class="w-full flex-none text-1xl font-medium leading-6 tracking-tight text-gray-900">8383</dd>-->
        </div>
        <!--                <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">-->
        <!--                    <dt class="text-sm font-medium leading-6 text-or">Health</dt>-->
        <!--                    <dd :class="'text-xs font-medium'">WARN</dd>-->
        <!--                    <div class="radial-progress text-primary" style="&#45;&#45;value:70;">70%</div>-->
        <!--                </div>-->
        <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
            <dt class="text-sm font-medium leading-6 text-or">Mon</dt>
            <ul role="list" class="divide-y divide-gray-100">
                <!--                    <li v-for="person in people" :key="person.email" class="flex justify-between gap-x-6 py-5">-->
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">daemon_count: <span class="font-bold text-gray-900">{{ data.services.mon.daemon_count }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">quorum: <span class="font-bold text-gray-900">{{ data.services.mon.quorum }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">age: <span class="font-bold text-gray-900">{{ data.services.mon.age }}</span></p>
                </li>
            </ul>
        </div>
        <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
            <dt class="text-sm font-medium leading-6 text-or">Mgr</dt>
            <ul role="list" class="divide-y divide-gray-100">
                <!--                    <li v-for="person in people" :key="person.email" class="flex justify-between gap-x-6 py-5">-->
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">standbys: <span class="font-bold text-gray-900">{{ data.services.mgr.standbys }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">active: <span class="font-bold text-gray-900">{{ data.services.mgr.active }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">since: <span class="font-bold text-gray-900">{{ data.services.mgr.since }}</span></p>
                </li>
            </ul>
        </div>
        <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
            <dt class="text-sm font-medium leading-6 text-or">Mds</dt>
            <ul role="list" class="divide-y divide-gray-100">
                <!--                    <li v-for="person in people" :key="person.email" class="flex justify-between gap-x-6 py-5">-->
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">daemons_up: <span class="font-bold text-gray-900">{{ data.services.mds.daemons_up }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">hot_standby_count: <span class="font-bold text-gray-900">{{ data.services.mds.hot_standby_count }}</span></p>
                </li>
            </ul>
        </div>
        <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
            <dt class="text-sm font-medium leading-6 text-or"> - Osd -</dt>
            <ul role="list" class="divide-y divide-gray-100">
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">osd_count: <span class="font-bold text-gray-900">{{ data.services.osd.osd_count }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">osd_up: <span class="font-bold text-gray-900">{{ data.services.osd.osd_up }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">osd_in: <span class="font-bold text-gray-900">{{ data.services.osd.osd_in }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">since: <span class="font-bold text-gray-900">{{ data.services.osd.since }}</span></p>
                </li>
            </ul>
        </div>
        <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
            <dt class="text-sm font-medium leading-6 text-or">RGW</dt>
            <ul role="list" class="divide-y divide-gray-100">
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">active_daemon: <span class="font-bold text-gray-900">{{ data.services.rgw.active_daemon }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">host_count: <span class="font-bold text-gray-900">{{ data.services.rgw.host_count }}</span></p>
                </li>
                <li class="flex justify-between gap-x-4">
                    <p class="text-sm leading-6 text-gray-900">zone_count: <span class="font-bold text-gray-900">{{ data.services.rgw.zone_count }}</span></p>
                </li>
            </ul>
        </div>
    </dl>
</template>
<script setup lang="ts">
defineProps({
    data: Object,
})
</script>
