<script lang="ts" setup>
import { useAuthStore } from '~/store/auth';
import { useNotificationsStore } from '~/store/notifications';
import LoginForm from '~/components/auth/LoginForm.vue';

const authStore = useAuthStore();
const notifications = useNotificationsStore();
const route = useRoute();

const { setAccessToken } = authStore;

const query = route.query;
// `t` and `exp` set, means social login was successful
if (query.t && query.t !== '' && query.exp) {
    setAccessToken(query.t as string, BigInt(query.exp as string));

    notifications.dispatchNotification({
        title: 'OAuth2 Login successful!',
        content: 'Redirecting you to the data-control-center Overview.',
        type: 'info',
    });

    navigateTo({ name: 'index' });
} else if (query.oauth2Login && query.oauth2Login === 'failed') {
    // `oauth2Login` can be `failed` (with `reason`)
    const reason = query.reason ?? 'N/A';

    notifications.dispatchNotification({
        title: 'OAuth2 Login failed!',
        content: 'Reason:' + reason.toString(),
        type: 'error',
    });
}
</script>

<template>
    <div class="max-w-lg w-full sm:w-[32rem] mx-auto">
        <div class="px-4 py-8 rounded-lg bg-accent-950 sm:px-10">
            <KoorLogo class="h-auto mx-auto mb-2 w-36" />
            <LoginForm />
        </div>
    </div>
</template>
