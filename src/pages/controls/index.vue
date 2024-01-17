<script setup lang="ts">
import { AccountGroupIcon, ApplicationEditIcon, HarddiskIcon } from 'mdi-vue3';
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
        description: 'Configuring OSD related settings such as, e.g., scrubbing schedule.',
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
    {
        title: 'Resource Editor',
        description: 'Edit Ceph custom resources via YAML editor.',
        href: { name: 'controls-editor-editor' },
        icon: markRaw(ApplicationEditIcon),
        iconForeground: 'text-amber-900',
        iconBackground: 'bg-amber-100',
    },
] as CardElements;
</script>

<template>
    <ContentCenterWrapper>
        <Cards :items="features" />
    </ContentCenterWrapper>
</template>
