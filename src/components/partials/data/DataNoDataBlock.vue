<script lang="ts" setup>
import { MagnifyIcon } from 'mdi-vue3';
import { type DefineComponent } from 'vue';

const props = withDefaults(
    defineProps<{
        message?: string;
        icon?: DefineComponent;
        type?: string;
        focus?: Function;
    }>(),
    {
        icon: markRaw(MagnifyIcon),
        message: undefined,
        type: undefined,
        focus: undefined,
    },
);
function click() {
    if (props.focus) {
        props.focus();
    }
}
</script>

<template>
    <button
        type="button"
        :disabled="!focus"
        class="mt-4 relative block w-full p-12 text-center border-2 border-dashed rounded-lg bg-base-100 border-base-300 hover:border-base-400 focus:outline-none focus:ring-2 focus:ring-neutral focus:ring-offset-2"
        @click="click"
    >
        <component :is="icon" class="w-12 h-12 mx-auto text-base-500" />
        <span class="block mt-2 text-sm font-semibold text-base-500">
            <span v-if="message">
                {{ message }}
            </span>
            <span v-else>
                {{ `No ${type ?? 'Data'} found!` }}
            </span>
        </span>
    </button>
</template>
