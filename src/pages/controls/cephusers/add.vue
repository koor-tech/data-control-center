<script setup lang="ts">
import { max, min, required } from '@vee-validate/rules';
import { defineRule } from 'vee-validate';
import Alert from '~/components/partials/elements/Alert.vue';
import { useCephStore } from '~/store/ceph';

useHead({
    title: 'Add new Ceph dashboard user',
});
definePageMeta({
    title: 'Add new Ceph dashboard user',
    requiresAuth: true,
});

const cephStorage = useCephStore();
const { addCephUser } = cephStorage;
const { addError } = storeToRefs(cephStorage);

defineRule('required', required);
defineRule('min', min);
defineRule('max', max);

interface FormData {
    username: string;
    name: string;
    email: string;
    password: string;
    roles: string;
}

const { handleSubmit } = useForm<FormData>({
    validationSchema: {
        username: { required: true, min: 3, max: 24 },
        name: { required: true, min: 3, max: 60 },
        email: { required: true, min: 3, max: 60 },
        password: { required: true, min: 3, max: 60 },
        roles: { required: true },
    },
});

const onSubmit = handleSubmit(async function (values): Promise<any> {
    await addCephUser(values.username, values.name, values.email, values.password, values.roles);
    await navigateTo({ name: 'controls-cephusers' });
});
</script>
<template>
    <div class="p-2">
        <div class="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
            <div class="sm:mx-auto sm:w-full sm:max-w-sm">
                <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Add new ceph user</h2>
            </div>
        </div>
        <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form class="space-y-6" @submit="onSubmit">
                <div>
                    <label for="username" class="block text-sm font-medium leading-6 text-gray-900">Username *</label>
                    <div class="mt-2">
                        <VeeField
                            name="username"
                            type="text"
                            autocomplete="username"
                            placeholder="Username"
                            label="username"
                            class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        />
                        <VeeErrorMessage name="username" as="p" class="mt-2 text-sm text-error-400" />
                    </div>
                </div>

                <div>
                    <label for="name" class="block text-sm font-medium leading-6 text-gray-900">Name *</label>
                    <div class="mt-2">
                        <VeeField
                            name="name"
                            type="text"
                            autocomplete="name"
                            placeholder="Name"
                            label="name"
                            class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        />
                        <VeeErrorMessage name="name" as="p" class="mt-2 text-sm text-error-400" />
                    </div>
                </div>
                <div>
                    <label for="email" class="block text-sm font-medium leading-6 text-gray-900">Email *</label>
                    <div class="mt-2">
                        <VeeField
                            name="email"
                            type="email"
                            autocomplete="email"
                            placeholder="Email"
                            label="email"
                            class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        />
                        <VeeErrorMessage name="email" as="p" class="mt-2 text-sm text-error-400" />
                    </div>
                </div>
                <div>
                    <label for="password" class="block text-sm font-medium leading-6 text-gray-900">Password *</label>
                    <div class="mt-2">
                        <VeeField
                            name="password"
                            type="password"
                            autocomplete="password"
                            placeholder="Password"
                            label="password"
                            class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        />
                        <VeeErrorMessage name="password" as="p" class="mt-2 text-sm text-error-400" />
                    </div>
                </div>

                <div>
                    <div class="flex items-center justify-between">
                        <div>
                            <label for="roles" class="block text-sm font-medium text-gray-700">Roles</label>
                            <VeeField
                                as="select"
                                name="roles"
                                label="roles"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            >
                                <option value="administrator">administrator</option>
                                <option value="read-only">Read-Only</option>
                                <option value="block-manager">Block Manager</option>
                            </VeeField>
                            <VeeErrorMessage name="roles" as="p" class="mt-2 text-sm text-error-400" />
                        </div>
                    </div>
                </div>
                <div>
                    <button
                        type="submit"
                        class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                    >
                        Add user
                    </button>
                </div>
            </form>
            <Alert v-if="addError" title="Error during add user" :message="addError" />
        </div>
    </div>
</template>
