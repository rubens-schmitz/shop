<script lang="ts">
	import { getCategory } from '$lib/api/category.js';
	import { getURLIdParam } from '$lib/util.js';
    import { modal } from '$lib/stores.js'

	export let actionFn: FormActionFn;
	export let actionName: string;
	export let makeRequest: MakeCategoryRequest;

	let id = undefined;
	let title: HTMLInputElement;

	function getInvalidField(): HTMLElement {
		if (title.value === '') return title;
		return undefined;
	}

	function tryAction() {
		const invalidField = getInvalidField();
		if (invalidField !== undefined) {
			invalidField.focus();
			return;
		}
		let request = makeRequest(id, title.value);
		actionFn(request);
        $modal = 'success'
	}

	async function fetchCategory() {
		id = getURLIdParam();
		if (id === undefined) return { title: '' };
		const category = await getCategory(id);
		return category;
	}
</script>

{#await fetchCategory()}
	<p>...waiting</p>
{:then category}
	<form>
		<div>
			<span>Title</span>
			<input value={category.title} bind:this={title} type="text" />
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
</style>
