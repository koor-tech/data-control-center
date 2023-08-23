<script lang="ts" setup>
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';
import HeroFull from '~/components/partials/HeroFull.vue';
import Login from '~/components/auth/Login.vue';
import Footer from '~/components/partials/Footer.vue';

useHead({
    title: 'Login',
});
definePageMeta({
    title: 'Login',
    requiresAuth: false,
});

const { $grpc } = useNuxtApp();

// Define an async function to fetch cluster stats
const fetchClusterStats = async () => {
    try {
        const statsClient = $grpc.getStatsClient();
        const emptyRequest = {}; // Create an instance of the request message
        const ClusterStatusResponse = {}; // Create an instance of the request message

        const response = await statsClient.getClusterStats(emptyRequest, ClusterStatusResponse);
        console.log('Cluster stats:', response);
        return response;
    } catch (error) {
        console.error('Error fetching cluster stats:', error);
        throw error;
    }
};

// Call the async function to fetch cluster stats
fetchClusterStats();
</script>

<template>
    <div class="h-full justify-between flex flex-col">
        <HeroFull>
            <ContentCenterWrapper>
                <Login />
            </ContentCenterWrapper>
        </HeroFull>
        <Footer />
    </div>
</template>
