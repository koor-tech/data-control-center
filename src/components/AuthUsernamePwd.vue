<template>
  <div>
    <div class="text-h4 q-mb-md">Use your email and password to sign in</div>
    <q-form>
      <div>
        <q-input
          name="email"
          outlined
          type="email"
          label="Email"
          placeholder="Your email"
          v-model="email"
        />
      </div>
      <div>
        <q-input
          name="password"
          outlined
          type="password"
          label="Password"
          placeholder="Your password"
          v-model="password"
        />
      </div>
      <div>
        <q-btn
          class="q-mt-sm"
          color="primary"
          text-color="dark"
          label="Sign In"
          :disabled="loading"
        />
      </div>
    </q-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const loading = ref(false)
const email = ref('')
const password = ref('')
</script>

<style lang="scss" scoped></style>
