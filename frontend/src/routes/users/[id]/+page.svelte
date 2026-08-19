<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { getUserById, getTweetsByUserId } from '$lib/api';
    import { getInitials, getAvatarColor } from '$lib/avatar';
    import type { User, Tweet } from '$lib/api';

    let user = $state<User | null>(null);
    let tweets = $state<Tweet[]>([]);
    let loading = $state(true);
    let error = $state('');

    $effect(() => {
        loadUser();
    });

    async function loadUser() {
        const userId = parseInt($page.params.id ?? '');

        if (isNaN(userId)) {
            error = 'Invalid user ID';
            loading = false;
            return;
        }

        try {
            user = await getUserById(userId);
            tweets = await getTweetsByUserId(userId);
        } catch (e) {
            error = 'User not found';
        } finally {
            loading = false;
        }
    }

    function formatDate(dateString: string): string {
        return new Date(dateString).toLocaleDateString('en-US', {
            month: 'short',
            day: 'numeric',
            year: 'numeric'
        });
    }

    function goToHome() {
        window.location.href = '/';
    }
</script>

<svelte:head>
    <title>Twitter Clone - User</title>
</svelte:head>

{#if loading}
    <div class="loading">Loading user profile...</div>
{:else if error}
    <div class="error">
        <p>{error}</p>
        <button onclick={goToHome}>Go Back Home</button>
    </div>
{:else if user}
    <div class="user-profile">
        <div class="profile-header">
            <span class="avatar" style="background: {getAvatarColor(user.name)}">
                {getInitials(user.name)}
            </span>
            <h1>{user.name}</h1>
            <p class="email">{user.email}</p>
            <p class="age">Age: {user.age}</p>
            <p class="joined">Joined: {formatDate(user.created_at)}</p>
        </div>

        <div class="tweets-section">
            <h2>Tweets ({tweets.length})</h2>

            {#if tweets.length === 0}
                <div class="empty">
                    <p>This user hasn't posted any tweets yet.</p>
                </div>
            {:else}
                <div class="tweets">
                    {#each tweets as tweet (tweet.id)}
                        <div class="tweet-card">
                            <p class="tweet-content">{tweet.content}</p>
                            <p class="tweet-time">{new Date(tweet.created_at).toLocaleString()}</p>
                        </div>
                    {/each}
                </div>
            {/if}
        </div>
    </div>
{/if}

<style>
    .loading {
        text-align: center;
        padding: 2rem;
        color: #657786;
    }

    .user-profile {
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
        margin: 0;
    }

    .empty {
        text-align: center;
        padding: 2rem;
        color: #657786;
    }

    .error {
        text-align: center;
        padding: 2rem;
    }

    .error button {
        background: #1da1f2;
        color: white;
        border: none;
        padding: 0.5rem 1rem;
        border-radius: 20px;
        cursor: pointer;
        margin-top: 1rem;
    }
</style>
