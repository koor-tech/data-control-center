import { RouteLocationNormalized } from 'vue-router';
import { useAuthStore } from '~/store/auth';
import { useNotificationsStore } from '~/store/notifications';
import slug from '~/utils/slugify';
import { toDate } from '~/utils/time';

export default defineNuxtRouteMiddleware(async (to: RouteLocationNormalized, from: RouteLocationNormalized) => {
    // Default is that a page requires authentication
    if (!to.meta.hasOwnProperty('requiresAuth') || to.meta.requiresAuth) {
        const authStore = useAuthStore();
        const { accessToken, permissions } = storeToRefs(authStore);

        if (to.meta.authOnlyToken && accessToken.value !== null) {
            return true;
        }

        // Check if user has access token
        if (accessToken.value !== null) {
            // If the user has an acitve char, check for perms otherwise, redirect to char selector
            if (permissions.value.includes('superuser')) {
                return true;
            }

            // Route has permission attached to it, check if user has required permission
            if (to.meta.permission) {
                const perm = slug(to.meta.permission as string);
                if (permissions.value.includes(perm)) {
                    // User has permission
                    return true;
                } else {
                    useNotificationsStore().dispatchNotification({
                        title: 'No permissions!',
                        content: 'No permissions to access page: ' + (to.name ? to.name?.toString() : to.path),
                        type: 'warning',
                    });

                    if (accessToken.value) {
                        return navigateTo({
                            name: 'index',
                        });
                    }
                }
            } else {
                // No route permission check, so we can go ahead and return
                return true;
            }
        }

        // Only update the redirect query param if it isn't set already
        const redirect = to.query.redirect ?? to.fullPath;
        return navigateTo({
            name: 'auth-login',
            // save the location we were at to come back later
            query: { redirect: redirect },
        });
    }

    return true;
});
