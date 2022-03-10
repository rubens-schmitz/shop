<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import List from '$lib/List.svelte';
	import { getDeals } from '$lib/api/deal.js';
</script>

<svelte:head>
	<title>Deals</title>
</svelte:head>

<section>
	<List type="deal" getElements={getDeals} actionName="Detail" />
</section>
