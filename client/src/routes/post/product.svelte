<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import ProductForm from '$lib/ProductForm.svelte';
	import { postProduct } from '$lib/api/product.js';

	function makePostRequest(
		id: number,
		title: string,
		price: number,
		categoryId: number,
		pictures: string[]
	) {
		return { title, price, pictures, categoryId };
	}
</script>

<svelte:head>
	<title>Post product</title>
</svelte:head>

<section>
	<ProductForm
		actionFn={postProduct}
		actionName="Post"
		makeRequest={makePostRequest}
	/>
</section>
