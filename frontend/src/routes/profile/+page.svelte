<script lang="ts">
    import { onMount } from 'svelte';
    import { getCurrentUser, getTweetsByUserId, getCurrentUserId, updateTweet, deleteTweet } from '$lib/api';
    import { getInitials, getAvatarColor } from '$lib/avatar';
    import type { User, Tweet } from '$lib/api';

    let user = $state<User | null>(null);
    let tweets = $state<Tweet[]>([]);
    let loading = $state(true);
    let error = $state('');

    // For editing
    let editingId = $state<number | null>(null);
    let editContent = $state('');
    let saving = $state(false);

    onMount(async () => {
        await loadProfile();
    });

    async function loadProfile() {
        try {
            user = await getCurrentUser();
            tweets = await getTweetsByUserId(user.id);
        } catch (e) {
            error = 'Failed to load profile';
        } finally {
            loading = false;
        }
    }

    function startEdit(tweet: Tweet) {
        editingId = tweet.id;
        editContent = tweet.content;
    }

    function cancelEdit() {
        editingId = null;
        editContent = '';
    }

    async function saveEdit(userId: number, tweetId: number) {
        if (!editContent.trim()) return;

        saving = true;
        try {
            await updateTweet(userId, tweetId, editContent);
            await loadProfile();
            editingId = null;
            editContent = '';
        } catch (e) {
            error = 'Failed to update tweet';
        } finally {
            saving = false;
        }
    }

    async function handleDelete(userId: number, tweetId: number) {
        if (!confirm('Are you sure you want to delete this tweet?')) return;

        try {
            await deleteTweet(userId, tweetId);
            await loadProfile();
        } catch (e) {
            error = 'Failed to delete tweet';
        }
    }

    function formatDate(dateString: string): string {
        return new Date(dateString).toLocaleDateString('en-US', {
            month: 'short',
            day: 'numeric',
            year: 'numeric'
        });
    }

    function gotoHome() {
        window.location.href = '/';
    }
</script>

<svelte:head>
    <title>Twitter Clone - Profile</title>
</svelte:head>

{#if loading}
    <div class="loading">Loading profile...</div>
{:else if user}
    <div class="profile">
        <div class="profile-header">
            <span class="avatar" style="background: {getAvatarColor(user.name)}">
                {getInitials(user.name)}
            </span>
            <h1>{user.name}</h1>
            <p class="email">{user.email}</p>
            <p class="age">Age: {user.age}</p>
            <p class="joined">Joined: {formatDate(user.created_at)}</p>
        </div>

        {#if error}
            <p class="error">{error}</p>
        {/if}

        <div class="tweets-section">
            <h2>Your Tweets</h2>

            {#if tweets.length === 0}
                <div class="empty">
                    <p>You haven't posted any tweets yet.</p>
                    <button onclick={gotoHome}>Post your first tweet</button>
                </div>
            {:else}
                <div class="tweets">
                    {#each tweets as tweet (tweet.id)}
                        <div class="tweet-card">
                            {#if editingId === tweet.id}
                                <div class="edit-form">
                                    <textarea
                                        bind:value={editContent}
                                        maxlength={280}
                                    ></textarea>
                                    <div class="edit-actions">
                                        <button
                                            class="save"
                                            onclick={() => saveEdit(user!.id, tweet.id)}
                                            disabled={saving}
                                        >
                                            {saving ? 'Saving...' : 'Save'}
                                        </button>
                                        <button
                                            class="cancel"
                                            onclick={cancelEdit}
                                            disabled={saving}
                                        >
                                            Cancel
                                        </button>
                                    </div>
                                </div>
                            {:else}
                                <p class="tweet-content">{tweet.content}</p>
                                <p class="tweet-time">{new Date(tweet.created_at).toLocaleString()}</p>
                                <div class="tweet-actions">
                                    <button onclick={() => startEdit(tweet)}>Edit</button>
                                    <button class="delete" onclick={() => handleDelete(user!.id, tweet.id)}>Delete</button>
                                </div>
                            {/if}
                        </div>
                    {/each}
                </div>
            {/if}
        </div>
    </div>
{:else}
    <div class="error">
        <p>Please <a href="/login">login</a> to view your profile.</p>
    </div>
{/if}

<style>
    .loading {
        text-align: center;
        padding: 2rem;
        color: #657786;
    }

    .profile {
        background: white;
        border-radius: 16px;
        overflow: hidden;
        box-shadow: 0 8px 30px rgba(29, 39, 59, 0.08);
    }

    .profile-header {
        padding: 2.5rem 2rem;
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        color: white;
        text-align: center;
    }

    .avatar {
        width: 72px;
        height: 72px;
        border-radius: 50%;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        font-size: 1.6rem;
        font-weight: 700;
        border: 3px solid rgba(255, 255, 255, 0.6);
        margin-bottom: 0.75rem;
    }

    .profile-header h1 {
        margin: 0 0 0.5rem 0;
    }

    .email {
        margin: 0.25rem 0;
        opacity: 0.9;
    }

    .age {
        margin: 0.25rem 0;
        opacity: 0.9;
    }

    .joined {
        margin: 0.25rem 0;
        opacity: 0.8;
        font-size: 0.9rem;
    }

    .tweets-section {
        padding: 1rem;
    }

    .tweets-section h2 {
        margin: 0 0 1rem 0;
        padding-bottom: 0.5rem;
        border-bottom: 1px solid #e1e8ed;
    }

    .tweet-card {
        padding: 1rem;
        border-bottom: 1px solid #e1e8ed;
    }

    .tweet-content {
        margin: 0 0 0.5rem 0;
        white-space: pre-wrap;
    }

    .tweet-time {
        color: #657786;
        font-size: 0.875rem;
        margin: 0 0 0.5rem 0;
    }

    .tweet-actions {
        display: flex;
        gap: 0.5rem;
        margin-top: 0.5rem;
    }

    .tweet-actions button {
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        color: white;
        border: none;
        padding: 0.3rem 0.9rem;
        border-radius: 999px;
        cursor: pointer;
        font-size: 0.85rem;
        font-weight: 600;
    }

    .tweet-actions button.delete {
        background: #e0245e;
    }

    .edit-form textarea {
        width: 100%;
        min-height: 80px;
        padding: 0.75rem;
        border: 1px solid #e1e8ed;
        border-radius: 10px;
        font-size: 1rem;
        resize: vertical;
        font-family: inherit;
    }

    .edit-actions {
        display: flex;
        gap: 0.5rem;
        margin-top: 0.5rem;
    }

    .edit-actions button {
        border: none;
        padding: 0.25rem 0.75rem;
        border-radius: 15px;
        cursor: pointer;
        font-size: 0.875rem;
    }

    .edit-actions button.save {
        background: #1da1f2;
        color: white;
    }

    .edit-actions button.cancel {
        background: #657786;
        color: white;
    }

    .empty {
        text-align: center;
        padding: 2rem;
        color: #657786;
    }

    .empty button {
        background: #1da1f2;
        color: white;
        border: none;
        padding: 0.5rem 1rem;
        border-radius: 20px;
        cursor: pointer;
        margin-top: 1rem;
    }

    .error {
        text-align: center;
        padding: 2rem;
        color: #e0245e;
    }

    .error a {
        color: #1da1f2;
    }
</style>
