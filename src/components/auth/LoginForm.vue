<script lang="ts" setup>
// eslint-disable-next-line camelcase
import { alpha_dash, max, min, required } from '@vee-validate/rules';
import { defineRule } from 'vee-validate';
import Alert from '~/components/partials/elements/Alert.vue';
import { useAuthStore } from '~/store/auth';
import { useConfigStore } from '~/store/config';

const authStore = useAuthStore();
const { loginError } = storeToRefs(authStore);
const { doLogin } = authStore;

const configStore = useConfigStore();
const { appConfig } = storeToRefs(configStore);

defineRule('required', required);
defineRule('min', min);
defineRule('max', max);
defineRule('alpha_dash', alpha_dash);

interface FormData {
    registrationToken: number;
    username: string;
    password: string;
}

const { handleSubmit, meta } = useForm<FormData>({
    validationSchema: {
        username: { required: true, min: 3, max: 24, alpha_dash: true },
        password: { required: true, min: 6, max: 70 },
    },
});

const onSubmit = handleSubmit(async (values): Promise<void> => await doLogin(values.username, values.password));
</script>

<template>
    <h2 class="pb-4 text-3xl text-center text-white">Login</h2>

    <form class="my-2 space-y-6" @submit="onSubmit">
        <template v-if="appConfig.login.providers.length === 0">
            <div>
                <label for="username" class="sr-only"> Username </label>
                <div>
                    <VeeField
                        name="username"
                        type="text"
                        autocomplete="username"
                        placeholder="Username"
                        label="Username"
                        class="block w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                    />
                    <VeeErrorMessage name="username" as="p" class="mt-2 text-sm text-error-400" />
                </div>
            </div>
            <div>
                <label for="password" class="sr-only"> Password </label>
                <div>
                    <VeeField
                        name="password"
                        type="password"
                        autocomplete="current-password"
                        placeholder="Password"
                        label="Password"
                        class="block w-full rounded-md border-0 py-1.5 bg-base-700 text-neutral placeholder:text-base-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                    />
                    <VeeErrorMessage name="password" as="p" class="mt-2 text-sm text-error-400" />
                </div>
            </div>

            <div>
                <button
                    type="submit"
                    class="flex justify-center w-full px-3 py-2 text-sm font-semibold transition-colors rounded-md text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                    :disabled="!meta.valid"
                    :class="[
                        !meta.valid
                            ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                            : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                    ]"
                >
                    Login
                </button>
            </div>
        </template>

        <div v-else class="my-4 space-y-2">
            <div v-for="prov in appConfig.login.providers" :key="prov.name">
                <NuxtLink
                    :external="true"
                    :to="`/api/oauth2/login/${prov.name}`"
                    class="flex justify-center w-full px-3 py-2 text-sm font-semibold transition-colors rounded-md bg-primary-500 text-neutral hover:bg-primary-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-base-300"
                >
                    {{ prov.label }} Login
                </NuxtLink>
            </div>
        </div>
    </form>

    <Alert v-if="loginError" title="Error during Login" :message="loginError" />
</template>
