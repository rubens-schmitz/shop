<script lang="ts">
	import ListElement from './ListElement.svelte';
	import { getCategories } from '$lib/api/category.js';

	let search = '';
	let elements = [];
	let offset = 0;

	let category: number;

	export let limit = 16;
	export let getElements: ListGetElementsFn;
	export let actionFn: ListActionFn = undefined;
	export let actionName: string = undefined;
	export let updateAfterAction = false;
	export let cart = false;
	export let categories = false;

	function searchChanged() {
		offset = 0;
		update();
	}

	async function action(id: number) {
		await actionFn(id);
		if (updateAfterAction) update();
	}

	async function update(offsetDelta: number = 0) {
		// TODO: fix this
		// if you remove this, the select categories won't work
		// properly, it will mix the values.
		await getCategories();

		const newOffset = offset + offsetDelta;
		if (newOffset < 0) return;
		const data = await getElements(limit, newOffset, search, category);
		if (data.length === 0 && newOffset !== offset) return;
		if (data.length === 0 && newOffset !== 0) {
			update(-limit);
			return;
		}
		offset = newOffset;
		elements = data;
	}
</script>

<div class="container">
	<div class="controls">
		<input
			type="text"
			placeholder="Search"
			bind:value={search}
			on:input={() => searchChanged()}
		/>

		{#if !categories}
			{#await getCategories()}
				<p>loading...</p>
			{:then categories}
				<select
					on:change={() => searchChanged()}
					bind:value={category}
					name="category"
					id="category"
				>
					<option selected value={0}>All</option>
					{#each categories as category}
						<option value={category.id}>{category.title}</option>
					{/each}
				</select>
			{/await}
		{/if}
	</div>

	{#await update()}
		<p>...waiting</p>
	{:then}
		<div class="elements">
			{#each elements as element}
				<ListElement
					{categories}
					{update}
					{cart}
					{element}
					{actionName}
					{action}
				/>
			{/each}
		</div>
		<div class="step">
			<button on:click={() => update(-limit)}>Previous</button>
			<button on:click={() => update(limit)}>Next</button>
		</div>
	{:catch error}
		<p>{error}</p>
	{/await}
</div>

<style>
	.container {
        width: 100%;
		gap: 32px;
		display: flex;
        flex: 1;
		flex-direction: column;
		align-items: center;
	}
    .controls {
        width: 200px;
        display: flex;
        flex-direction: column;
        gap: 8px;
    }
	.elements {
		display: flex;
        flex: 1;
		flex-wrap: wrap;
        justify-content: center;
        gap: 32px;
	}
	.step {
        width: 100%;
		display: flex;
		justify-content: center;
		gap: 16px;
        padding: 16px;
        background: var(--tertiary-color);
	}
    .step > button {
        width: 100px;
    }
</style>
