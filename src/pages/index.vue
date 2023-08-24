<template>
    <div>
        <div class="row">
            <div class="col-4">
                <Radar />
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
    </div>
</template>
  
<script setup>
import Radar from '~/components/charts/cluster/Radar.vue'
import StatusDials from '~/components/StatusDials.vue'
import NodeSummaryList from '~/components/NodeSummaryList.vue'
import OSDHealthWidget from '~/components/OSDHealthWidget.vue'
</script>
