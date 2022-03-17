import { writable } from 'svelte/store';
import { emptyDialog } from '$lib/modal.js'

export const dialog = writable(emptyDialog);
