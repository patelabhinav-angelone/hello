<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { browser } from '$app/environment';
    import { isLoggedIn, currentUser } from '$lib/stores';
    import { isAuthenticated, getCurrentUser, logout } from '$lib/api';
    import { getInitials, getAvatarColor } from '$lib/avatar';

    let { children } = $props();

    let showNav = $state(false);
    let mounted = $state(false);
    let searchQuery = $state('');

    const publicRoutes = ['/login', '/register'];

    onMount(async () => {
        if (browser && isAuthenticated()) {
            try {
                const user = await getCurrentUser();
                currentUser.set(user);
                isLoggedIn.set(true);
            } catch (e) {
                console.error('Failed to fetch user', e);
            }
        }

        if (browser && !isAuthenticated() && !publicRoutes.includes($page.url.pathname)) {
            window.location.href = '/login';
            return;
        }

        mounted = true;
    });

    $effect(() => {
        if (mounted && browser) {
            showNav = $isLoggedIn;
        }
    });

    function handleLogout() {
        logout();
        isLoggedIn.set(false);
        currentUser.set(null);
        window.location.href = '/login';
    }

    function handleSearch(e: Event) {
        e.preventDefault();
        if (!searchQuery.trim()) return;
        window.location.href = `/search?q=${encodeURIComponent(searchQuery.trim())}`;
    }
</script>

<main>
    {#if showNav}
        <nav class="navbar">
            <a href="/" class="logo">
                <span class="logo-icon">🐦</span>
                TweetApp
            </a>

            <form class="search-form" onsubmit={handleSearch}>
                <span class="search-icon">🔍</span>
                <input
                    type="text"
                    placeholder="Search people by name or email..."
                    bind:value={searchQuery}
                />
            </form>

            <div class="nav-links">
                <a href="/">Home</a>
                <a href="/profile">Profile</a>
                {#if $currentUser}
                    <span class="avatar" style="background: {getAvatarColor($currentUser.name)}">
                        {getInitials($currentUser.name)}
                    </span>
                {/if}
                <button onclick={handleLogout}>Logout</button>
            </div>
        </nav>
    {/if}

    <div class="container">
        {@render children()}
    </div>
</main>

<style>
    :global(body) {
        margin: 0;
        font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
        background: linear-gradient(160deg, #eef4fb 0%, #f5f7fa 45%, #eef2f8 100%);
        min-height: 100vh;
    }

    :global(*) {
        box-sizing: border-box;
    }

    .navbar {
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        color: white;
        padding: 0.85rem 2rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 1.5rem;
        box-shadow: 0 4px 20px rgba(29, 161, 242, 0.25);
        position: sticky;
        top: 0;
        z-index: 50;
    }

    .logo {
        font-size: 1.4rem;
        font-weight: 800;
        text-decoration: none;
        color: white;
        display: flex;
        align-items: center;
        gap: 0.4rem;
        white-space: nowrap;
        letter-spacing: -0.02em;
    }

    .logo-icon {
        font-size: 1.3rem;
    }

    .search-form {
        flex: 1;
        max-width: 380px;
        display: flex;
        align-items: center;
        background: rgba(255, 255, 255, 0.18);
        border: 1px solid rgba(255, 255, 255, 0.3);
        border-radius: 999px;
        padding: 0.5rem 1rem;
        gap: 0.5rem;
        transition: background 0.2s ease;
    }

    .search-form:hover,
    .search-form:focus-within {
        background: rgba(255, 255, 255, 0.28);
    }

    .search-icon {
        font-size: 0.9rem;
        opacity: 0.9;
    }

    .search-form input {
        flex: 1;
        border: none;
        background: transparent;
        color: white;
        font-size: 0.9rem;
        outline: none;
        font-family: inherit;
    }

    .search-form input::placeholder {
        color: rgba(255, 255, 255, 0.75);
    }

    .nav-links {
        display: flex;
        gap: 1.1rem;
        align-items: center;
        white-space: nowrap;
    }

    .nav-links a {
        color: white;
        text-decoration: none;
        font-weight: 500;
        opacity: 0.95;
        transition: opacity 0.15s ease;
    }

    .nav-links a:hover {
        opacity: 1;
        text-decoration: underline;
    }

    .avatar {
        width: 32px;
        height: 32px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        color: white;
        font-weight: 700;
        font-size: 0.8rem;
        border: 2px solid rgba(255, 255, 255, 0.6);
    }

    .nav-links button {
        background-color: white;
        color: #1da1f2;
        border: none;
        padding: 0.5rem 1.1rem;
        border-radius: 999px;
        cursor: pointer;
        font-weight: 700;
        transition: transform 0.15s ease, box-shadow 0.15s ease;
    }

    .nav-links button:hover {
        transform: translateY(-1px);
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15);
    }

    .container {
        max-width: 640px;
        margin: 0 auto;
        padding: 2rem 1rem;
    }

    @media (max-width: 700px) {
        .search-form {
            display: none;
        }
    }
</style>
