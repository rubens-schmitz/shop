<script lang="ts">
	import Nav from '$lib/Nav.svelte';
	import Modal from '$lib/Modal.svelte';
	import Main from '$lib/Main.svelte';
	import '../app.css';

	$: innerWidth = 0;
</script>

<svelte:head>
	<link rel="icon" href="/favicon.png" />
</svelte:head>

<svelte:window bind:innerWidth />

{#if innerWidth < 512}
	<div class="container bottom">
		<Nav />
	</div>
{:else}
	<div class="container top">
		<Nav />
	</div>
{/if}

<Main>
	<slot />
</Main>

<Modal />

<style>
	.container {
		width: 100%;
		height: var(--nav-heigh);
		position: fixed;
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 20px;
		background: var(--tertiary-color);
	}
	.top {
		top: 0;
	}
	.bottom {
		bottom: 0;
	}
</style>
