<template>
  <div class="q-pa-md" style="max-width: 350px">
    <div class="text-h3 text-center">Licenses</div>
    <q-list bordered padding>
      <q-item-label header>Subscriptions</q-item-label>

      <q-item clickable v-ripple>
        <q-item-section>
          <q-item-label>Koor Pro</q-item-label>
          <q-item-label caption> Koor Data Control Center </q-item-label>
          <q-item-label> Renewal on 15 July 2024 </q-item-label>
        </q-item-section>
        <q-item-section side top>
          <q-chip color="primary">12 nodes</q-chip>
        </q-item-section>
      </q-item>

      <q-item clickable v-ripple>
        <q-item-section>
          <q-item-label>Koor Free Trial</q-item-label>
          <q-item-label caption> Limited Free Trial </q-item-label>
          <q-item-label> Registered on 1 July 2023 </q-item-label>
        </q-item-section>
        <q-item-section side top>
          <q-chip color="primary">4 nodes</q-chip>
        </q-item-section>
      </q-item>
    </q-list>

    <div class="q-mt-md text-h3 text-center">Settings</div>
    <q-list bordered padding>
      <q-item-label header>General</q-item-label>

      <q-item tag="label" v-ripple>
        <q-item-section side top>
          <q-checkbox v-model="emailConsent" />
        </q-item-section>
        <q-item-section>
          <q-item-label>Communication</q-item-label>
          <q-item-label caption> Keep me informed about important updates from Koor. </q-item-label>
        </q-item-section>
      </q-item>

      <q-item tag="label" v-ripple>
        <q-item-section side top>
          <q-checkbox v-model="cookiesConsent" />
        </q-item-section>

        <q-item-section>
          <q-item-label>Cookies</q-item-label>
          <q-item-label caption>
            I understand that cookies are used so that applications will function correctly.
          </q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const emailConsent = ref(true)
const cookiesConsent = ref(true)
</script>

<style lang="scss" scoped></style>
