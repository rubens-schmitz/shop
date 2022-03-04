<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { goto } from '$app/navigation';

	import { getCategories } from '$lib/api/category.js';
	import List from '$lib/List.svelte';

	function gotoPutCategoryPage(id: number) {
		goto(`/put/category?id=${id}`);
		return new Promise(() => null);
	}
</script>

<svelte:head>
	<title>Edit category</title>
</svelte:head>

<section>
	<List
		categories
		getElements={getCategories}
		actionFn={gotoPutCategoryPage}
		actionName="Edit"
	/>
</section>
