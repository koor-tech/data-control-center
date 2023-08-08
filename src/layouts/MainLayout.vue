<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

        <koor-logo />
        <q-toolbar-title> Koor Data Control Center </q-toolbar-title>

        <div>KDCC v0.1.0</div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label header> Access Panel </q-item-label>
        <InternalLink v-for="link in panelViewLinks" :key="link.title" v-bind="link" />

        <q-item-label header> Outside Resources </q-item-label>
        <ExternalLink v-for="link in externalLinks" :key="link.title" v-bind="link" />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { defineComponent, ref } from 'vue'
import ExternalLink from 'components/ExternalLink.vue'
import InternalLink from 'components/InternalLink.vue'
import KoorLogo from 'components/KoorLogo.vue'

const panelViewsList = [
  {
    title: 'Dashboard',
    caption: 'Get a quick look',
    icon: 'visibility',
    to: 'dashboard',
    exact: true,
  },
  {
    title: 'Charts',
    caption: 'Examine system metrics',
    icon: 'dashboard',
    to: 'charts',
  },
  {
    title: 'Controls',
    caption: 'Make system adjustments',
    icon: 'tune',
    to: 'controls',
  },
  {
    title: 'User Account',
    caption: 'All about you',
    icon: 'tune',
    to: 'account',
  },
]
const linksList = [
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

export default defineComponent({
  name: 'MainLayout',

  components: {
    ExternalLink,
    InternalLink,
    KoorLogo,
  },

  setup() {
    const leftDrawerOpen = ref(false)

    return {
      panelViewLinks: panelViewsList,
      externalLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value
      },
    }
  },
})
</script>
