<template>
  <div class="px-4">
    <div v-if="accessToken">
      <div class="flex">
        <KoorAccount v-if="accessToken" :accessToken="accessToken" />
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
  </div>
</template>

<script setup lang="ts">
import KoorAccount from '../components/KoorAccount.vue';
import AccountSettings from '../components/AccountSettings.vue';
import AccountSignIn from '../components/AccountSignIn.vue';
import { useAuthStore } from '~/store/auth';

const authStore = useAuthStore();
const { accessToken } = storeToRefs(authStore);

async function signOut() {
  authStore.signOut();
}
</script>
