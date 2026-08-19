<script lang="ts">
    import { onMount } from 'svelte';
    import { getTweetsByUserId, createTweet, getCurrentUserId } from '$lib/api';
    import type { Tweet } from '$lib/api';

    let tweets = $state<Tweet[]>([]);
    let content = $state('');
    let loading = $state(true);
    let posting = $state(false);
    let error = $state('');
    let showCompose = $state(false);

    onMount(async () => {
        await loadTweets();
    });

    function focusOnMount(node: HTMLElement) {
        node.focus();
    }

    async function loadTweets() {
        const userId = getCurrentUserId();
        if (!userId) {
            window.location.href = '/login';
            return;
        }

        try {
            tweets = await getTweetsByUserId(userId);
        } catch (e) {
            error = 'Failed to load tweets';
        } finally {
            loading = false;
        }
    }

    function openCompose() {
        showCompose = true;
        error = '';
    }

    function closeCompose() {
        showCompose = false;
        content = '';
        error = '';
    }

    async function handlePost(e: Event) {
        e.preventDefault();
        if (!content.trim()) return;

        posting = true;
        error = '';

        try {
            const userId = getCurrentUserId();
            if (!userId) {
                window.location.href = '/login';
                return;
            }

            await createTweet(userId, content);
            content = '';
            showCompose = false;
            await loadTweets();
        } catch (e) {
            error = 'Failed to post tweet';
        } finally {
            posting = false;
        }
    }

    function formatDate(dateString: string): string {
        return new Date(dateString).toLocaleDateString('en-US', {
            month: 'short',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
        });
    }
</script>

<svelte:head>
    <title>Twitter Clone - Home</title>
</svelte:head>

<div class="toolbar">
    <button class="tweet-button" onclick={openCompose}>Tweet</button>
</div>

{#if showCompose}
    <div
        class="modal-overlay"
        role="presentation"
        onclick={(e) => e.target === e.currentTarget && closeCompose()}
        onkeydown={(e) => e.key === 'Escape' && closeCompose()}
    >
        <div class="modal">
            <form onsubmit={handlePost}>
                <textarea
                    bind:value={content}
                    placeholder="What's happening?"
                    maxlength={280}
                    use:focusOnMount
                ></textarea>
                {#if error}
                    <p class="error">{error}</p>
                {/if}
                <div class="actions">
                    <span class="char-count">{content.length}/280</span>
                    <div class="buttons">
                        <button type="button" class="cancel" onclick={closeCompose}>Cancel</button>
                        <button type="submit" disabled={posting || !content.trim()}>
                            {posting ? 'Posting...' : 'Tweet'}
                        </button>
                    </div>
                </div>
            </form>
        </div>
    </div>
{/if}

<div class="feed">
    <h1>My Tweets</h1>

    {#if loading}
        <div class="loading">Loading tweets...</div>
    {:else if tweets.length === 0}
        <div class="empty">No tweets yet. Click "Tweet" to post your first one!</div>
    {:else}
        <div class="tweets">
            {#each tweets as tweet (tweet.id)}
                <div class="tweet-card">
                    <div class="tweet-header">
                        <span class="time">{formatDate(tweet.created_at)}</span>
                    </div>
                    <p class="tweet-content">{tweet.content}</p>
                </div>
            {/each}
        </div>
    {/if}
</div>

<style>
    .toolbar {
        display: flex;
        justify-content: flex-end;
        margin-bottom: 1rem;
    }

    .tweet-button {
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        color: white;
        border: none;
        padding: 0.65rem 1.75rem;
        border-radius: 999px;
        font-size: 1rem;
        font-weight: 700;
        cursor: pointer;
        box-shadow: 0 6px 16px rgba(29, 161, 242, 0.3);
        transition: transform 0.15s ease, box-shadow 0.15s ease;
    }

    .tweet-button:hover {
        transform: translateY(-1px);
        box-shadow: 0 8px 20px rgba(29, 161, 242, 0.4);
    }

    .modal-overlay {
        position: fixed;
        inset: 0;
        background: rgba(20, 23, 26, 0.55);
        backdrop-filter: blur(2px);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 100;
    }

    .modal {
        background: white;
        border-radius: 16px;
        padding: 1.5rem;
        width: 90%;
        max-width: 500px;
        box-shadow: 0 20px 50px rgba(0, 0, 0, 0.25);
    }

    .modal textarea {
        width: 100%;
        min-height: 100px;
        padding: 0.75rem;
        border: 1px solid #e1e8ed;
        border-radius: 10px;
        font-size: 1rem;
        resize: vertical;
        font-family: inherit;
        box-sizing: border-box;
    }

    .actions {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-top: 0.75rem;
    }

    .char-count {
        color: #657786;
        font-size: 0.875rem;
    }

    .buttons {
        display: flex;
        gap: 0.5rem;
    }

    .buttons button {
        border: none;
        padding: 0.5rem 1.25rem;
        border-radius: 20px;
        font-size: 1rem;
        cursor: pointer;
        font-weight: bold;
    }

    .buttons button[type='submit'] {
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        color: white;
    }

    .buttons button[type='submit']:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }

    .buttons button.cancel {
        background-color: #e1e8ed;
        color: #333;
    }

    .feed {
        background: white;
        border-radius: 16px;
        overflow: hidden;
        box-shadow: 0 8px 30px rgba(29, 39, 59, 0.08);
    }

    h1 {
        padding: 1.25rem 1.5rem;
        margin: 0;
        font-size: 1.4rem;
        border-bottom: 1px solid #eef2f6;
    }

    .tweets {
        display: flex;
        flex-direction: column;
    }

    .tweet-card {
        padding: 1rem 1.5rem;
        border-bottom: 1px solid #eef2f6;
        transition: background 0.15s ease;
    }

    .tweet-card:hover {
        background: #f8fbfe;
    }

    .tweet-header {
        display: flex;
        gap: 0.5rem;
        margin-bottom: 0.5rem;
    }

    .time {
        color: #657786;
        font-size: 0.875rem;
    }

    .tweet-content {
        margin: 0;
        white-space: pre-wrap;
    }

    .loading,
    .empty {
        text-align: center;
        padding: 2rem;
        color: #657786;
    }

    .error {
        color: red;
        text-align: center;
        padding: 0.5rem 0 0 0;
        margin: 0;
    }
</style>
