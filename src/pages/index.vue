<template>
    <q-page padding>
        <div class="row">
            <div class="col-4">
                <ClusterRadar />
            </div>
            <div class="col-8">
                <StatusDials />
            </div>
        </div>
        <div class="row">
            <div class="col-4">
                <NodeSummaryList />
            </div>
            <div class="col-8">
                <OSDHealthWidget />
            </div>
        </div>
    </q-page>
</template>
  
<script setup>
import ClusterRadar from '~/components/ClusterRadar.vue'
import StatusDials from '~/components/StatusDials.vue'
import NodeSummaryList from '~/components/NodeSummaryList.vue'
import OSDHealthWidget from '~/components/OSDHealthWidget.vue'
</script>
