<script lang="ts">
	import { modalOpen } from '$lib/stores.js';

	function closeModal() {
		$modalOpen = false;
	}
</script>

{#if $modalOpen}
	<div class="modal-overlay">
		<div class="modal-inside">
			<span>Operation completed successfully</span>
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
		gap: 16px;
	}
</style>
