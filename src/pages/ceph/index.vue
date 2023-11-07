<script setup lang="ts">
import UsersWidget from '~/components/ceph/UsersWidget.vue';

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
        <ul role="list" class="grid grid-cols-1 gap-x-6 gap-y-8 lg:grid-cols-3 xl:gap-x-8">
            <UsersWidget />
        </ul>
    </div>
</template>
