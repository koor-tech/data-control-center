import * as resources_timestamp_timestamp from '~~/gen/ts/api/resources/timestamp/timestamp_pb';


export function toDate(ts: resources_timestamp_timestamp.Timestamp | undefined): Date {
    if (ts === undefined || ts?.timestamp === undefined) {
        return new Date();
    }
    return ts?.timestamp.toDate();
}

export function toDateLocaleString(ts: resources_timestamp_timestamp.Timestamp | undefined): string {
    if (typeof ts === undefined) {
        return '-';
    }

    return ts?.timestamp?.toDate().toLocaleDateString()!;
}

export function fromString(time: string): Date {
    return new Date(Date.parse(time));
}
