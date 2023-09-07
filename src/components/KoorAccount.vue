<template>
  <div>
    <form class="q-gutter-sm" @submit.prevent="">
      <h2 class="text-xl">Profile</h2>
      <div>
        <label for="username" class="block text-sm font-medium leading-6 text-gray-900">
          Username
        </label>
        <input id="username" name="username" disabled v-model="username"
          class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '~/store/auth';

const authStore = useAuthStore();
const { username } = storeToRefs(authStore);
</script>
