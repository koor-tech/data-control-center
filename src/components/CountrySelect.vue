<template>
  <div>
    <q-select
      v-model="selectedCountry"
      label="Country"
      :options="countryList"
      optionLabel="name"
      optionValue="id"
    />
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { getCountries } from '../apis/supabase'

const countryList = ref(['USA', 'France', 'Japan'])
const selectedCountry = ref()

onMounted(async () => {
  countryList.value = await getCountries()
})
</script>

<style lang="scss" scoped></style>
