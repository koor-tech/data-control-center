<script lang="ts" setup>
import { ConnectError } from '@connectrpc/connect';
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';
import Footer from '~/components/partials/Footer.vue';
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

const { $grpc } = useNuxtApp();
const authStore = useAuthStore();

const { accessToken } = storeToRefs(authStore);
const { signOut } = authStore;

function redirect(): void {
    setTimeout(async () => {
        const route = useRoute();
        if (route.name === 'auth-logout') {
            await navigateTo({ name: 'index' });
        }
    }, 1500);
}

onBeforeMount(async () => {
    signOut();

    if (!accessToken.value) {
        redirect();
        return;
    }

    try {
        await $grpc.getAuthClient().logout({});
    } catch (e) {
        if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
    }

    redirect();
});
</script>

<template>
    <div class="h-full justify-between flex flex-col">
        <HeroFull>
            <ContentCenterWrapper class="max-w-2xl mx-auto text-center">
                <h2 class="text-4xl font-bold tracking-tight text-neutral sm:text-6xl">
                    Logout
                </h2>
                <p class="mt-6 text-lg leading-8 text-gray-300">
                    You are being logged out.
                </p>
            </ContentCenterWrapper>
        </HeroFull>
        <Footer />
    </div>
</template>
