<script lang="ts">
	import { goto } from '$app/navigation';

	import Fa from 'svelte-fa/src/fa.svelte';
	import { faPlus } from '@fortawesome/free-solid-svg-icons';

	import { getProduct } from '$lib/api/product.js';
	import { getCategories } from '$lib/api/category.js';
	import { getURLIdParam } from '$lib/util.js';
	import { openModal } from '$lib/modal.js';

	interface Picture {
		id: number;
		preview: string;
	}

	export let actionFn: FormActionFn;
	export let actionName: string;
	export let makeRequest: MakeProductRequest;

	let id = undefined;
	let title: HTMLInputElement;
	let price: HTMLInputElement;
	let category: HTMLSelectElement;

	let picture: HTMLInputElement;
	let choose: HTMLButtonElement;

	let pictures: Picture[] = [{ id: 0, preview: '' }];
	let curPicId = 0;

	let container: HTMLDivElement;
	let preview: HTMLImageElement;
	let placeholder: HTMLSpanElement;
	let showPicture = false;

	/**
	 * Clean undefined pictures.
	 * Add an undefined picture at the end.
	 */
	function prep() {
		let result: Picture[] = [];
		for (let i = 0; i < pictures.length; i++)
			if (pictures[i].preview !== '') {
				if (pictures[i].id === curPicId) curPicId = result.length;
				result.push({
					id: result.length,
					preview: pictures[i].preview
				});
			}
		const id = result.length;
		result.push({ id, preview: '' });
		pictures = result;
		if (pictures[curPicId].preview === '') curPicId = 0;
	}

	function selectPic(id: number) {
		curPicId = id;
		if (pictures[id].preview === '') picture.click();
		else {
			showPicture = true;
			setTimeout(
				() => preview.setAttribute('src', pictures[id].preview),
				1
			);
		}
	}

	function deletePic(id: number) {
		if (pictures[curPicId].preview === '') return;
		picture.value = null;
		let result: Picture[] = [];
		for (let i = 0; i < pictures.length; i++)
			if (pictures[i].id !== id)
				result.push({
					id: result.length,
					preview: pictures[i].preview
				});
		pictures = result;
		curPicId = 0;
		if (pictures[0].preview === '') showPicture = false;
		else preview.setAttribute('src', pictures[0].preview);
	}

	function onChange() {
		const file = picture.files[0];
		if (file) {
			showPicture = true;
			const reader = new FileReader();
			reader.addEventListener('load', function () {
				const p = String(reader.result);
				preview.setAttribute('src', p);
				pictures[curPicId] = { id: curPicId, preview: p };
				prep();
			});
			reader.readAsDataURL(file);
			return;
		}
	}

	function getInvalidField(): HTMLElement {
		if (title.value === '') return title;
		if (price.value === '') return price;
		if (category.value === '') return category;
		if (pictures[0].preview === '') return choose;
		return undefined;
	}

	function getResquestPictures() {
		let result: string[] = [];
		for (let i = 0; i < pictures.length; i++)
			if (pictures[i].preview !== '') result.push(pictures[i].preview);
		return result;
	}

	async function tryAction() {
		const invalidField = getInvalidField();
		if (invalidField !== undefined) {
			invalidField.focus();
			return;
		}
		let request = makeRequest(
			id,
			title.value,
			parseFloat(price.value),
			parseInt(category.value),
			getResquestPictures()
		);
		const rawRes = await actionFn(request);
		const res = await rawRes.json();
		openModal({ body: res.msg, task: 'alert' });
		goto('/login');
	}

	async function fetchElement() {
		id = getURLIdParam();
		if (id === undefined) {
			return { title: '', price: '', categoryId: 0, pictures: [''] };
		}
		const product = await getProduct(id);
		pictures = [];
		for (let i = 0; i < product.pictures.length; i++) {
			pictures.push({ id: i, preview: product.pictures[i] });
		}
		pictures = [...pictures, { id: pictures.length, preview: '' }];
		showPicture = true;
		return product;
	}

	function onKeypress(e: KeyboardEvent) {
		if (e.key === 'Enter') tryAction();
	}
</script>

{#await fetchElement()}
	<p>...waiting</p>
{:then product}
	<form>
		<div class="photos">
			{#each pictures as pic}
				{#if pic.preview !== ''}
					<button
						on:click={() => selectPic(pic.id)}
						type="button"
						class="photo"
					>
						<img class="photo-image" src={pic.preview} alt="preview" />
					</button>
				{:else}
					<button
						on:click={() => selectPic(pic.id)}
						type="button"
						class="new-image-button"
					>
						<Fa icon={faPlus} />
					</button>
				{/if}
			{/each}
		</div>

		<div>
			<span>Title</span>
			<input
				on:keypress={onKeypress}
				value={product.title}
				bind:this={title}
				type="text"
			/>

			<span>Price</span>
			<input
				on:keypress={onKeypress}
				value={product.price}
				bind:this={price}
				type="number"
			/>

			<span>Category</span>
			{#await getCategories()}
				<p>loading...</p>
			{:then categories}
				<select
					value={product.categoryId}
					bind:this={category}
					name="category"
					id="category"
				>
					<option value="" />
					{#each categories as category}
						<option value={category.id}>{category.title}</option>
					{/each}
				</select>
			{/await}

			<span>Images</span>
			<input
				tabindex="-1"
				hidden
				type="file"
				bind:this={picture}
				on:change={onChange}
			/>
			<div class="buttons">
				<button
					bind:this={choose}
					on:click={() => picture.click()}
					type="button">Choose</button
				>
				<button on:click={() => deletePic(curPicId)} type="button"
					>Delete</button
				>
			</div>
			<div class="preview-container" bind:this={container}>
				{#if showPicture}
					<img
						bind:this={preview}
						src={product.pictures[0]}
						alt="product"
						class="preview-image"
					/>
				{:else}
					<span bind:this={placeholder}>Add image</span>
				{/if}
			</div>

			<button type="button" on:click={tryAction}>{actionName}</button>
		</div>
	</form>
{/await}

<style>
	form {
		display: flex;
		gap: 16px;
	}
	div {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
	.photo {
		padding: 0;
	}
	.new-image-button {
		width: 64px;
		height: 64px;
		display: flex;
		align-items: center;
		justify-content: center;
		background: var(--accent-color);
	}
	.buttons {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
</style>
