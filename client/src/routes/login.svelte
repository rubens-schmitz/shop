<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import Fa from 'svelte-fa/src/fa.svelte';
	import {
		faPlus,
		faEdit,
		faMinus,
		faEye,
		faUser
	} from '@fortawesome/free-solid-svg-icons';

	import { adminExist, createAdmin } from '$lib/api/access.js';
	import { modal, qrcode } from '$lib/stores.js';
	import { getCookieValue } from '$lib/util.js';
	import { loginAdmin } from '$lib/api/access.js';

	let picture: HTMLInputElement;

	let admin = false;

	onMount(() => updateAdmin());

	function updateAdmin() {
		admin = Boolean(getCookieValue('admin'));
	}

	function onChange() {
		const file = picture.files[0];
		if (file) {
			const reader = new FileReader();
			reader.addEventListener('load', function () {
				tryLoginAdmin(String(reader.result));
			});
			reader.readAsDataURL(file);
			return;
		}
	}

	async function tryLoginAdmin(qrcode) {
		let res = await loginAdmin(qrcode);
		if (res.success) $modal = 'loginSuccess';
		else $modal = 'loginFailure';
	}

	async function onCreateAdmin() {
		let res = await createAdmin();
		$qrcode = res.qrcode;
		$modal = 'accessResponse';
	}
</script>

<svelte:head>
	<title>Login</title>
</svelte:head>

<section>
	{#if admin}
		<div class:active={$page.url.pathname === '/post/product'}>
			<a sveltekit:prefetch href="/post/product">
				<span>Post product</span>
				<Fa icon={faPlus} />
			</a>
		</div>
		<div class:active={$page.url.pathname === '/put/product/choose'}>
			<a sveltekit:prefetch href="/put/product/choose">
				<span>Edit product</span>
				<Fa icon={faEdit} />
			</a>
		</div>
		<div class:active={$page.url.pathname === '/delete/product'}>
			<a sveltekit:prefetch href="/delete/product">
				<span>Delete product</span>
				<Fa icon={faMinus} />
			</a>
		</div>

		<br />

		<div class:active={$page.url.pathname === '/post/category'}>
			<a sveltekit:prefetch href="/post/category">
				<span>Post category</span>
				<Fa icon={faPlus} />
			</a>
		</div>
		<div class:active={$page.url.pathname === '/put/category/choose'}>
			<a sveltekit:prefetch href="/put/category/choose">
				<span>Edit category</span>
				<Fa icon={faEdit} />
			</a>
		</div>
		<div class:active={$page.url.pathname === '/delete/category'}>
			<a sveltekit:prefetch href="/delete/category">
				<span>Delete category</span>
				<Fa icon={faMinus} />
			</a>
		</div>

		<br />

		<div class:active={$page.url.pathname === '/get/deals'}>
			<a sveltekit:prefetch href="/get/deals">
				<span>See deals</span>
				<Fa icon={faEye} />
			</a>
		</div>
		<div class:active={$page.url.pathname === '/delete/deal'}>
			<a sveltekit:prefetch href="/delete/deal">
				<span>Delete deal</span>
				<Fa icon={faMinus} />
			</a>
		</div>
	{:else}
		{#await adminExist() then res}
			{#if !res.success}
				<button on:click={() => onCreateAdmin()}>
					<span>Create admin</span>
					<Fa icon={faPlus} />
				</button>
			{:else}
				<input
					tabindex="-1"
					hidden
					type="file"
					bind:this={picture}
					on:change={onChange}
				/>
				<button on:click={() => picture.click()} type="button"
					>Upload access qrcode</button
				>
			{/if}
		{/await}
	{/if}
</section>

<style>
	section {
		gap: 16px;
	}
	a {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 16px;
	}
	span {
		display: flex;
		align-items: center;
		justify-content: center;
		text-align: center;
	}
	button {
		width: auto;
		gap: 16px;
	}
</style>
