<script setup lang="ts">
import type { ClusterStats } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import IopsStats from '~/components/cluster/services/IopsStats.vue';
import MdsService from '~/components/cluster/services/MdsService.vue';
import MgrService from '~/components/cluster/services/MgrService.vue';
import MonService from '~/components/cluster/services/MonService.vue';
import ObjectsBox from '~/components/cluster/services/ObjectsBox.vue';
import OsdService from '~/components/cluster/services/OsdService.vue';
import PgsBox from '~/components/cluster/services/PgsBox.vue';
import PoolsBox from '~/components/cluster/services/PoolsBox.vue';
import RgwService from '~/components/cluster/services/RgwService.vue';
import UsageStats from '~/components/cluster/services/UsageStats.vue';
import VolumesBox from '~/components/cluster/services/VolumesBox.vue';
import AlertsList from '~/components/cluster/services/AlertsList.vue';

defineProps<{
    stats: ClusterStats;
}>();
</script>

<template>
    <div class="w-full relative isolate">
        <!-- Stats -->
        <dl class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-5 px-2">
            <AlertsList :alerts="stats.crashes!" />
            <MonService :stats="stats.services!.mon!" />
            <MgrService :stats="stats.services!.mgr!" />
            <OsdService :stats="stats.services!.osd!" />
            <MdsService :stats="stats.services!.mds!" />
            <RgwService :stats="stats.services!.rgw!" />
            <ObjectsBox :stats="stats.data!.objects!" />
            <PoolsBox :stats="stats.data!.pools!" />
            <PgsBox :stats="stats.data!.pools!.pgs!" />
            <IopsStats :stats="stats.iops!" />
            <UsageStats :stats="stats.data!.usage!" />
            <VolumesBox :stats="stats.data!" />
        </dl>
    </div>
</template>
