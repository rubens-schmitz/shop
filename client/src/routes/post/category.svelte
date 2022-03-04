<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import CategoryForm from '$lib/CategoryForm.svelte';
	import { postCategory } from '$lib/api/category.js';

	function makePostCategoryRequest(id: number, title: string) {
		return { title };
	}
</script>

<svelte:head>
	<title>Post category</title>
</svelte:head>

<section>
	<CategoryForm
		actionFn={postCategory}
		actionName="Post"
		makeRequest={makePostCategoryRequest}
	/>
</section>
