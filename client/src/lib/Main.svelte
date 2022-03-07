<script lang="ts">
	$: innerWidth = 0;
</script>

<svelte:window bind:innerWidth />

{#if innerWidth < 512}
	<main class="paddingBottom">
		<slot />
	</main>
{:else}
	<main class="paddingTop">
		<slot />
	</main>
{/if}

<style>
	.paddingTop {
		padding-top: calc(var(--nav-height) + 48px);
	}
	.paddingBottom {
		padding-bottom: calc(var(--nav-height) + 48px);
	}
</style>
