<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import List from '$lib/List.svelte';
	import { getProducts } from '$lib/api/product.js';
	import { postItem } from '$lib/api/item.js';
	import { modal } from '$lib/stores.js';

	async function actionFn(id: number) {
		await postItem(id);
		$modal = 'success';
	}
</script>

<svelte:head>
	<title>Catalog</title>
</svelte:head>

<section>
	<List getElements={getProducts} {actionFn} actionName="Add" />
</section>
