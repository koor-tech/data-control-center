<template>
  <q-page padding>
    <div v-if="authStore.accessToken">
      <div class="flex">
        <KoorAccount v-if="authStore.accessToken" :accessToken="authStore.accessToken" />
        <AccountSettings />
      </div>
      <hr />
      <div class="q-mt-lg">
        <q-btn color="warning" @click="signOut">Sign Out</q-btn>
      </div>
    </div>
    <div v-else>
      <AccountSignIn />
    </div>
  </q-page>
</template>

<script setup lang="ts">
import KoorAccount from '../components/KoorAccount.vue'
import AccountSettings from '../components/AccountSettings.vue'
import AccountSignIn from '../components/AccountSignIn.vue'
import { useAuthStore } from 'src/stores/auth';

const authStore = useAuthStore();

async function signOut() {
  authStore.signOut();
}
</script>
