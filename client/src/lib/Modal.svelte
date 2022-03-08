<script lang="ts">
	import { modal, qrcode } from '$lib/stores.js';

	let a: HTMLAnchorElement;

	function prepHref() {
		a.href = $qrcode;
	}

	function closeModal() {
		$modal = '';
	}
</script>

{#if $modal != ''}
	<div class="modal-overlay">
		<div class="modal-inside">
			{#if $modal === 'buyCart'}
				<span>Operation completed successfully</span>
				<span>Please download the invoice</span>
				<img src={$qrcode} alt="qrcode" />
				<a
					bind:this={a}
					class="qrcode-button"
					download="qrcode.png"
					href={$qrcode}
				>
					Download
				</a>
			{:else}
				<span>Operation completed successfully</span>
			{/if}
			<button on:click={closeModal}>OK</button>
		</div>
	</div>
{/if}

<style>
	.modal-overlay {
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		position: fixed;
		background: rgba(0, 0, 0, 0.5);
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.modal-inside {
		padding: 16px;
		background: white;
		border-radius: 4px;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 16px;
	}
	.qrcode-button {
		width: 100%;
		display: flex;
		justify-content: center;
		cursor: pointer;
		background: var(--accent-color);
		color: white;
		font-weight: bold;
		border-radius: 4px;
		padding: 8px;
	}
	.qrcode-button:hover {
		background-color: var(--accent-color-darker);
	}
</style>
