<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';
import { Disclosure, DisclosureButton, DisclosurePanel } from '@headlessui/vue';
import { ChevronDownIcon } from 'mdi-vue3';
import ClusterRadar from '~/components/ClusterRadar.vue';
import StatusDials from '~/components/StatusDials.vue';
import ClusterHealthServices from '~/components/health/ClusterHealthServices.vue';
import Container from '~/components/partials/Container.vue';
//import NodeSummaryList from '~/components/NodeSummaryList.vue';

useHead({
    title: 'Overview',
});
definePageMeta({
    title: 'Overview',
    requiresAuth: true,
});

const { $grpc } = useNuxtApp();

const {
    data: stats,
    pending,
    refresh,
    error: statsErr,
} = useLazyAsyncData(`clusterstats`, () => $grpc.getStatsClient().getClusterStats({}));
watch(statsErr, () => {
    if (statsErr.value !== null) $grpc.handleError(statsErr.value as ConnectError);
});

const { data: radar, error: radarErr } = useLazyAsyncData(`clusterradar`, () => $grpc.getStatsClient().getClusterRadar({}));
watch(radarErr, () => {
    if (radarErr.value !== null) $grpc.handleError(radarErr.value as ConnectError);
});

const { data: resources } = useLazyAsyncData(`clusterresources`, async () => {
    try {
        return await $grpc.getStatsClient().getClusterResources({});
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});
</script>

<template>
    <div class="flex flex-col my-2 mx-2 gap-2">
        <div class="grid grid-cols-3 gap-2">
            <ClusterRadar v-if="radar && radar.radar" :radar="radar.radar" />
            <StatusDials v-if="radar && radar.radar" class="col-span-2" :radar="radar.radar" />
        </div>
        <div class="grid-cols-1">
            <ClusterHealthServices v-if="stats && stats.stats" :stats="stats.stats" />
        </div>
        <div class="grid-cols-2">
            <div class="col-4">
                <!-- <NodeSummaryList />-->
            </div>
            <div class="grid grid-cols-1" v-if="resources">
                <dl class="space-y-2">
                    <Container>
                        <Disclosure as="div" v-slot="{ open }">
                            <dt>
                                <DisclosureButton class="flex w-full items-start justify-between text-left text-base">
                                    <span class="text-lg font-semibold leading-7"> Deployment Resources </span>
                                    <span class="ml-6 flex h-7 items-center">
                                        <ChevronDownIcon
                                            :class="[open ? 'upsidedown' : '', 'h-6 w-6 transition-transform']"
                                            aria-hidden="true"
                                        />
                                    </span>
                                </DisclosureButton>
                            </dt>
                            <DisclosurePanel as="dd" class="mt-2 pr-12">
                                <div class="sm:flex sm:items-center">
                                    <div class="sm:flex-auto">
                                        <h3 class="text-base font-semibold leading-6 text-gray-900">
                                            Cluster Component Deployments
                                        </h3>
                                        <p class="mt-2 text-sm text-gray-700">
                                            A list of the Storage Cluster relevant resources.
                                        </p>
                                    </div>
                                </div>
                                <ResourceInfoList :rows="resources.deployments" />
                            </DisclosurePanel>
                        </Disclosure>
                    </Container>
                    <Container>
                        <Disclosure as="div" v-slot="{ open }">
                            <dt>
                                <DisclosureButton class="flex w-full items-start justify-between text-left text-base">
                                    <span class="text-lg font-semibold leading-7"> Ceph Resources </span>
                                    <span class="ml-6 flex h-7 items-center">
                                        <ChevronDownIcon
                                            :class="[open ? 'upsidedown' : '', 'h-6 w-6 transition-transform']"
                                            aria-hidden="true"
                                        />
                                    </span>
                                </DisclosureButton>
                            </dt>
                            <DisclosurePanel as="dd" class="mt-2 pr-12">
                                <div class="sm:flex sm:items-center">
                                    <div class="sm:flex-auto">
                                        <h2 class="text-base font-semibold leading-6 text-gray-900">Resource Info</h2>
                                        <p class="mt-2 text-sm text-gray-700">
                                            A list of the Storage Cluster relevant resources.
                                        </p>
                                    </div>
                                </div>
                                <ResourceInfoList :rows="resources.resources" />
                            </DisclosurePanel>
                        </Disclosure>
                    </Container>
                </dl>
            </div>
        </div>
    </div>
</template>
