<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { goto } from '$app/navigation';

	import { getProducts } from '$lib/api/product.js';
	import List from '$lib/List.svelte';

	function gotoPutPage(id: number) {
		goto(`/put/product?id=${id}`);
		return new Promise(() => null);
	}
</script>

<svelte:head>
	<title>Edit product</title>
</svelte:head>

<section>
	<List
		getElements={getProducts}
		actionFn={gotoPutPage}
		actionName="Edit"
	/>
</section>
