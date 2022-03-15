<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import List from '$lib/List.svelte';
	import { getProducts } from '$lib/api/product.js';
	import { postItem } from '$lib/api/item.js';
</script>

<svelte:head>
	<title>Catalog</title>
</svelte:head>

<section>
	<List
		showAlert
		type="product"
		getElements={getProducts}
		actionFn={postItem}
		actionName="Add"
	/>
</section>
