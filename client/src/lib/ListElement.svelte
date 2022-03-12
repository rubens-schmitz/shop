<script lang="ts">
	import { onMount } from 'svelte';

	import { putItem, deleteItem } from '$lib/api/item.js';

	export let element: ListElement;
	export let actionName: string;
	export let action: (id: number) => Promise<void>;
	export let update: (offsetDelta?: number) => Promise<any>;
	export let type: string;

	let href: string;

	onMount(() => {
		if (type === 'item' || type === 'dealItem') {
			href = `/get/product?id=${element.productId}`;
		} else if (type === 'deal') {
			href = `/get/deal?id=${element.id}`;
		} else {
			href = `/get/product?id=${element.id}`;
		}
	});

	async function removeAndUpdate() {
		await deleteItem(element.id);
		update();
	}

	async function changeQuantityAndUpdate(delta: number) {
		await putItem(element.id, element.quantity + delta);
		update();
	}
</script>

<div class="container">
	<a sveltekit:prefetch {href}>
		{#if type === 'product' || type === 'item'}
			<img class="preview" src={element.pictures[0]} alt="preview" />
			<span>{element.title}</span>
			<span>${element.price}</span>
		{:else if type == 'dealItem'}
			<img class="preview" src={element.pictures[0]} alt="preview" />
			<span>{element.title}</span>
			<span>${element.price}</span>
			{#if element.quantity > 1}
				<span>{element.quantity} units</span>
			{:else}
				<span>{element.quantity} unit</span>
			{/if}
		{:else if type === 'deal'}
			<div class="deal">
				<span>{element.datestamp}</span>
				<span>Price: {element.price}</span>
				<span>Quantity: {element.quantity}</span>
			</div>
		{:else}
			<span>{element.title}</span>
		{/if}
	</a>
	{#if type === 'item'}
		<div class="buttons">
			<button on:click={removeAndUpdate}>Remove</button>
			<div class="buttons-step">
				<button on:click={() => changeQuantityAndUpdate(-1)}>-</button>
				<div class="buttons-step-quantity">
					<span>{element.quantity}</span>
				</div>
				<button on:click={() => changeQuantityAndUpdate(+1)}>+</button>
			</div>
		</div>
	{:else if type === 'deal' || type === 'dealItem'}
		{#if actionName === 'Detail'}
			<a class="deal-button" {href}>{actionName}</a>
		{:else if actionName === 'Delete'}
			<button on:click={() => action(element.id)}>
				{actionName}
			</button>
		{/if}
	{:else if type === 'product' || type === 'category'}
		<button on:click={() => action(element.id)}>
			{actionName}
		</button>
	{/if}
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
	a {
		display: flex;
		flex-direction: column;
		color: var(--text-color);
		gap: 8px;
	}
	.preview {
		width: 200px;
		height: 200px;
	}

	.buttons {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
	.buttons-step {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr;
	}
	.buttons-step-quantity {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.deal {
		display: flex;
		flex-direction: column;
	}
	.deal-button {
		width: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
		cursor: pointer;
		background: var(--accent-color);
		color: white;
		font-weight: bold;
		border-radius: 4px;
		padding: 8px;
	}
	.deal-button:hover {
		text-decoration: none;
		background-color: var(--accent-color-darker);
	}
</style>
