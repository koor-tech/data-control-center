<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';
import { Disclosure, DisclosureButton, DisclosurePanel } from '@headlessui/vue';
import { ChevronDownIcon } from 'mdi-vue3';
import ClusterRadar from '~/components/cluster/ClusterRadar.vue';
import NodeSummaryList from '~/components/cluster/nodes/NodeSummaryList.vue';
import StatusDials from '~/components/cluster/StatusDials.vue';
import ClusterHealthBar from '~/components/cluster/ClusterHealthBar.vue';
import Container from '~/components/partials/Container.vue';
import ClusterServices from '~/components/cluster/ClusterServices.vue';
import ResourceInfoList from '~/components/cluster/ResourceInfoList.vue';

useHead({
    title: 'Overview',
});
definePageMeta({
    title: 'Overview',
    requiresAuth: true,
});

const { $grpc } = useNuxtApp();

const { data: stats } = useLazyAsyncData('clusterStats', async () => {
    try {
        const response = await $grpc.getStatsClient().getClusterStats({});

        return response.stats!;
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});

const { data: radar, error: radarErr } = useLazyAsyncData(`clusterRadar`, () => $grpc.getStatsClient().getClusterRadar({}));
watch(radarErr, () => {
    if (radarErr.value !== null) $grpc.handleError(radarErr.value as ConnectError);
});

const { data: nodes, error: nodesErr } = useLazyAsyncData(`clusterNodes`, () => $grpc.getStatsClient().getClusterNodes({}));
watch(nodesErr, () => {
    if (nodesErr.value !== null) $grpc.handleError(nodesErr.value as ConnectError);
});

const { data: resources } = useLazyAsyncData(`clusterResources`, async () => {
    try {
        return await $grpc.getStatsClient().getClusterResources({});
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});
</script>

<template>
    <div class="p-2">
        <div class="flex flex-col gap-2">
            <div v-if="radar" class="grid grid-cols-3 gap-2">
                <ClusterRadar v-if="radar.radar" :radar="radar.radar" />
                <StatusDials v-if="radar.radar" class="col-span-2" :radar="radar.radar" />
            </div>
            <div v-if="stats" class="grid-cols-1">
                <ClusterHealthBar :health="stats.status" :cluster-id="stats.id" />
                <ClusterServices :stats="stats" />
            </div>
            <div class="grid-cols-2">
                <div v-if="nodes" class="grid grid-cols-1">
                    <dl class="space-y-2">
                        <Container>
                            <Disclosure v-slot="{ open }" as="div">
                                <dt>
                                    <DisclosureButton
                                        class="flex w-full items-start justify-between text-left text-base border-b border-gray-300"
                                    >
                                        <span class="text-lg font-semibold leading-7"> Storage Nodes </span>
                                        <span class="ml-6 flex h-7 items-center">
                                            <ChevronDownIcon
                                                :class="[open ? 'upsidedown' : '', 'h-6 w-6 transition-transform']"
                                                aria-hidden="true"
                                            />
                                        </span>
                                    </DisclosureButton>
                                </dt>
                                <DisclosurePanel as="dd" class="mt-2 pr-12">
                                    <NodeSummaryList v-if="nodes" :nodes="nodes.nodes" />
                                </DisclosurePanel>
                            </Disclosure>
                        </Container>
                        <template v-if="resources">
                            <Container>
                                <Disclosure v-slot="{ open }" as="div">
                                    <dt>
                                        <DisclosureButton
                                            class="flex w-full items-start justify-between text-left text-base border-b border-gray-300"
                                        >
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
                                <Disclosure v-slot="{ open }" as="div">
                                    <dt>
                                        <DisclosureButton
                                            class="flex w-full items-start justify-between text-left text-base border-b border-gray-300"
                                        >
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
                        </template>
                    </dl>
                </div>
            </div>
        </div>
    </div>
</template>
