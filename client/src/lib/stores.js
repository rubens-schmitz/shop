import { writable } from 'svelte/store';

export const modalOpen = writable(false);
export const modalContent = writable('');
