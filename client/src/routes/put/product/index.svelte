<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import ProductForm from '$lib/ProductForm.svelte';
	import { putProduct } from '$lib/api/product.js';

	function makePutRequest(
		id: number,
		title: string,
		price: number,
		categoryId: number,
		pictures: string[]
	) {
		return { id, title, price, categoryId, pictures };
	}
</script>

<svelte:head>
	<title>Put product</title>
</svelte:head>

<section>
	<ProductForm
		actionFn={putProduct}
		actionName="Put"
		makeRequest={makePutRequest}
	/>
</section>
