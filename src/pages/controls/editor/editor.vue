<script setup lang="ts">
import { CheckIcon, ChevronDownIcon, FileDocumentEditIcon } from 'mdi-vue3';
import { Listbox, ListboxButton, ListboxOption, ListboxOptions } from '@headlessui/vue';
import { ConnectError } from '@connectrpc/connect';
import { ref } from 'vue';
import MonacoEditor from '~/components/editor/MonacoEditor.vue';
import { useClusterStore } from '~/store/cluster';
import AlertSuccess from '~/components/editor/AlertSuccess.vue';
useHead({
    title: 'Koor objects Editor',
});
definePageMeta({
    title: 'Koor objects Editor',
    requiresAuth: true,
});

const { $grpc } = useNuxtApp();

const clusterStore = useClusterStore();

interface CephResource {
    name: string;
    namespace: string;
    kind: string;
    content: string;
    object: string;
    current: boolean;
}

const isSuccess = ref(false);
const savingStatus = ref('');

function hideAlert() {
    isSuccess.value = false;
}

const { data: resources } = useLazyAsyncData('Resources', async () => {
    try {
        const resp = await clusterStore.getCustomResources();

        return resp.resources?.resources.map(
            (item, index) =>
                ({
                    name: item.name,
                    namespace: item.namespace,
                    kind: item.kind,
                    object: item.object,
                    content: item.content,
                    current: index === 0,
                }) as CephResource,
        );
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});

// Ref for the selected item
const selected = ref<CephResource | null>(null);

watch(
    resources,
    (newValue) => {
        if (Array.isArray(newValue) && newValue.length > 0) {
            selected.value = newValue[0] as CephResource;
        } else {
            selected.value = { name: '', content: '', namespace: '', kind: '', object: '', current: false };
        }
    },
    { immediate: true },
);

const monacoEditor = ref<typeof MonacoEditor | null>(null);

function delay(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

async function saveChanges() {
    savingStatus.value = '.........saving';
    if (monacoEditor.value) {
        const updatedContent = monacoEditor.value.getCurrentContent();
        const resource = selected.value;
        console.log(selected.value);
        const res = await clusterStore.updateResource(
            resource?.name,
            resource?.namespace,
            resource?.kind,
            resource?.object,
            updatedContent,
        );
        await delay(1000);
        if (res) {
            isSuccess.value = true;
        }
        savingStatus.value = '';
    }
}
</script>

<template>
    <div class="p-2">
        <div class="mt-10 sm:mx-auto sm:w-2/3 mb-10">
            <div class="overflow-hidden rounded-lg bg-white shadow">
                <div class="bg-white p-6">
                    <div>
                        <div class="sm:flex sm:space-x-5">
                            <div class="flex-shrink-0">
                                <FileDocumentEditIcon class="mx-auto h-20 w-20 rounded-full" />
                            </div>
                            <div class="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left">
                                <p class="text-xl font-bold text-gray-900 sm:text-2xl">Koor Editor</p>
                                <p class="pt-3">Please select the Ceph object from the dropdown that you want to edit</p>
                            </div>
                        </div>
                    </div>
                    <div class="flex pt-4 w-full">
                        <Listbox v-if="selected" v-model="selected" as="div" class="w-full">
                            <div class="relative">
                                <div class="w-full inline-flex divide-x divide-accent-500 rounded-md shadow-sm">
                                    <div
                                        class="w-full inline-flex items-center gap-x-1.5 rounded-l-md bg-accent-500 px-3 py-2 text-white shadow-sm"
                                    >
                                        <CheckIcon class="-ml-0.5 h-5 w-5" aria-hidden="true" />
                                        <p class="text-sm font-semibold">{{ selected.kind }} / {{ selected.name }}</p>
                                    </div>
                                    <ListboxButton
                                        class="inline-flex items-center rounded-l-none rounded-r-md bg-accent-500 p-2 hover:bg-accent-500 focus:outline-none focus:ring-2 focus:ring-accent-500 focus:ring-offset-2 focus:ring-offset-gray-50"
                                    >
                                        <span class="sr-only">Change Ceph object</span>
                                        <ChevronDownIcon class="h-5 w-5 text-white" aria-hidden="true" />
                                    </ListboxButton>
                                </div>
                                <transition
                                    leave-active-class="transition ease-in duration-100"
                                    leave-from-class="opacity-100"
                                    leave-to-class="opacity-0"
                                >
                                    <ListboxOptions
                                        class="w-full absolute z-10 w-72 origin-top-right divide-y divide-gray-200 overflow-hidden rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
                                    >
                                        <ListboxOption
                                            v-for="option in resources"
                                            :key="option?.name"
                                            v-slot="{ active, current }"
                                            as="template"
                                            :value="option"
                                        >
                                            <li
                                                :class="[
                                                    active ? 'bg-accent-500 text-white' : 'text-gray-900',
                                                    'cursor-default select-none p-4 text-sm',
                                                ]"
                                            >
                                                <div class="flex flex-col">
                                                    <div class="flex justify-between">
                                                        <p :class="current ? 'font-semibold' : 'font-normal'">
                                                            {{ option?.kind }} / {{ option?.name }}
                                                        </p>
                                                        <span v-if="current" :class="active ? 'text-white' : 'text-indigo-600'">
                                                            <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                                        </span>
                                                    </div>
                                                </div>
                                            </li>
                                        </ListboxOption>
                                    </ListboxOptions>
                                </transition>
                            </div>
                        </Listbox>
                    </div>
                    <div class="pt-4">
                        <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                            <AlertSuccess v-if="isSuccess" @click.prevent="hideAlert" />
                            <dl v-if="selected" class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-1">
                                <div
                                    class="overflow-hidden rounded-l-md rounded-r-md border-t border-gray-200 bg-gray-50 shadow sm:p-3 pl-4"
                                >
                                    <dd class="tracking-tight text-gray-500">
                                        {{ selected.kind }} / {{ selected.name }}
                                        <span class="text-sm font-bold text-accent-500">{{ savingStatus }}</span>
                                    </dd>
                                </div>
                            </dl>
                            <MonacoEditor
                                v-if="selected"
                                ref="monacoEditor"
                                width="800"
                                height="500"
                                theme="vs-dark"
                                language="yaml"
                                :content="selected.content"
                                class="border-t border-gray-200"
                            />
                        </div>
                    </div>
                </div>
                <div
                    v-if="selected"
                    class="grid grid-cols-1 divide-y divide-gray-200 border-t border-gray-200 bg-gray-50 sm:grid-cols-3 sm:divide-x sm:divide-y-0"
                >
                    <div class="px-6 py-5 text-center text-sm font-medium">
                        <div class="text-gray-600">
                            Kind:
                            <span class="text-base font-medium">{{ selected.kind }}</span>
                        </div>
                    </div>
                    <div class="px-6 py-5 text-center text-sm font-medium">
                        <div class="text-gray-600">
                            Namespace: <span class="text-base font-medium">{{ selected.namespace }}</span>
                        </div>
                    </div>
                    <div class="px-6 py-5 text-center text-sm font-medium">
                        <span class="text-gray-600">
                            <button
                                type="button"
                                class="rounded-md bg-accent-500 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-accent-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                @click.prevent="saveChanges"
                            >
                                Save Changes
                            </button>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
