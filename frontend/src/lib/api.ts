import { browser } from '$app/environment';

const API_BASE = 'http://localhost:8080/api';

export interface User {
	id: number;
	name: string;
	email: string;
	age: number;
	created_at: string;
	updated_at: string;
}

export interface Tweet {
	id: number;
	content: string;
	user_id: number;
	created_at: string;
	updated_at: string;
}

export interface AuthResponse {
	token: string;
	user: User;
}

function getToken(): string | null {
	if (browser) {
		return localStorage.getItem('token');
	}
	return null;
}

function getUserId(): string | null {
	if (browser) {
		return localStorage.getItem('user_id');
	}
	return null;
}

function setToken(token: string, userId: number): void {
	if (browser) {
		localStorage.setItem('token', token);
		localStorage.setItem('user_id', userId.toString());
	}
}

function clearToken(): void {
	if (browser) {
		localStorage.removeItem('token');
		localStorage.removeItem('user_id');
	}
}

function getAuthHeaders(): HeadersInit {
	const headers: HeadersInit = {
		'Content-Type': 'application/json'
	};
	const token = getToken();
	if (token) {
		headers['Authorization'] = `Bearer ${token}`;
	}
	return headers;
}

// Auth APIs
export async function login(email: string, password: string): Promise<AuthResponse> {
	const response = await fetch(`${API_BASE}/auth/login`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ email, password })
	});

	if (!response.ok) {
		const error = await response.json();
		throw new Error(error.error || 'Login failed');
	}

	const body = await response.json();
	const user: User = body.data;
	setToken(body.token, user.id);
	return { token: body.token, user };
}

export async function register(name: string, email: string, password: string, age: number): Promise<AuthResponse> {
	const response = await fetch(`${API_BASE}/auth/register`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ name, email, password, age })
	});

	if (!response.ok) {
		const error = await response.json();
		throw new Error(error.error || 'Registration failed');
	}

	// Register does not return a token, so log in right after to establish a session
	return login(email, password);
}

export function logout(): void {
	clearToken();
}

// User APIs
export async function getCurrentUser(): Promise<User> {
	const response = await fetch(`${API_BASE}/users/me`, {
		headers: getAuthHeaders()
	});

	if (!response.ok) {
		throw new Error('Failed to fetch user');
	}

	const data = await response.json();
	return data.data;
}

export async function getUserById(id: number): Promise<User> {
	const response = await fetch(`${API_BASE}/users/${id}`);

	if (!response.ok) {
		throw new Error('User not found');
	}

	const data = await response.json();
	return data.data;
}

export async function getAllUsers(): Promise<User[]> {
	const response = await fetch(`${API_BASE}/users`);

	if (!response.ok) {
		throw new Error('Failed to fetch users');
	}

	const data = await response.json();
	return data.data;
}

export async function searchUsers(query: string): Promise<User[]> {
	const response = await fetch(`${API_BASE}/users/search?q=${encodeURIComponent(query)}`);

	if (!response.ok) {
		throw new Error('Failed to search users');
	}

	const data = await response.json();
	return data.data;
}

// Tweet APIs
export async function getAllTweets(): Promise<Tweet[]> {
	const response = await fetch(`${API_BASE}/tweets`);

	if (!response.ok) {
		throw new Error('Failed to fetch tweets');
	}

	const data = await response.json();
	return data.data;
}

export async function getTweetsByUserId(userId: number): Promise<Tweet[]> {
	const response = await fetch(`${API_BASE}/tweets/user/${userId}`);

	if (!response.ok) {
		throw new Error('Failed to fetch tweets');
	}

	const data = await response.json();
	return data.data;
}

export async function createTweet(userId: number, content: string): Promise<Tweet> {
	const response = await fetch(`${API_BASE}/tweets/user/${userId}`, {
		method: 'POST',
		headers: getAuthHeaders(),
		body: JSON.stringify({ content })
	});

	if (!response.ok) {
		const error = await response.json();
		throw new Error(error.error || 'Failed to create tweet');
	}

	const data = await response.json();
	return data.data;
}

export async function updateTweet(userId: number, tweetId: number, content: string): Promise<Tweet> {
	const response = await fetch(`${API_BASE}/users/${userId}/tweets/${tweetId}`, {
		method: 'PUT',
		headers: getAuthHeaders(),
		body: JSON.stringify({ content })
	});

	if (!response.ok) {
		const error = await response.json();
		throw new Error(error.error || 'Failed to update tweet');
	}

	const data = await response.json();
	return data.data;
}

export async function deleteTweet(userId: number, tweetId: number): Promise<void> {
	const response = await fetch(`${API_BASE}/users/${userId}/tweets/${tweetId}`, {
		method: 'DELETE',
		headers: getAuthHeaders()
	});

	if (!response.ok) {
		const error = await response.json();
		throw new Error(error.error || 'Failed to delete tweet');
	}
}

export function isAuthenticated(): boolean {
	return getToken() !== null;
}

export function getCurrentUserId(): number | null {
	const userId = getUserId();
	return userId ? parseInt(userId) : null;
}
