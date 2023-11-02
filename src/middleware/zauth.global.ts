import { RouteLocationNormalized } from 'vue-router';
import { useAuthStore } from '~/store/auth';

export default defineNuxtRouteMiddleware(async (to: RouteLocationNormalized, _: RouteLocationNormalized) => {
    const authStore = useAuthStore();
    const { accessToken } = storeToRefs(authStore);
    // Default is that a page requires authentication, but if it doesn't exit quickly
    if (to.meta.requiresAuth === false) {
        return true;
    }

    // Auth token is not null and only needed
    if (to.meta.authOnlyToken && accessToken.value !== null) {
        return true;
    }

    // Check if user has access token
    if (accessToken.value !== null) {
        // No route permission check, so we can go ahead and return true
        if (!to.meta.permission) {
            return true;
        }

        // Route has permission attached to it, check if user "can" go there
        if (can(to.meta.permission)) {
            // User has permission
            return true;
        } else {
            useNotificationsStore().dispatchNotification({
                title: 'No permissions!',
                content: 'No permissions to access page: ' + (to.name ? to.name?.toString() : to.path),
                type: 'warning',
            });

            if (accessToken.value !== null) {
                return navigateTo({
                    name: 'overview',
                });
            }
        }
    }

    // Only update the redirect query param if it isn't set already
    const redirect = to.query.redirect ?? to.fullPath;
    return navigateTo({
        name: 'auth-login',
        // save the location we were at to come back later
        query: { redirect },
    });
});
