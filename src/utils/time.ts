import { Timestamp } from '@bufbuild/protobuf';
import { useTimeAgo as vueUseUseTimeAgo } from '@vueuse/core';

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
 * @param {Date | Timestamp | undefined} date - The timestamp to convert. If `undefined` the current date will be used.
 * @returns {ComputedRef<string>} A computed ref string indicating time elapsed (auto updates).
 */
export function useTimeAgo(date: Date | Timestamp | undefined): ComputedRef<string> {
    if (date === undefined) {
        date = new Date();
    }
    if (date instanceof Timestamp) {
        date = date.toDate();
    }

    return vueUseUseTimeAgo(date);
}
