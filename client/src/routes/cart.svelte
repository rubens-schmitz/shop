<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import List from '$lib/List.svelte';
	import { getItems } from '$lib/api/item.js';
</script>

<svelte:head>
	<title>Cart</title>
</svelte:head>

<section>
	<List cart getElements={getItems} />
</section>
