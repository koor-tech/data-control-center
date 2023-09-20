import { Timestamp } from '@bufbuild/protobuf';

export function toDate(ts: Timestamp | undefined): Date {
    if (ts === undefined) {
        return new Date();
    }
    return ts.toDate();
}

export function toDateLocaleString(ts: Timestamp | undefined): string {
    if (typeof ts === undefined) {
        return '-';
    }

    return ts?.toDate().toLocaleDateString()!;
}

export function fromString(time: string): Date {
    return new Date(Date.parse(time));
}

/**
 * Converts a timestamp to a human-readable string representing time elapsed.
 * @param {Timestamp} timestamp - The timestamp to convert.
 * @returns {string} A human-readable string indicating time elapsed.
 */
export function convertTimestampToAgoString(timestamp: Timestamp): string {
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
