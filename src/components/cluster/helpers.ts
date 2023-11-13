import { ClusterHealth } from '~~/gen/ts/api/resources/stats/v1/stats_pb';

export function clusterHealthToBGColor(status: ClusterHealth | undefined): string {
    switch (status) {
        case ClusterHealth.UNSPECIFIED:
        case ClusterHealth.OFFLINE:
        case ClusterHealth.WARN:
            return 'bg-warn-600';
        case ClusterHealth.OK:
            return 'bg-success-600';
        case ClusterHealth.ERR:
        default:
            return 'bg-error-500';
    }
}

export function clusterHealthToTextColor(status: ClusterHealth | undefined): string {
    switch (status) {
        case ClusterHealth.UNSPECIFIED:
        case ClusterHealth.OFFLINE:
        case ClusterHealth.WARN:
            return 'text-warn-600';
        case ClusterHealth.OK:
            return 'text-success-600';
        case ClusterHealth.ERR:
        default:
            return 'text-error-500';
    }
}
