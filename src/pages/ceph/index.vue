<script setup lang="ts">
useHead({
    title: 'Ceph settings',
});
definePageMeta({
    title: 'Ceph settings',
    requiresAuth: true,
});
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Ceph Settings</h2>
        <CephSettings />
    </div>
</template>
