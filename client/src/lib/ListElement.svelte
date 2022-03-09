<script lang="ts">
	import { onMount } from 'svelte';

	import { putItem, deleteItem } from '$lib/api/item.js';

	export let element: ListElement;
	export let actionName: string;
	export let action: (id: number) => Promise<void>;
	export let update: (offsetDelta?: number) => Promise<any>;
	export let cart: boolean;
	export let categories: boolean;

	let href: string;

	onMount(() => {
		if (cart) {
			href = `/get/product?id=${element.productId}`;
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
		{#if !categories}
			<img class="preview" src={element.pictures[0]} alt="preview" />
		{/if}
		<span>{element.title}</span>
		{#if !categories}
			<span>${element.price}</span>
		{/if}
	</a>
	{#if cart}
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
	{:else}
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
</style>
