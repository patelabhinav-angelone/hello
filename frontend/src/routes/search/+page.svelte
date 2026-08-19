<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { searchUsers, getTweetsByUserId } from '$lib/api';
    import { getInitials, getAvatarColor } from '$lib/avatar';
    import type { User, Tweet } from '$lib/api';

    let query = $state('');
    let results = $state<User[]>([]);
    let loading = $state(false);
    let error = $state('');
    let searched = $state(false);

    let expandedUserId = $state<number | null>(null);
    let userTweets = $state<Record<number, Tweet[]>>({});
    let tweetsLoading = $state<number | null>(null);

    onMount(() => {
        const q = $page.url.searchParams.get('q');
        if (q) {
            query = q;
            runSearch();
        }
    });

    async function runSearch() {
        if (!query.trim()) return;
        loading = true;
        error = '';
        searched = true;
        expandedUserId = null;

        try {
            results = await searchUsers(query.trim());
        } catch (e) {
            error = 'Failed to search users';
            results = [];
        } finally {
            loading = false;
        }
    }

    function handleSubmit(e: Event) {
        e.preventDefault();
        const url = new URL(window.location.href);
        url.searchParams.set('q', query.trim());
        window.history.pushState({}, '', url);
        runSearch();
    }

    async function toggleUser(userId: number) {
        if (expandedUserId === userId) {
            expandedUserId = null;
            return;
        }

        expandedUserId = userId;

        if (!userTweets[userId]) {
            tweetsLoading = userId;
            try {
                userTweets[userId] = await getTweetsByUserId(userId);
            } catch (e) {
                userTweets[userId] = [];
            } finally {
                tweetsLoading = null;
            }
        }
    }

    function formatDate(dateString: string): string {
        return new Date(dateString).toLocaleDateString('en-US', {
            month: 'short',
            day: 'numeric',
            year: 'numeric'
        });
    }
</script>

<svelte:head>
    <title>Search - TweetApp</title>
</svelte:head>

<div class="search-page">
    <h1>Find People</h1>

    <form class="search-input-form" onsubmit={handleSubmit}>
        <input
            type="text"
            bind:value={query}
            placeholder="Search by name or email..."
        />
        <button type="submit" disabled={!query.trim()}>Search</button>
    </form>

    {#if error}
        <p class="error">{error}</p>
    {/if}

    {#if loading}
        <div class="loading">Searching...</div>
    {:else if searched && results.length === 0}
        <div class="empty">No users found matching "{query}"</div>
    {:else if results.length > 0}
        <div class="results">
            {#each results as user (user.id)}
                <div class="user-card" class:expanded={expandedUserId === user.id}>
                    <button class="user-row" onclick={() => toggleUser(user.id)}>
                        <span class="avatar" style="background: {getAvatarColor(user.name)}">
                            {getInitials(user.name)}
                        </span>
                        <span class="user-info">
                            <span class="user-name">{user.name}</span>
                            <span class="user-email">{user.email}</span>
                        </span>
                        <span class="chevron">{expandedUserId === user.id ? '▲' : '▼'}</span>
                    </button>

                    {#if expandedUserId === user.id}
                        <div class="tweets-panel">
                            {#if tweetsLoading === user.id}
                                <p class="panel-loading">Loading tweets...</p>
                            {:else if (userTweets[user.id] ?? []).length === 0}
                                <p class="panel-empty">No tweets yet.</p>
                            {:else}
                                {#each userTweets[user.id] as tweet (tweet.id)}
                                    <div class="tweet-item">
                                        <p class="tweet-content">{tweet.content}</p>
                                        <p class="tweet-time">{formatDate(tweet.created_at)}</p>
                                    </div>
                                {/each}
                            {/if}
                        </div>
                    {/if}
                </div>
            {/each}
        </div>
    {/if}
</div>

<style>
    .search-page {
        background: white;
        border-radius: 16px;
        overflow: hidden;
        box-shadow: 0 8px 30px rgba(29, 39, 59, 0.08);
    }

    h1 {
        padding: 1.25rem 1.5rem 0.5rem;
        margin: 0;
        font-size: 1.4rem;
    }

    .search-input-form {
        display: flex;
        gap: 0.5rem;
        padding: 1rem 1.5rem 1.25rem;
        border-bottom: 1px solid #eef2f6;
    }

    .search-input-form input {
        flex: 1;
        padding: 0.7rem 1rem;
        border: 1px solid #e1e8ed;
        border-radius: 999px;
        font-size: 1rem;
        outline: none;
        font-family: inherit;
        transition: border-color 0.15s ease;
    }

    .search-input-form input:focus {
        border-color: #1da1f2;
    }

    .search-input-form button {
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        color: white;
        border: none;
        padding: 0.7rem 1.5rem;
        border-radius: 999px;
        font-weight: 700;
        cursor: pointer;
        transition: transform 0.15s ease;
    }

    .search-input-form button:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .search-input-form button:not(:disabled):hover {
        transform: translateY(-1px);
    }

    .loading,
    .empty {
        text-align: center;
        padding: 2.5rem 1rem;
        color: #657786;
    }

    .error {
        color: #e0245e;
        text-align: center;
        padding: 1rem;
        margin: 0;
    }

    .results {
        display: flex;
        flex-direction: column;
    }

    .user-card {
        border-bottom: 1px solid #eef2f6;
    }

    .user-card.expanded {
        background: #f8fbfe;
    }

    .user-row {
        width: 100%;
        display: flex;
        align-items: center;
        gap: 0.9rem;
        padding: 1rem 1.5rem;
        background: none;
        border: none;
        cursor: pointer;
        text-align: left;
        font-family: inherit;
        transition: background 0.15s ease;
    }

    .user-row:hover {
        background: #f4f8fc;
    }

    .avatar {
        width: 42px;
        height: 42px;
        min-width: 42px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        color: white;
        font-weight: 700;
        font-size: 0.95rem;
    }

    .user-info {
        flex: 1;
        display: flex;
        flex-direction: column;
        gap: 0.1rem;
        min-width: 0;
    }

    .user-name {
        font-weight: 700;
        color: #14171a;
    }

    .user-email {
        font-size: 0.85rem;
        color: #657786;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }

    .chevron {
        color: #aab8c2;
        font-size: 0.75rem;
    }

    .tweets-panel {
        padding: 0 1.5rem 1.25rem 4.9rem;
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
    }

    .panel-loading,
    .panel-empty {
        color: #657786;
        margin: 0;
        font-size: 0.9rem;
    }

    .tweet-item {
        background: white;
        border: 1px solid #eef2f6;
        border-radius: 12px;
        padding: 0.75rem 1rem;
    }

    .tweet-content {
        margin: 0 0 0.35rem 0;
        white-space: pre-wrap;
    }

    .tweet-time {
        margin: 0;
        font-size: 0.78rem;
        color: #aab8c2;
    }
</style>
