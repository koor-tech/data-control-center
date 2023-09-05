<template>
  <div class="px-4">
    <div class="flex">
      <KoorAccount :accessToken="accessToken" />
      <AccountSettings />
    </div>
    <hr />
    <div class="q-mt-lg">
      <q-btn color="warning" @click="signOut">Sign Out</q-btn>
    </div>
  </div>
</template>

<script setup lang="ts">
import AccountSettings from '../components/AccountSettings.vue';
import { useAuthStore } from '~/store/auth';
import KoorAccount from '../components/KoorAccount.vue';

const authStore = useAuthStore();
const { accessToken } = storeToRefs(authStore);

async function signOut() {
  authStore.signOut();
}
</script>
