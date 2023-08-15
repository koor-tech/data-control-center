<template>
  <div>
    <q-toggle v-model="createAccount" color="primary" label="I need an account, please." />
    <div>
      <div class="q-pa-lg">
        <div class="text-bold">How would you like to identify yourself?</div>
        <q-option-group name="authMethod" v-model="methodForAuth" :options="options" color="primary" inline />
      </div>
      <AuthUsernamePwd v-if="methodForAuth === 'pwd'" />
      <div v-if="methodForAuth === 'oauth'" class="text-h4">Coming soon...</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import AuthUsernamePwd from './AuthUsernamePwd.vue'

const options = [
  {
    label: 'Email/Password',
    value: 'pwd',
  },
  {
    label: 'OAuth',
    value: 'oauth',
  },
]
const createAccount = ref(false)
const methodForAuth = ref(options[0].value)
</script>

<style lang="scss" scoped></style>
