<template>
  <div class="q-pa-md" style="width: 350px">
    <div class="text-h3 text-center">Profile</div>
    <q-form class="q-gutter-sm" @submit.prevent="updateProfile">
      <q-input id="email" label="Email" outlined v-model="session.user.email" disable />
      <q-input id="username" label="Alias" outlined v-model="username" />
      <q-input id="full_name" label="Full Name" outlined v-model="full_name" />
      <q-input id="website" label="Website" outlined type="url" v-model="website" />
      <AvatarHandler v-model:path="avatar_url" @upload="updateProfile" size="10" />
      <div>
        <q-btn
          type="submit"
          color="primary"
          :label="loading ? 'Loading' : 'Update'"
          :disabled="loading"
        />
      </div>
    </q-form>
  </div>
</template>

<script setup>
import { getUserProfile, updateUserProfile, signOutUser } from '../apis/supabase'
import { onMounted, ref, toRefs } from 'vue'
import AvatarHandler from './AvatarHandler.vue'

const props = defineProps(['session'])
const { session } = toRefs(props)

const loading = ref(true)
const username = ref('')
const full_name = ref('')
const website = ref('')
const avatar_url = ref('')

onMounted(() => {
  getProfile()
})

async function getProfile() {
  try {
    loading.value = true
    const { user } = session.value
    let { data, error, status } = await getUserProfile(user.id)
    if (error && status != 406) throw error
    if (data) {
      username.value = data.username
      full_name.value = data.full_name
      website.value = data.website
      avatar_url.value = data.avatar_url
    }
  } catch (error) {
    alert(error.message)
  } finally {
    loading.value = false
  }
}

async function updateProfile() {
  try {
    loading.value = true
    const { user } = session.value
    const updates = {
      id: user.id,
      username: username.value,
      full_name: full_name.value,
      website: website.value,
      avatar_url: avatar_url.value,
      updated_at: new Date(),
    }
    const { error } = await updateUserProfile(updates)
    if (error) throw error
  } catch (error) {
    alert(error.message)
  } finally {
    loading.value = false
  }
}
</script>

<style lang="scss" scoped></style>
