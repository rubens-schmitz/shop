<script lang="ts">
	import { goto } from '$app/navigation';

	import { getCategory } from '$lib/api/category.js';
	import { getURLIdParam } from '$lib/util.js';
	import { openModal } from '$lib/modal.js';

	export let actionFn: FormActionFn;
	export let actionName: string;
	export let makeRequest: MakeCategoryRequest;

	let id = undefined;
	let title: HTMLInputElement;

	function getInvalidField(): HTMLElement {
		if (title.value === '') return title;
		return undefined;
	}

	async function tryAction() {
		const invalidField = getInvalidField();
		if (invalidField !== undefined) {
			invalidField.focus();
			return;
		}
		let request = makeRequest(id, title.value);
		const rawRes = await actionFn(request);
		const res = await rawRes.json();
		openModal({ body: res.msg, task: 'alert' });
		goto('/login');
	}

	async function fetchCategory() {
		id = getURLIdParam();
		if (id === undefined) return { title: '' };
		const category = await getCategory(id);
		return category;
	}

	function onKeypress(e: KeyboardEvent) {
		if (e.key === 'Enter') tryAction();
	}
</script>

{#await fetchCategory()}
	<p>...waiting</p>
{:then category}
	<form>
		<div>
			<span>Title</span>
			<input
				on:keypress={onKeypress}
				value={category.title}
				bind:this={title}
				type="text"
			/>
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
