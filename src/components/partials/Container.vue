<template>
    <div class="bg-white rounded-lg shadow dark:bg-gray-800 px-2 py-2">
        <div v-if="$slots.title" class="flex justify-center mb-3 items-center">
            <h3 class="text-xl font-bold leading-none text-gray-900 dark:text-white pr-1">
                <slot name="title" />
            </h3>
        </div>
        <div>
            <slot />
        </div>
    </div>
</template>
