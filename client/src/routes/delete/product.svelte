<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { deleteProduct, getProducts } from '$lib/api/product.js';
	import List from '$lib/List.svelte';
</script>

<svelte:head>
	<title>Delete product</title>
</svelte:head>

<section>
	<List
		type="product"
		getElements={getProducts}
		actionFn={deleteProduct}
		actionName="Delete"
		updateAfterAction={true}
	/>
</section>
