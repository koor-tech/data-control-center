<script setup lang="ts">
import { useClipboard } from '@vueuse/core';
import HeroFull from '~/components/partials/HeroFull.vue';
import LoadingBar from '~/components/partials/LoadingBar.vue';
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';

useHead({
    title: 'Error occured - Data Control Center',
});

const clipboard = useClipboard();

const props = defineProps<{
    error: Error | Object;
}>();

const buttonDisabled = ref(true);

function handleError(url?: string): void {
    startButtonTimer();
    if (url === undefined) url = '/';
    clearError({ redirect: url });
}

function copyError(): void {
    if (!props.error) {
        return;
    }

    clipboard.copy(JSON.stringify(props.error));
}

function startButtonTimer(): void {
    buttonDisabled.value = true;

    setTimeout(() => (buttonDisabled.value = false), 2000);
}

startButtonTimer();
</script>

<template>
    <div class="h-dscreen">
        <LoadingBar />
        <HeroFull>
            <ContentCenterWrapper class="max-w-3xl mx-auto text-center">
                <KoorLogo class="h-auto w-24 mx-auto mb-2" />

                <h1 class="text-5xl font-bold text-neutral">Error occured</h1>
                <h2 class="text-xl text-neutral">Critical Error Occured</h2>
                <div class="py-2 text-neutral mb-4">
                    <p class="py-2 font-semibold">Error Message:</p>
                    <span v-if="error">
                        <!-- @vue-expect-error -->
                        <pre
                            v-if="error.statusMessage"
                            v-text="
                                // @ts-ignore
                                error.statusMessage
                            "
                        />
                        <!-- @vue-expect-error -->
                        <pre
                            v-else-if="
                                // @ts-ignore
                                error.message
                            "
                            v-text="
                                // @ts-ignore
                                error.message
                            "
                        />
                        <pre v-else>Unable to get error message</pre>
                    </span>
                    <span v-else>
                        <pre>Unknown error</pre>
                    </span>
                </div>

                <div class="flex justify-center">
                    <button
                        :disabled="buttonDisabled"
                        class="rounded-md w-60 px-3.5 py-2.5 text-sm font-semibold text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                        :class="[
                            buttonDisabled
                                ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                                : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                        ]"
                        @click="handleError()"
                    >
                        Home
                    </button>

                    <button
                        :disabled="buttonDisabled"
                        class="rounded-md w-60 px-3.5 py-2.5 sm:ml-4 text-sm font-semibold text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                        :class="[
                            buttonDisabled
                                ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                                : 'bg-secondary-500 hover:bg-secondary-400 focus-visible:outline-secondary-500',
                        ]"
                        @click="handleError(useRoute().fullPath)"
                    >
                        Retry
                    </button>

                    <!-- @vue-expect-error -->
                    <button
                        v-if="error && (error.statusMessage || error.message)"
                        class="rounded-md w-60 bg-base-600 sm:ml-4 px-3.5 py-2.5 text-sm font-semibold text-neutral hover:bg-base-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-base-500"
                        @click="copyError"
                    >
                        Copy Error
                    </button>
                </div>
            </ContentCenterWrapper>
        </HeroFull>
    </div>
</template>
