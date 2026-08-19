<script lang="ts">
    import { register } from '$lib/api';
    import { setCurrentUser } from '$lib/stores';

    let name = $state('');
    let email = $state('');
    let password = $state('');
    let age = $state(18);
    let error = $state('');
    let loading = $state(false);

    async function handleRegister(e: Event) {
        e.preventDefault();
        loading = true;
        error = '';

        try {
            const response = await register(name, email, password, age);
            setCurrentUser(response.user);
            window.location.href = '/';
        } catch (e: any) {
            if (e instanceof Error) {
                error = e.message;
            } else {
                error = 'Registration failed';
            }
        } finally {
            loading = false;
        }
    }
</script>

<div class="auth-container">
    <div class="brand">
        <span class="brand-icon">🐦</span>
        <h1>Join TweetApp</h1>
        <p class="subtitle">Create an account to start tweeting</p>
    </div>

    <form onsubmit={handleRegister}>
        <div class="form-group">
            <label for="name">Name</label>
            <input
                type="text"
                id="name"
                bind:value={name}
                placeholder="Enter your name"
                required
            />
        </div>

        <div class="form-group">
            <label for="email">Email</label>
            <input
                type="email"
                id="email"
                bind:value={email}
                placeholder="Enter your email"
                required
            />
        </div>

        <div class="form-group">
            <label for="password">Password</label>
            <input
                type="password"
                id="password"
                bind:value={password}
                placeholder="Create a password"
                required
            />
        </div>

        <div class="form-group">
            <label for="age">Age</label>
            <input
                type="number"
                id="age"
                bind:value={age}
                placeholder="Enter your age"
                min="0"
                max="150"
                required
            />
        </div>

        {#if error}
            <p class="error">{error}</p>
        {/if}

        <button type="submit" disabled={loading}>
            {loading ? 'Registering...' : 'Register'}
        </button>
    </form>

    <p class="link">
        Already have an account? <a href="/login">Login</a>
    </p>
</div>

<style>
    .auth-container {
        max-width: 400px;
        margin: 60px auto;
        padding: 2.5rem 2rem;
        background: white;
        border-radius: 20px;
        box-shadow: 0 20px 50px rgba(29, 39, 59, 0.12);
    }

    .brand {
        text-align: center;
        margin-bottom: 1.5rem;
    }

    .brand-icon {
        font-size: 2.2rem;
    }

    h1 {
        margin: 0.5rem 0 0.25rem;
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        -webkit-background-clip: text;
        background-clip: text;
        color: transparent;
        font-size: 1.6rem;
    }

    .subtitle {
        margin: 0;
        color: #657786;
        font-size: 0.9rem;
    }

    .form-group {
        margin-bottom: 1.1rem;
    }

    label {
        display: block;
        margin-bottom: 0.4rem;
        font-weight: 600;
        color: #14171a;
        font-size: 0.9rem;
    }

    input {
        width: 100%;
        padding: 0.75rem 0.9rem;
        border: 1px solid #e1e8ed;
        border-radius: 10px;
        font-size: 1rem;
        font-family: inherit;
        transition: border-color 0.15s ease, box-shadow 0.15s ease;
    }

    input:focus {
        outline: none;
        border-color: #1da1f2;
        box-shadow: 0 0 0 3px rgba(29, 161, 242, 0.15);
    }

    button {
        width: 100%;
        padding: 0.8rem;
        background: linear-gradient(120deg, #1da1f2 0%, #6a5cf5 100%);
        color: white;
        border: none;
        border-radius: 999px;
        font-size: 1rem;
        font-weight: 700;
        cursor: pointer;
        margin-top: 0.5rem;
        box-shadow: 0 8px 20px rgba(29, 161, 242, 0.3);
        transition: transform 0.15s ease, box-shadow 0.15s ease;
    }

    button:hover:not(:disabled) {
        transform: translateY(-1px);
        box-shadow: 0 10px 24px rgba(29, 161, 242, 0.4);
    }

    button:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }

    .error {
        color: #e0245e;
        text-align: center;
        margin: 0.5rem 0;
        font-size: 0.9rem;
    }

    .link {
        text-align: center;
        margin-top: 1.5rem;
        color: #657786;
        font-size: 0.9rem;
    }

    .link a {
        color: #1da1f2;
        font-weight: 600;
        text-decoration: none;
    }

    .link a:hover {
        text-decoration: underline;
    }
</style>
