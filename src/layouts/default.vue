<template>
  <div view="lHh Lpr lFf">
    <header elevated>
      <div>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

        <KoorLogo />
        <h2> Koor Data Control Center </h2>

        <div>KDCC v0.1.0</div>
      </div>
    </header>

    <div show-if-above bordered>
      <ul>
        <li header> Access Panel </li>
        <NuxtLink v-for="link in panelViewsList" :key="link.title" :to="link.to" />

        <li header> Outside Resources </li>
        <NuxtLink v-for="link in externalLinks" :key="link.title" :to="link.link" :external="true" />
      </ul>
    </div>

    <div>
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import KoorLogo from '~/components/KoorLogo.vue';
import { RoutesNamedLocations } from '../../.nuxt/typed-router/__routes';

const panelViewsList: {
  title: string;
  caption: string;
  icon: string;
  to: RoutesNamedLocations;
  exact?: boolean;
}[] = [
    {
      title: 'Dashboard',
      caption: 'Get a quick look',
      icon: 'visibility',
      to: { name: 'index' },
      exact: true,
    },
    {
      title: 'Charts',
      caption: 'Examine system metrics',
      icon: 'dashboard',
      to: { name: 'charts' },
    },
    {
      title: 'Controls',
      caption: 'Make system adjustments',
      icon: 'tune',
      to: { name: 'controls' },
    },
    {
      title: 'User Account',
      caption: 'All about you',
      icon: 'tune',
      to: { name: 'accounts' },
    },
  ];
const externalLinks = [
  {
    title: 'Knowledge Center',
    caption: 'Information about data storage',
    icon: 'school',
    link: 'https://kb.koor.tech',
  },
  {
    title: 'Github',
    caption: 'github.com/koor-tech',
    icon: 'code',
    link: 'https://github.com/koor-tech',
  },
  {
    title: 'X (Twitter)',
    caption: 'Follow @koor_tech',
    icon: 'rss_feed',
    link: 'https://twitter.com/koor_tech',
  },
  {
    title: 'YouTube',
    caption: 'Videos by @koor-tech',
    icon: 'public',
    link: 'https://www.youtube.com/@koor-tech',
  },
  {
    title: 'Quasar docs',
    caption: 'Powered by Quasar on VueJS',
    icon: 'public',
    link: 'https://quasar.dev/docs',
  },
]

const leftDrawerOpen = ref(false);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>
