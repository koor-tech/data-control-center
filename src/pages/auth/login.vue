<script lang="ts" setup>
import Login from '~/components/auth/Login.vue';
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';
import PageFooter from '~/components/partials/PageFooter.vue';
import HeroFull from '~/components/partials/HeroFull.vue';

useHead({
    title: 'Login',
});
definePageMeta({
    title: 'Login',
    requiresAuth: false,
});
</script>

<template>
    <div class="h-full justify-between flex flex-col">
        <HeroFull>
            <ContentCenterWrapper>
                <Login />
            </ContentCenterWrapper>
        </HeroFull>
        <PageFooter />
    </div>
</template>
