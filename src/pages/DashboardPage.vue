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
import { ref } from 'vue'
import ClusterRadar from 'components/ClusterRadar.vue'
import StatusDials from 'components/StatusDials.vue'
import NodeSummaryList from 'components/NodeSummaryList.vue'
import OSDHealthWidget from 'components/OSDHealthWidget.vue'
</script>

<style scoped lang="scss">
.graph {
  border: 5px groove $grey-3;
}
.caption {
  margin-left: 5px;
}
</style>
