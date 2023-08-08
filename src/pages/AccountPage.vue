<template>
  <q-page padding>
    <div v-if="session">
      <div class="flex">
        <KoorAccount v-if="session" :session="session" />
        <AccountSettings />
      </div>
      <hr />
      <div class="q-mt-lg">
        <q-btn color="warning" @click="signOut">Sign Out</q-btn>
      </div>
    </div>
    <div v-else><AccountSignIn /></div>
  </q-page>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { supabase } from '../supabase'
import { signOutUser } from '../apis/supabase'
import KoorAccount from 'components/KoorAccount.vue'
import AccountSettings from 'components/AccountSettings.vue'
import AccountSignIn from 'components/AccountSignIn.vue'

const session = ref()

async function signOut() {
  await signOutUser()
}
onMounted(() => {
  supabase.auth.getSession().then(({ data }) => {
    session.value = data.session
  })

  supabase.auth.onAuthStateChange((_, _session) => {
    session.value = _session
  })
})
</script>
