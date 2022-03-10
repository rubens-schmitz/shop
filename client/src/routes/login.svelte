<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { page } from '$app/stores';
	import Fa from 'svelte-fa/src/fa.svelte';
	import {
		faPlus,
		faEdit,
		faMinus,
		faEye
	} from '@fortawesome/free-solid-svg-icons';
</script>

<svelte:head>
	<title>Login</title>
</svelte:head>

<section>
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
</style>
