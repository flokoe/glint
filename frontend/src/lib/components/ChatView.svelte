<script lang="ts">
	import { onMount } from 'svelte';

	import Sidebar from './Sidebar.svelte';
	import SidebarCollapsed from './SidebarCollapsed.svelte';
	import SearchDialog from './SearchDialog.svelte';

	let sidebarCollapsed = false;
	let searchDialog: HTMLDialogElement;

	// 	let messages: Array<{
	// 	id: number;
	// 	role: 'user' | 'assistant';
	// 	content: string;
	// 	timestamp: string;
	// }> = [];

	onMount(() => {
		const handleKeydown = (e: KeyboardEvent) => {
			if (e.ctrlKey && e.key === 'b') {
				e.preventDefault();
				sidebarCollapsed = !sidebarCollapsed;
			}
			if ((e.ctrlKey && e.key === 'k') || e.key === '/') {
				e.preventDefault();
				searchDialog?.showModal();
			}
			if (e.ctrlKey && e.shiftKey && e.key === 'O') {
				e.preventDefault();
				// TODO: Trigger new chat
			}
		};

		window.addEventListener('keydown', handleKeydown);
		return () => window.removeEventListener('keydown', handleKeydown);
	});
</script>


<div>
	<SearchDialog bind:dialog={searchDialog} />

	<Sidebar bind:open={sidebarCollapsed} bind:searchDialog />
	<SidebarCollapsed bind:open={sidebarCollapsed} bind:searchDialog />

	<main class="{sidebarCollapsed ? 'ml-80' : 'ml-16'}">
		<slot />
	</main>
</div>
