<script lang="ts" setup>
import { Timestamp } from '@bufbuild/protobuf';

const props = withDefaults(
    defineProps<{
        value: Date | Timestamp | undefined;
        ago?: boolean;
    }>(),
    {
        type: 'short',
        ago: true,
    },
);

const date: Date = props.value instanceof Date ? props.value : toDate(props.value) ?? new Date();
</script>

<template>
    <time :datetime="date.toLocaleTimeString()" :title="!ago ? date.toLocaleString() : useTimeAgo(date).value">
        {{ !ago ? date.toLocaleString() : useTimeAgo(date).value }}
    </time>
</template>
