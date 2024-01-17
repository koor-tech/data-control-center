import { Timestamp } from '@bufbuild/protobuf';
import { useTimeAgo as vueUseUseTimeAgo } from '@vueuse/core';

export function toDate(ts: Timestamp | undefined): Date {
    if (ts === undefined) {
        return new Date();
    }
    return ts.toDate();
}

export function toDateLocaleString(ts: Timestamp | undefined): string {
    if (typeof ts === 'undefined') {
        return '-';
    }

    return ts.toDate().toLocaleDateString()!;
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

const secondsPerMinute = 60;
const secondsPerHour = secondsPerMinute * 60;
const secondsPerDay = secondsPerHour * 24;
const secondsPerWeek = secondsPerDay * 7;
const secondsPerYear = secondsPerWeek * 52;

export function fromSecondsToFormattedDuration(seconds: number, options?: { seconds?: boolean; emptyText?: string }): string {
    const years = Math.floor(seconds / secondsPerYear);
    seconds -= years * secondsPerYear;
    const weeks = Math.floor(seconds / secondsPerWeek);
    seconds -= weeks * secondsPerWeek;
    const days = Math.floor(seconds / secondsPerDay);
    seconds -= days * secondsPerDay;
    const hours = Math.floor(seconds / secondsPerHour);
    seconds -= hours * secondsPerHour;
    const minutes = Math.floor(seconds / secondsPerMinute);
    seconds -= minutes * secondsPerMinute;

    const parts: String[] = [];
    if (years > 0) {
        parts.push(`${years} Year(s)`);
    }
    if (weeks > 0) {
        parts.push(`${weeks} Week(s)`);
    }
    if (days > 0) {
        parts.push(`${days} Day(s)`);
    }
    if (hours > 0) {
        parts.push(`${hours} Hour(s)`);
    }
    if (minutes > 0) {
        parts.push(`${minutes} Minute(s)`);
    }
    if ((!options || options.seconds) && seconds > 0) {
        parts.push(`${seconds} Second(s)`);
    }

    const text = parts.join(', ');
    return text.length > 0 ? text : 'Unknown';
}
