export function toTitleCase(s: string): string {
    return s.replace(/\w\S*/g, (w) => w.charAt(0).toUpperCase() + w.slice(1).toLowerCase());
}

export function lowercaseFirstLetter(s: string): string {
    return s.charAt(0).toLowerCase() + s.slice(1);
}

export function camelCaseToTitleCase(s: string): string {
    const result = s.replace(/([A-Z])/g, ' $1');
    return result.charAt(0).toUpperCase() + result.slice(1);
}

// Taken from https://stackoverflow.com/a/18650828
export function formatBytes(bytes: number, decimals = 2) {
    if (!+bytes) return '0 Bytes';

    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'];

    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
}

export function formatBytesBigInt(bytes: bigint, decimals = 2) {
    return formatBytes(parseInt(bytes.toString()), decimals);
}
