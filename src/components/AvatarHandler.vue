<template>
  <div class="q-mt-md">
    <q-avatar square :size="size + 'em'">
      <q-img v-if="src" :src="src" />
      <q-icon v-else name="face" />
    </q-avatar>
    <q-file
      v-model="avatarFile"
      label="Click here to pick a new avatar"
      filled
      style="max-width: 300px"
    />
  </div>
</template>

<script setup>
import { ref, toRefs, watch } from 'vue'
import { downloadAvatar, uploadAvatar } from '../apis/supabase'

const prop = defineProps(['path', 'size'])
const { path, size } = toRefs(prop)

const emit = defineEmits(['upload', 'upload:path'])
const uploading = ref(false)
const src = ref('')
const avatarFile = ref(null)
const files = ref()

const downloadImage = async () => {
  try {
    const { data, error } = await downloadAvatar(path.value)
    if (error) throw error
    src.value = URL.createObjectURL(data)
  } catch (error) {
    console.error('Error downloading image: ', error.message)
  }
}

const uploadNewAvatar = async (e) => {
  try {
    uploading.value = true
    if (!avatarFile.value) {
      throw new Error('You must select an image to upload.')
    }

    const file = avatarFile.value
    const fileExt = file.name.split('.').pop()
    const filePath = `${Math.random()}.${fileExt}`
    let { error: uploadError } = await uploadAvatar(filePath, file)
    if (uploadError) throw uploadError
    emit('update:path', filePath)
    emit('upload')
  } catch (error) {
    alert(error.message)
  } finally {
    uploading.value = false
  }
}

watch(avatarFile, () => {
  if (avatarFile.value) uploadNewAvatar()
})
watch(path, () => {
  if (path.value) downloadImage()
})
</script>

<style lang="scss" scoped></style>
