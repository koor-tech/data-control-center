<script lang="ts" setup>
import { Dialog, DialogPanel, Menu, MenuButton, MenuItem, MenuItems, TransitionChild, TransitionRoot } from '@headlessui/vue';
import {
    AccountDetailsIcon,
    AccountIcon,
    AntennaIcon,
    ChevronRightIcon,
    CloseIcon,
    GithubIcon,
    GlassesIcon,
    HelpBoxMultipleIcon,
    HomeIcon,
    MenuIcon,
    PlusBoxMultipleIcon,
    SchoolIcon,
    TuneIcon,
    TwitterIcon,
    YoutubeIcon,
} from 'mdi-vue3';
import { type DefineComponent } from 'vue';
import { useAuthStore } from '~/store/auth';
import { type RoutesNamedLocations } from '@typed-router';
import { useBreadcrumbs } from '~/composables/breadcrumbs';

const authStore = useAuthStore();
const { accessToken } = storeToRefs(authStore);
const router = useRouter();

type SidebarNavigationitem = {
    name: string;
    permission?: string;
    icon: DefineComponent;
    position: 'top' | 'bottom';
    current: boolean;
    loggedIn?: boolean;
    charSelected?: boolean;
};

const sidebarNavigation = ref<
    (SidebarNavigationitem &
        (
            | {
                  href: string;
                  external: true;
              }
            | { href: RoutesNamedLocations; external?: false }
        ))[]
>([
    {
        name: 'Overview',
        href: { name: 'index' },
        icon: markRaw(GlassesIcon),
        position: 'top',
        current: false,
    },
    {
        name: 'Stats',
        href: { name: 'health' },
        icon: markRaw(AntennaIcon),
        position: 'top',
        current: false,
    },
    {
        name: 'Controls',
        href: { name: 'controls' },
        icon: markRaw(TuneIcon),
        position: 'top',
        current: false,
    },
    {
        name: 'Recommender',
        href: { name: 'recommender' },
        icon: markRaw(PlusBoxMultipleIcon),
        position: 'top',
        current: false,
    },
    {
        name: 'Troubleshooting',
        href: { name: 'troubleshoot' },
        icon: markRaw(HelpBoxMultipleIcon),
        position: 'top',
        current: false,
    },

    // Bottom
    {
        name: 'Koor Account',
        external: true,
        href: 'https://account.koor.tech/login',
        icon: markRaw(AccountDetailsIcon),
        position: 'bottom',
        current: false,
    },
    {
        name: 'Knowledge Center',
        href: 'https://kb.koor.tech',
        external: true,
        icon: markRaw(SchoolIcon),
        position: 'bottom',
        current: false,
    },
    {
        name: 'Github',
        href: 'https://github.com/koor-tech',
        external: true,
        icon: markRaw(GithubIcon),
        position: 'bottom',
        current: false,
    },
    {
        name: 'X (Twitter)',
        href: 'https://twitter.com/koor_tech',
        external: true,
        icon: markRaw(TwitterIcon),
        position: 'bottom',
        current: false,
    },
    {
        name: 'YouTube',
        href: 'https://www.youtube.com/@koor-tech',
        external: true,
        icon: markRaw(YoutubeIcon),
        position: 'bottom',
        current: false,
    },
]);
const userNavigation = ref<{ name: string; href: RoutesNamedLocations; permission?: string }[]>([
    { name: 'Login', href: { name: 'auth-login' } },
]);
const breadcrumbs = useBreadcrumbs();
const mobileMenuOpen = ref(false);

onMounted(() => {
    updateUserNav();
    updateActiveItem();
});

function updateUserNav(): void {
    userNavigation.value.length = 0;
    if (accessToken.value) {
        userNavigation.value.push({ name: 'Sign Out', href: { name: 'auth-logout' } });
    }
    if (userNavigation.value.length === 0) {
        userNavigation.value = [{ name: 'Login', href: { name: 'auth-login' } }];
    }
}

function updateActiveItem(): void {
    const route = router.currentRoute.value;
    if (route.name) {
        sidebarNavigation.value.forEach((e) => {
            if (e.external) return;

            if (route.name.toLowerCase().includes(e.href.name.toLowerCase())) {
                e.current = true;
            } else {
                e.current = false;
            }
        });
    } else {
        sidebarNavigation.value.forEach((e) => (e.current = false));
    }
}

