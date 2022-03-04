<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import CategoryForm from '$lib/CategoryForm.svelte';
	import { putCategory } from '$lib/api/category.js';

	function makePutCategoryRequest(id: number, title: string) {
		return { id, title };
	}
</script>

<svelte:head>
	<title>Put category</title>
</svelte:head>

<section>
	<CategoryForm
		actionFn={putCategory}
		actionName="Put"
		makeRequest={makePutCategoryRequest}
	/>
</section>
