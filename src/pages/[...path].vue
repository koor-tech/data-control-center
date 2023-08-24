<template>
  <div class="fullscreen bg-blue text-white text-center q-pa-md flex flex-center">
    <div>
      <div style="font-size: 30vh">404</div>

      <div class="text-h2" style="opacity: 0.4">Oops. Nothing here...</div>

      <q-btn class="q-mt-xl" color="white" text-color="blue" unelevated to="/" label="Go Home" no-caps />
    </div>
  </div>
  <div class="h-full justify-between flex flex-col">
    <HeroFull>
      <ContentCenterWrapper>
        <div style="font-size: 30vh">404</div>

        <div class="text-h2" style="opacity: 0.4">Oops. Nothing here...</div>
      </ContentCenterWrapper>
    </HeroFull>
    <Footer />
  </div>
</template>

<script lang="ts" setup>
import ContentCenterWrapper from 'components/partials/ContentCenterWrapper.vue';
import Footer from 'components/partials/Footer.vue';
import HeroFull from 'components/partials/HeroFull.vue';


</script>
