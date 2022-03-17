<script lang="ts">
	import { dialog } from '$lib/stores.js';
	import { modalOkButton } from '$lib/modal.js';

	function closeModal() {
		if ($dialog.reload === true) window.location.reload();
		$dialog.task = '';
	}
</script>

{#if $dialog.task != ''}
	<div class="modal-overlay">
		<div class="modal-inside">
			{#if $dialog.task === 'successWithQRCode'}
				<span>Operation completed successfully</span>
				<span>Download the access qrcode</span>
				<img src={$dialog.qrcode} alt="qrcode" />
				<a
					class="qrcode-button"
					download="qrcode.png"
					href={$dialog.qrcode}
				>
					Download
				</a>
			{:else if $dialog.task === 'loginSuccess'}
				<span>Login completed successfully</span>
			{:else if $dialog.task === 'loginFailure'}
				<span>Login failed</span>
			{:else}
				<span>{$dialog.body}</span>
			{/if}
			<button id={modalOkButton} on:click={closeModal}>OK</button>
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
