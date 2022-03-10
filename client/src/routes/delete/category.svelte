<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { deleteCategory, getCategories } from '$lib/api/category.js';
	import List from '$lib/List.svelte';
</script>

<svelte:head>
	<title>Delete product</title>
</svelte:head>

<section>
	<List
		type="category"
		getElements={getCategories}
		actionFn={deleteCategory}
		actionName="Delete"
		updateAfterAction={true}
	/>
</section>
