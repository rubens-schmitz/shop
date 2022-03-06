<script lang="ts">
	import { page } from '$app/stores';
	import Fa from 'svelte-fa/src/fa.svelte';
	import {
		faHome,
		faSearch,
		faShoppingCart,
		faUser
	} from '@fortawesome/free-solid-svg-icons';

	console.log($page.url.pathname);
</script>

<nav>
	<div class:active={$page.url.pathname === '/'}>
		<a sveltekit:prefetch href="/">
			<Fa icon={faHome} size="2x" />
            <span>Home</span>
		</a>
	</div>
	<div class:active={$page.url.pathname === '/catalog'}>
		<a sveltekit:prefetch href="/catalog">
			<Fa icon={faSearch} size="2x" />
            <span>Catalog</span>
		</a>
	</div>
	<div class:active={$page.url.pathname === '/cart'}>
		<a sveltekit:prefetch href="/cart">
			<Fa icon={faShoppingCart} size="2x" />
            <span>Cart</span>
		</a>
	</div>
	<div class:active={$page.url.pathname === '/login'}>
		<a sveltekit:prefetch href="/login">
			<Fa icon={faUser} size="2x" />
            <span>Login</span>
		</a>
	</div>
</nav>

<style>
	nav {
		max-width: 375px;
		width: 100%;
		display: flex;
		justify-content: space-around;
	}
    a {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
	div > * {
		color: var(--text-color);
	}
	.active > * {
		color: var(--accent-color);
	}
</style>
