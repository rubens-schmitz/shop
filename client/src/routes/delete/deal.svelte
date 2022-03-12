<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import List from '$lib/List.svelte';
	import { deleteDeal, getDeals } from '$lib/api/deal';
</script>

<svelte:head>
	<title>Delete deal</title>
</svelte:head>

<section>
	<List
		type="deal"
		getElements={getDeals}
		actionFn={deleteDeal}
		actionName="Delete"
		updateAfterAction={true}
	/>
</section>
