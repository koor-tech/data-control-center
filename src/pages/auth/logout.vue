<script lang="ts" setup>
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';
import PageFooter from '~/components/partials/PageFooter.vue';
import HeroFull from '~/components/partials/HeroFull.vue';
import { useAuthStore } from '~/store/auth';

useHead({
    title: 'Logout',
});
definePageMeta({
    title: 'Logout',
    requiresAuth: true,
    authOnlyToken: true,
    showCookieOptions: true,
});

const authStore = useAuthStore();
const { doLogout } = authStore;

function redirect(): void {
    setTimeout(async () => {
        const route = useRoute();
        if (route.name === 'auth-logout') {
            await navigateTo({ name: 'index' });
        }
    }, 1500);
}

onMounted(async () => {
    try {
        await doLogout();
    } finally {
        redirect();
    }
});
</script>

<template>
    <div class="h-full justify-between flex flex-col">
        <HeroFull>
            <ContentCenterWrapper class="max-w-2xl mx-auto text-center">
                <h2 class="text-4xl font-bold tracking-tight text-neutral sm:text-6xl">Logout</h2>
                <p class="mt-6 text-lg leading-8 text-gray-300">You are being logged out.</p>
            </ContentCenterWrapper>
        </HeroFull>
        <PageFooter />
    </div>
</template>
