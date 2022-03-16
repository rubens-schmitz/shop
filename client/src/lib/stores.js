import { writable } from 'svelte/store';

/**
 * @type {Dialog}
 */
const emptyDialog = { task: '', body: '', qrcode: '', reload: false };

export const dialog = writable(emptyDialog);
