import { ConnectError } from '@connectrpc/connect';
import {
    type GetResourcesResponse,
    type SaveResourceRequest,
    type SaveResourceResponse,
} from '~~/gen/ts/api/services/k8sresources/v1/editor_pb';

export async function getResources(): Promise<GetResourcesResponse> {
    const { $grpc } = useNuxtApp();
    try {
        return await $grpc.getK8sResourcesClient().getResources({});
    } catch (e) {
        if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
        throw e;
    }
}

export async function saveResource(
    name: string | undefined,
    namespace: string | undefined,
    kind: string | undefined,
    object: string | undefined,
    content: string,
): Promise<SaveResourceResponse> {
    const { $grpc } = useNuxtApp();
    try {
        const updateResource = {
            resource: { name, namespace, content, kind, object },
        } as SaveResourceRequest;
        return await $grpc.getK8sResourcesClient().saveResource(updateResource);
    } catch (e) {
        if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
        throw e;
    }
}
