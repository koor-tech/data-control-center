import { Timestamp } from '@bufbuild/protobuf';
import { ClusterHealthStats, DisplayStatsData, meta } from '~/composables/stats/types';
import { camelCaseToTitleCase, formatBytes } from '~/utils/strings';
import { ClusterStats, PGs } from '~~/gen/ts/api/resources/stats/v1/stats_pb';

export class TransformStats {
    constructor(private clusterStats: ClusterStats) {}

    /**
     * Transforms descriptions for display.
     * @param {string} title - The title of the data.
     * @param {any} data - The data to transform.
     * @returns {DisplayStatsData[]} An array of transformed data.
     */
    private transformDescriptions(title: string, data: any): DisplayStatsData[] {
        return Object.keys(data).map((serviceName) => {
            const serviceData = data[serviceName];
            const serviceMeta = meta[serviceName] ?? meta['default'];

            if (serviceName === 'volumes') {
                return {
                    title: serviceMeta.title || serviceName,
                    icon: markRaw(serviceMeta.icon),
                    color: serviceMeta.color,
                    description: [{ title: serviceMeta.title || serviceName, description: serviceData }],
                };
            }

            const description = Object.entries(serviceData)
                .map(([key, value]) => {
                    if (value instanceof Timestamp) {
                        return `${camelCaseToTitleCase(key)}: ${useTimeAgo(value).value}`;
                    }
                    if (value instanceof PGs) {
                        return `Active + Clean: ${value.activeClean}`;
                    }
                    return `${camelCaseToTitleCase(key)}: ${value}`;
                })
                .join(', ');

            if (serviceName === 'usage') {
                return {
                    title: serviceMeta.title || serviceName,
                    icon: markRaw(serviceMeta.icon),
                    color: serviceMeta.color,
                    description: description.split(', ').map((entry) => {
                        const [key, value] = entry.split(': ');
                        return { title: key, description: formatBytes(parseInt(value)) };
                    }),
                };
            }

            return {
                title: serviceMeta.title || serviceName,
                icon: markRaw(serviceMeta.icon),
                color: serviceMeta.color,
                description: description.split(', ').map((entry) => {
                    const [key, value] = entry.split(': ');
                    return { title: key, description: value };
                }),
            };
        });
    }

    /**
     * Transforms cluster statistics for display.
     * @returns {ClusterHealthStats} Transformed cluster health statistics.
     */
    public display(): ClusterHealthStats {
        const transformedArray: DisplayStatsData[] = [];

        transformedArray.push({
            title: 'Alerts',
            icon: meta['alerts'].icon,
            color: meta['alerts']?.color,

            description: this.clusterStats.crashes.map((daemon) => ({
                description: daemon.description,
            })),
        });

        transformedArray.push(...this.transformDescriptions('services', this.clusterStats.services));
        transformedArray.push(...this.transformDescriptions('data', this.clusterStats.data));

        if (this.clusterStats.iops) {
            transformedArray.push({
                title: 'Input / Output',
                icon: markRaw(meta['iops'].icon),
                color: meta['iops']?.color,
                description: [
                    {
                        title: 'Read',
                        description: this.clusterStats.iops.clientRead,
                    },
                    {
                        title: 'Read Ops',
                        description: this.clusterStats.iops.clientReadOps,
                    },
                    {
                        title: 'Write',
                        description: this.clusterStats.iops.clientWrite,
                    },
                    {
                        title: 'Write Ops',
                        description: this.clusterStats.iops.clientWriteOps,
                    },
                ],
            });
        }

        return {
            id: this.clusterStats.id,
            health: this.clusterStats.status,
            stats: transformedArray,
        };
    }
}
