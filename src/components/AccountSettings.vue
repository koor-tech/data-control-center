<template>
  <div class="q-pa-md" style="max-width: 350px">
    <div class="text-h3 text-center">Licenses</div>
    <ul bordered padding>
      <li header>Subscriptions</li>

      <q-item clickable v-ripple>
        <q-item-section>
          <li>Koor Pro</li>
          <li caption> Koor Data Control Center </li>
          <li> Renewal on 15 July 2024 </li>
        </q-item-section>
        <q-item-section side top>
          <q-chip color="primary">12 nodes</q-chip>
        </q-item-section>
      </q-item>

      <q-item clickable v-ripple>
        <q-item-section>
          <li>Koor Free Trial</li>
          <li caption> Limited Free Trial </li>
          <li> Registered on 1 July 2023 </li>
        </q-item-section>
        <q-item-section side top>
          <q-chip color="primary">4 nodes</q-chip>
        </q-item-section>
      </q-item>
    </ul>

    <div class="q-mt-md text-h3 text-center">Settings</div>
    <ul bordered padding>
      <li header>General</li>

      <q-item tag="label" v-ripple>
        <q-item-section side top>
          <q-checkbox v-model="emailConsent" />
        </q-item-section>
        <q-item-section>
          <li>Communication</li>
          <li caption> Keep me informed about important updates from Koor. </li>
        </q-item-section>
      </q-item>

      <q-item tag="label" v-ripple>
        <q-item-section side top>
          <q-checkbox v-model="cookiesConsent" />
        </q-item-section>

        <q-item-section>
          <li>Cookies</li>
          <li caption>
            I understand that cookies are used so that applications will function correctly.
          </li>
        </q-item-section>
      </q-item>
    </ul>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const emailConsent = ref(true)
const cookiesConsent = ref(true)
</script>
