<script setup lang="ts">
import { AccountGroupIcon, HarddiskIcon } from 'mdi-vue3';
import Cards from '~/components/partials/Cards.vue';
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';
import type { CardElements } from '~/utils/types';

useHead({
    title: 'Controls',
});
definePageMeta({
    title: 'Controls',
    requiresAuth: true,
});

const features = [
    {
        title: 'OSD Settings',
        description: 'Settings for OSDs (e.g., Scrubbing schedule)',
        href: { name: 'controls-osds' },
        icon: markRaw(HarddiskIcon),
        iconForeground: 'text-purple-900',
        iconBackground: 'bg-purple-100',
    },
    {
        title: 'Ceph Dashboard Users',
        description: 'Manage Ceph Dashboard users.',
        href: { name: 'controls-cephusers' },
        icon: markRaw(AccountGroupIcon),
        iconForeground: 'text-zinc-900',
        iconBackground: 'bg-zinc-100',
    },
] as CardElements;
</script>

<template>
    <ContentCenterWrapper>
        <Cards :items="features" />
    </ContentCenterWrapper>
</template>
