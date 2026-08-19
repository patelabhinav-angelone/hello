import { writable } from 'svelte/store';
import type { User } from './api';

export const currentUser = writable<User | null>(null);
export const isLoggedIn = writable<boolean>(false);

export function setCurrentUser(user: User | null) {
	currentUser.set(user);
	isLoggedIn.set(user !== null);
}
