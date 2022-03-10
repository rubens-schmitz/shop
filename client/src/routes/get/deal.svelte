<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import List from '$lib/List.svelte';
	import { getDeal } from '$lib/api/deal.js';
	import { getItems } from '$lib/api/item';
	import { getURLIdParam } from '$lib/util.js';

	let title = '';
	let dealItem: GetDealResponse;

	async function fetchDealItem() {
		const id = getURLIdParam();
		dealItem = await getDeal(id);
		title = dealItem.datestamp;
		return dealItem;
	}

	async function getElements(limit, offset, title, categoryId) {
		return await getItems(
			limit,
			offset,
			title,
			categoryId,
			dealItem.cartId
		);
	}
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<section>
	{#await fetchDealItem()}
		<span>loading...</span>
	{:then dealItem}
		<List {dealItem} type="dealItem" {getElements} actionName="Detail" />
	{/await}
</section>
