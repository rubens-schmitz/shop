<script lang="ts">
	import { modal, modalMsg, qrcode } from '$lib/stores.js';

	function closeModal() {
		$modal = '';
	}

	function closeModalAndReload() {
		closeModal();
		window.location.reload();
	}
</script>

{#if $modal != ''}
	<div class="modal-overlay">
		<div class="modal-inside">
			{#if $modal === 'buyCart'}
				<span>Operation completed successfully</span>
				<span>Download the access qrcode</span>
				<img src={$qrcode} alt="qrcode" />
				<a class="qrcode-button" download="qrcode.png" href={$qrcode}>
					Download
				</a>
				<button on:click={closeModal}>OK</button>
			{:else if $modal === 'loginSuccess'}
				<span>Login completed successfully</span>
				<button on:click={closeModalAndReload}>OK</button>
			{:else if $modal === 'loginFailure'}
				<span>Login failed</span>
				<button on:click={closeModal}>OK</button>
			{:else}
				<span>{$modalMsg}</span>
				<button on:click={closeModal}>OK</button>
			{/if}
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
		text-decoration: none;
		background-color: var(--accent-color-darker);
	}
</style>