watch(accessToken, () => updateUserNav());

watch(router.currentRoute, () => {
    updateActiveItem();
});

const appVersion = accessToken ? ' v' + __APP_VERSION__ + (import.meta.env.DEV ? '-dev' : '-prod') : '';
</script>

<template>
    <div class="flex h-dscreen">
        <!-- Sidebar -->
        <div class="hidden overflow-y-auto bg-accent-600 w-28 md:block">
            <div class="flex flex-col items-center w-full py-6 h-full">
                <div class="flex items-center flex-shrink-0">
                    <NuxtLink :to="{ name: 'index' }" aria-current-value="page">
                        <KoorLogo class="w-20 h-20" :title="`Koor Data Control Center ${appVersion}`" />
                    </NuxtLink>
                </div>
                <div class="flex-grow w-full px-2 mt-6 space-y-1 text-center">
                    <template v-if="accessToken">
                        <!-- @vue-ignore nuxt typed router makes sure the "href" is correct for internal links -->
                        <NuxtLink
                            v-for="item in sidebarNavigation.filter(
                                (e) => e.position === 'top' && (!e.permission || can(e.permission)),
                            )"
                            :key="item.name"
                            :external="item.external"
                            :to="item.href"
                            :class="[
                                item.current
                                    ? 'bg-accent-100/20 text-neutral font-bold'
                                    : 'text-accent-100 hover:bg-accent-100/10 hover:text-neutral font-medium',
                                'hover:transition-all group flex w-full flex-col items-center rounded-md p-3 text-xs my-2',
                            ]"
                            :aria-current="item.current ? 'page' : undefined"
                        >
                            <component
                                :is="item.icon"
                                :class="[item.current ? 'text-neutral' : 'text-accent-100 group-hover:text-neutral', 'h-6 w-6']"
                                aria-hidden="true"
                            />
                            <span class="mt-2">{{ item.name }}</span>
                        </NuxtLink>
                    </template>
                </div>
                <div class="flex-initial w-full px-2 space-y-1 text-center">
                    <!-- @vue-ignore nuxt typed router makes sure the "href" is correct for internal links -->
                    <NuxtLink
                        v-for="item in sidebarNavigation.filter(
                            (e) => e.position === 'bottom' && (!e.permission || can(e.permission)),
                        )"
                        :key="item.name"
                        :external="item.external"
                        :to="item.href"
                        :class="[
                            item.current
                                ? 'bg-accent-100/20 text-neutral font-bold'
                                : 'text-accent-100 hover:bg-accent-100/10 hover:text-neutral font-medium',
                            'hover:transition-all group flex w-full flex-col items-center rounded-md p-1 text-xs my-2',
                        ]"
                        :aria-current="item.current ? 'page' : undefined"
                    >
                        <component
                            :is="item.icon"
                            :class="[item.current ? 'text-neutral' : 'text-accent-100 group-hover:text-neutral', 'h-6 w-6']"
                            aria-hidden="true"
                        />
                        <span class="mt-2">{{ item.name }}</span>
                    </NuxtLink>
                </div>
            </div>
        </div>

        <!-- Mobile Sidebar -->
        <TransitionRoot as="template" :show="mobileMenuOpen">
            <Dialog as="div" class="relative z-20 md:hidden" @close="mobileMenuOpen = false">
                <TransitionChild
                    as="template"
                    enter="transition-opacity ease-linear duration-300"
                    enter-from="opacity-0"
                    enter-to="opacity-100"
                    leave="transition-opacity ease-linear duration-300"
                    leave-from="opacity-100"
                    leave-to="opacity-0"
                >
                    <div class="fixed inset-0 bg-opacity-75 bg-base-900/10" />
                </TransitionChild>

                <div class="fixed inset-0 z-40 flex">
                    <TransitionChild
                        as="template"
                        enter="transition ease-in-out duration-300 transform"
                        enter-from="-translate-x-full"
                        enter-to="translate-x-0"
                        leave="transition ease-in-out duration-300 transform"
                        leave-from="translate-x-0"
                        leave-to="-translate-x-full"
                    >
                        <DialogPanel class="relative flex flex-col flex-1 w-full max-w-xs pt-5 pb-4 bg-accent-600">
                            <TransitionChild
                                as="template"
                                enter="ease-in-out duration-300"
                                enter-from="opacity-0"
                                enter-to="opacity-100"
                                leave="ease-in-out duration-300"
                                leave-from="opacity-100"
                                leave-to="opacity-0"
                            >
                                <div class="absolute p-1 -right-3 top-1 -mr-14">
                                    <button
                                        type="button"
                                        class="flex items-center justify-center w-12 h-12 rounded-full focus:outline-none ring-2 ring-neutral"
                                        @click="mobileMenuOpen = false"
                                    >
                                        <CloseIcon class="w-6 h-6 text-neutral" aria-hidden="true" />
                                        <span class="sr-only">Close Sidebar</span>
                                    </button>
                                </div>
                            </TransitionChild>
                            <div class="flex items-center flex-shrink-0 px-4">
                                <KoorLogo class="w-20 h-20" :title="`Koor Data Control Center ${appVersion}`" />
                            </div>
                            <div class="flex-grow h-0 px-2 mt-5 overflow-y-auto">
                                <nav class="flex flex-col h-full">
                                    <div class="space-y-1">
                                        <NuxtLink
                                            v-if="!accessToken"
                                            :to="{ name: 'index' }"
                                            class="text-accent-100 hover:bg-accent-100/10 hover:text-neutral font-medium group flex items-center rounded-md py-2 px-3 text-sm"
                                            @click="mobileMenuOpen = false"
                                        >
                                            <HomeIcon
                                                class="text-accent-100 group-hover:text-neutral mr-3 h-6 w-6"
                                                aria-hidden="true"
                                            />
                                            <span>Home</span>
                                        </NuxtLink>
                                        <template v-else-if="accessToken">
                                            <NuxtLink
                                                v-for="item in sidebarNavigation.filter(
                                                    (e) => e.position === 'top' && (!e.permission || can(e.permission)),
                                                )"
                                                :key="item.name"
                                                :external="item.external"
                                                :to="item.href"
                                                :class="[
                                                    item.current
                                                        ? 'bg-accent-100/20 text-neutral font-bold'
                                                        : 'text-accent-100 hover:bg-accent-100/10 hover:text-neutral font-medium',
                                                    'group flex items-center rounded-md py-2 px-3 text-sm',
                                                ]"
                                                :aria-current="item.current ? 'page' : undefined"
                                                @click="mobileMenuOpen = false"
                                            >
                                                <component
                                                    :is="item.icon"
                                                    :class="[
                                                        item.current
                                                            ? 'text-neutral'
                                                            : 'text-accent-100 group-hover:text-neutral',
                                                        'mr-3 h-6 w-6',
                                                    ]"
                                                    aria-hidden="true"
                                                />
                                                <span>{{ item.name }}</span>
                                            </NuxtLink>
                                        </template>
                                    </div>
                                </nav>
                            </div>
                            <div class="flex-initial px-2 overflow-y-auto">
                                <nav class="flex flex-col h-full">
                                    <div class="space-y-1">
                                        <!-- @vue-ignore nuxt typed router makes sure the "href" is correct for internal links -->
                                        <NuxtLink
                                            v-for="item in sidebarNavigation.filter(
                                                (e) => e.position === 'bottom' && (!e.permission || can(e.permission)),
                                            )"
                                            :key="item.name"
                                            :external="item.external"
                                            :to="item.href"
                                            :class="[
                                                item.current
                                                    ? 'bg-accent-100/20 text-neutral font-bold'
                                                    : 'text-accent-100 hover:bg-accent-100/10 hover:text-neutral font-medium',
                                                'group flex items-center rounded-md py-2 px-3 text-sm',
                                            ]"
                                            :aria-current="item.current ? 'page' : undefined"
                                            @click="mobileMenuOpen = false"
                                        >
                                            <component
                                                :is="item.icon"
                                                :class="[
                                                    item.current ? 'text-neutral' : 'text-accent-100 group-hover:text-neutral',
                                                    'mr-3 h-6 w-6',
                                                ]"
                                                aria-hidden="true"
                                            />
                                            <span>{{ item.name }}</span>
                                        </NuxtLink>
                                    </div>
                                </nav>
                            </div>
                        </DialogPanel>
                    </TransitionChild>
                    <div class="flex-shrink-0 w-14" aria-hidden="true"></div>
                </div>
            </Dialog>
        </TransitionRoot>

        <!-- Content area -->
        <div class="flex flex-col flex-1 overflow-hidden">
            <header class="w-full">
                <div class="relative z-10 flex flex-shrink-0 h-16 bg-accent-950">
                    <button
                        type="button"
                        class="px-4 text-neutral focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500 md:hidden"
                        @click="mobileMenuOpen = true"
                    >
                        <span class="sr-only">Open Sidebar</span>
                        <MenuIcon class="w-6 h-6" aria-hidden="true" />
                    </button>
                    <div class="flex justify-between flex-1 px-4 sm:px-6">
                        <div class="flex flex-1">
                            <nav class="flex" aria-label="Breadcrumb">
                                <ol role="list" class="flex items-center space-x-4">
                                    <li>
                                        <div>
                                            <NuxtLink
                                                :to="{
                                                    name: 'index',
                                                }"
                                                class="text-base-400 hover:text-neutral hover:transition-colors"
                                            >
                                                <HomeIcon class="flex-shrink-0 w-5 h-5" aria-hidden="true" />
                                                <span class="sr-only">Home</span>
                                            </NuxtLink>
                                        </div>
                                    </li>
                                    <li v-for="(page, key) in breadcrumbs" :key="key">
                                        <div class="flex items-center">
                                            <ChevronRightIcon class="flex-shrink-0 w-5 h-5 text-base-400" aria-hidden="true" />
                                            <!-- @vue-ignore the route should be valid, as we construct it based on our pages -->
                                            <NuxtLink
                                                :to="page.href"
                                                :class="[
                                                    page.current ? 'font-bold text-base-200' : 'font-medium text-base-400',
                                                    'ml-4 text-sm hover:text-neutral hover:transition-colors',
                                                ]"
                                                :aria-current="page.current ? 'page' : undefined"
                                            >
                                                {{ page.title }}
                                            </NuxtLink>
                                        </div>
                                    </li>
                                </ol>
                            </nav>
                        </div>
                        <div class="flex items-center ml-2 space-x-4 sm:ml-6 sm:space-x-6">
                            <!-- Account dropdown -->
                            <Menu as="div" class="relative flex-shrink-0">
                                <div>
                                    <MenuButton
                                        class="flex text-sm rounded-full bg-accent-950 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
                                    >
                                        <span class="sr-only"> Open User Menu </span>
                                        <AccountIcon
                                            class="w-auto h-10 rounded-full hover:transition-colors text-base-300 bg-base-800 fill-base-300 hover:text-base-100 hover:fill-base-100"
                                        />
                                    </MenuButton>
                                </div>
                                <transition
                                    enter-active-class="transition duration-100 ease-out"
                                    enter-from-class="transform scale-95 opacity-0"
                                    enter-to-class="transform scale-100 opacity-100"
                                    leave-active-class="transition duration-75 ease-in"
                                    leave-from-class="transform scale-100 opacity-100"
                                    leave-to-class="transform scale-95 opacity-0"
                                >
                                    <MenuItems
                                        class="absolute right-0 w-48 py-1 mt-2 origin-top-right rounded-md shadow-float bg-accent-950 ring-1 ring-base-100 ring-opacity-5 focus:outline-none z-40"
                                    >
                                        <MenuItem
                                            v-for="item in userNavigation.filter(
                                                (e) => e.permission === undefined || can(e.permission),
                                            )"
                                            :key="item.name"
                                            v-slot="{ active }"
                                        >
                                            <!-- @vue-ignore nuxt typed router makes sure the "href" is correct for internal links -->
                                            <NuxtLink
                                                :to="item.href"
                                                :class="[
                                                    active ? 'bg-base-800' : '',
                                                    'block px-4 py-2 text-sm text-neutral hover:transition-colors',
                                                ]"
                                            >
                                                {{ item.name }}
                                            </NuxtLink>
                                        </MenuItem>
                                    </MenuItems>
                                </transition>
                            </Menu>
                        </div>
                    </div>
                </div>
            </header>

            <main class="h-full overflow-y-auto">
                <section aria-labelledby="primary-heading" class="h-full min-w-0 lg:order-last">
                    <slot></slot>
                </section>
            </main>
        </div>
    </div>
</template>
