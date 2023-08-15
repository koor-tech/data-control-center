<template>
  <div class="q-pa-md" style="width: 350px">
    <div class="text-h3 text-center">Profile</div>
    <q-form class="q-gutter-sm" @submit.prevent="updateProfile">
      <q-input id="email" label="Email" outlined v-model="session.user.email" disable />
      <q-input id="username" label="Alias" outlined v-model="username" />
      <q-input id="full_name" label="Full Name" outlined v-model="full_name" />
      <q-input id="website" label="Website" outlined type="url" v-model="website" />
      <div>
        <q-btn type="submit" color="primary" :label="loading ? 'Loading' : 'Update'" :disabled="loading" />
      </div>
    </q-form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

defineProps<{
  accessToken: string | null;
}>();

const loading = ref(true)
const username = ref('')
const full_name = ref('')
const website = ref('')
</script>

<style lang="scss" scoped></style>
