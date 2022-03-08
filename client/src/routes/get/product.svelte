<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { getProduct } from '$lib/api/product.js';
	import { postItem } from '$lib/api/item.js';
	import { getURLIdParam } from '$lib/util.js';
	import { modal } from '$lib/stores.js';

	let title = '';
	let pictureContainer: HTMLDivElement;
	let preview: HTMLImageElement;
	let product: GetProductResponse;

	async function fetchProduct() {
		const id = getURLIdParam();
		product = await getProduct(id);
		title = product.title;
	}

	function select(pic: string) {
		preview.setAttribute('src', pic);
	}

	async function action(id: number) {
		await postItem(id);
		$modal = 'success';
	}
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<section>
	{#await fetchProduct()}
		<p>...waiting</p>
	{:then}
		<div class="flex-row">
			<div class="photos">
				{#each product.pictures as pic}
					<button on:click={() => select(pic)} type="button" class="photo">
						<img class="photo-image" src={pic} alt="preview" />
					</button>
				{/each}
			</div>

			<div>
				<div>
					Title
					<input type="text" value={product.title} readonly />
				</div>
				<div>
					Price
					<input type="number" value={product.price} readonly />
				</div>
				<div>
					Image
					<div class="preview-container" bind:this={pictureContainer}>
						<img
							bind:this={preview}
							src={product.pictures[0]}
							alt="product"
							class="preview-image"
						/>
					</div>
				</div>
				<button on:click={() => action(product.id)} type="button"
					>Add</button
				>
			</div>
		</div>
	{:catch error}
		<p>{error}</p>
	{/await}
</section>

<style>
	div {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
	button {
		width: 100%;
	}
	.flex-row {
		display: flex;
		flex-direction: row;
	}
</style>
