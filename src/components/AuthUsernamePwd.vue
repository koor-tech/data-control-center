<template>
  <div>
    <div class="text-h4 q-mb-md">Use your username and password to sign in</div>
    <q-form>
      <div>
        <q-input name="username" outlined type="username" label="username" placeholder="Your username"
          v-model="username" />
      </div>
      <div>
        <q-input name="password" outlined type="password" label="Password" placeholder="Your password"
          v-model="password" />
      </div>
      <div>
        <q-btn class="q-mt-sm" color="primary" text-color="dark" label="Sign In" :disabled="loading" />
      </div>
    </q-form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const loading = ref(false);
const username = ref('');
const password = ref('');
</script>

<style lang="scss" scoped></style>
