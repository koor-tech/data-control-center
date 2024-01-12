import { ResourceStatus } from '~~/gen/ts/api/resources/k8s/v1/resources_pb';

export function resourceStatusToBGColor(status: ResourceStatus | undefined): string {
    switch (status) {
        case ResourceStatus.UNSPECIFIED:
        case ResourceStatus.UNKNOWN:
            return 'bg-warn-600';
        case ResourceStatus.PROGRESSING:
            return 'bg-info-600';
        case ResourceStatus.READY:
            return 'bg-success-600';
        case ResourceStatus.NOT_READY:
        default:
            return 'bg-error-500';
    }
}

export function resourceStatusToTextColor(status: ResourceStatus | undefined): string {
    switch (status) {
        case ResourceStatus.UNSPECIFIED:
        case ResourceStatus.UNKNOWN:
            return 'text-warn-600';
        case ResourceStatus.PROGRESSING:
            return 'text-info-600';
        case ResourceStatus.READY:
            return 'text-success-600';
        case ResourceStatus.NOT_READY:
        default:
            return 'text-error-500';
    }
}
