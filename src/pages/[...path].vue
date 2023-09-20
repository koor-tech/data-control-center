<script lang="ts" setup>
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';
import Footer from '~/components/partials/Footer.vue';
import HeroFull from '~/components/partials/HeroFull.vue';
</script>

<template>
    <div class="h-full justify-between flex flex-col">
        <HeroFull>
            <ContentCenterWrapper>
                <div class="text-white text-9xl">404</div>

                <div class="text-base-100 text-2xl">Oops. Nothing here...</div>
                <NuxtLink
                    class="bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500 rounded-md px-3.5 py-2.5 text-sm font-semibold text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                    :to="{ name: 'index' }"
                >
                    Go Home
                </NuxtLink>
            </ContentCenterWrapper>
        </HeroFull>
        <Footer />
    </div>
</template>
