<script lang="ts" setup>
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';
import HeroFull from '~/components/partials/HeroFull.vue';
import Login from '~/components/auth/Login.vue';
import Footer from '~/components/partials/Footer.vue';

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
        <Footer />
    </div>
</template>
