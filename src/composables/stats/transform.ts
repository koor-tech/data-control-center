import { Timestamp } from '@bufbuild/protobuf';
import { ClusterHealthStats, DisplayStatsData, meta } from '~/composables/stats/types';
import { ClusterStats, PGs } from '~~/gen/ts/api/resources/stats/stats_pb';

export class TransformStats {
    constructor(private clusterStats: ClusterStats) {}

    /**
     * Converts a timestamp to a human-readable string representing time elapsed.
     * @param {Timestamp} timestamp - The timestamp to convert.
     * @returns {string} A human-readable string indicating time elapsed.
     */
    private convertTimestampToAgoString(timestamp: Timestamp): string {
        const currentTimestamp = BigInt(Math.floor(Date.now() / 1000));
        const seconds = timestamp.seconds;
        const difference = currentTimestamp - seconds;

        if (difference < 60n) {
            return `${difference.toString()} seconds ago`;
        } else if (difference < 3600n) {
            const minutes = difference / 60n;
            return `${minutes.toString()} minutes ago`;
        } else if (difference < 86400n) {
            const hours = difference / 3600n;
            return `${hours.toString()} hours ago`;
        } else {
            const days = difference / 86400n;
            return `${days.toString()} days ago`;
        }
    }

    /**
     * Transforms descriptions for display.
     * @param {string} title - The title of the data.
     * @param {any} data - The data to transform.
     * @returns {DisplayStatsData[]} An array of transformed data.
     */
    private transformDescriptions(title: string, data: any): DisplayStatsData[] {
        return Object.keys(data).map((serviceName) => {
            const serviceData = data[serviceName];

            if (serviceName === 'volumes') {
                return {
                    title: meta[serviceName]?.title || serviceName,
                    icon: meta[serviceName]?.icon,
                    color: meta[serviceName]?.color,
                    description: [{ title: meta[serviceName]?.title || serviceName, description: serviceData }],
                };
            }

            const description = Object.entries(serviceData)
                .map(([key, value]) => {
                    if (value instanceof Timestamp) {
                        return `${key}: ${this.convertTimestampToAgoString(value)}`;
                    }
                    if (value instanceof PGs) {
                        return `active+clean: ${value.activeClean}`;
                    }
                    return `${key}: ${value}`;
                })
                .join(', ');

            return {
                title: meta[serviceName]?.title || serviceName,
                icon: meta[serviceName]?.icon,
                color: meta[serviceName]?.color,
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

            description: this.clusterStats.crashes.map((daemon: any) => daemon.description),
        });

        transformedArray.push(...this.transformDescriptions('services', this.clusterStats.services));
        transformedArray.push(...this.transformDescriptions('data', this.clusterStats.data));

        if (this.clusterStats.iops) {
            transformedArray.push({
                title: 'Input / Output',
                icon: meta['io'].icon,
                color: meta['io']?.color,
                description: [
                    {
                        title: 'Client Read',
                        description: this.clusterStats.iops.clientRead,
                    },
                    {
                        title: '',
                        description: this.clusterStats.iops.clientReadOps,
                    },
                    {
                        title: '',
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
