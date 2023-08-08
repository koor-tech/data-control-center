<template>
  <div class="signInForm">
    <div class="text-h4">Use a Magic Link</div>
    <p>
      This may be the easiest way to verify who you are. Tell us your email address, and we will
      send you a magic link. Click on that link to prove that it's you, and everything you can do
      will be available.
    </p>
    <p class="text-bold text-italic">It's like magic.</p>
    <q-form @submit.prevent="handleLogin">
      <div>
        <q-input
          name="email"
          outlined
          type="email"
          label="Email"
          placeholder="Where to send magic link"
          v-model="email"
        />
      </div>
      <div>
        <q-btn
          @click="handleLogin"
          class="q-mt-sm"
          color="primary"
          text-color="dark"
          :label="loading ? 'Loading' : 'Send magic link'"
          :disabled="loading"
        />
      </div>
    </q-form>
    <q-banner
      v-if="showMessage"
      :class="success ? 'bg-info text-white' : 'bg-warning text-dark'"
      class="q-mt-sm"
    >
      {{ message }}
      <template v-slot:action>
        <q-btn @click="showMessage = false" flat color="white" label="Dismiss" />
      </template>
    </q-banner>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { signInMagicLinkStyle } from '../apis/supabase'

const loading = ref(false)
const email = ref('')
const showMessage = ref(false)
const message = ref('')
const success = ref(true)

const handleLogin = async () => {
  try {
    if (email.value === '') {
      notify('A valid email address is required')
      return
    }
    loading.value = true // TODO: think about using a spinner
    const { error } = await signInMagicLinkStyle(email.value)
    if (error) throw error
    notifySuccess()
  } catch (error) {
    if (error instanceof Error) {
      notifyFailure()
    }
  } finally {
    loading.value = false
  }
}

const notify = (msg, ok = true) => {
  message.value = msg
  success.value = ok
  showMessage.value = true
}

const notifySuccess = () => {
  notify('Check your email for the login link')
}

const notifyFailure = () => {
  notify('Something went wrong. Please contact support@koor.tech for assistance.', false)
}

/*
TODO: process this
redirected with:
http://localhost:9000/#access_token=eyJhbGciOiJIUzI1NiIsImtpZCI6ImdPWDFnYXRDclovWEl0QjQiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNjkwOTUwOTUxLCJpYXQiOjE2OTA5NDczNTEsImlzcyI6Imh0dHBzOi8vaG9iY3Vsenp2ZWNra2tyZXVneXIuc3VwYWJhc2UuY28vYXV0aC92MSIsInN1YiI6ImI1NTBmMzI1LWQ5ZWUtNDIzNS04NjM4LTJjM2Y0ODIwZTUyYSIsImVtYWlsIjoiZGF2ZUBrb29yLnRlY2giLCJwaG9uZSI6IiIsImFwcF9tZXRhZGF0YSI6eyJwcm92aWRlciI6ImVtYWlsIiwicHJvdmlkZXJzIjpbImVtYWlsIl19LCJ1c2VyX21ldGFkYXRhIjp7fSwicm9sZSI6ImF1dGhlbnRpY2F0ZWQiLCJhYWwiOiJhYWwxIiwiYW1yIjpbeyJtZXRob2QiOiJvdHAiLCJ0aW1lc3RhbXAiOjE2OTA5NDczNTF9XSwic2Vzc2lvbl9pZCI6ImYxMTI3NWYwLTE5MWQtNGJkMS1hZWRjLTA3ZTM2NThjOGRlNyJ9.6Yi7FBhSozalQOLQaF8wsrFdmZo6htfYKQRi0tAvxXE&expires_in=3600&refresh_token=Fjh6krZXZ9pj-rMeGT-_Ng&token_type=bearer&type=magiclink
*/
</script>

<style lang="scss" scoped>
.signInForm {
  width: 600px;
  padding: 2em;
}
</style>
