<script lang="ts">
  // UI Components
	import Button from '$lib/components/ui/button/button.svelte';
	import { goto } from '$app/navigation';
	import { buttonVariants } from '$lib/components/ui/button';
  
  const projectsList = [
    { name: "Starlight" },
    { name: "Godgifu Server Code Server Code Server Code"},
    { name: "project#3" },
  ]

</script>
<!-- class="leftsidebar flex-col" --> 
{#each projectsList as project}
<!-- class="bg-[#3a3a3a] flex items-center mb-2 w-full text-md font-medium border border-gray-50/15 duration-75 group transition dark:text-white dark:hover:bg-gray-700" -->
  <p>redo Sidebar.svelte to use something other than button</p>
  <Button
    variant="link"
    href="/dashboard"
    class="flex-1 flex-col truncate flex-grow w-full mb-2 font-medium border text-white bg-[#3a3a3a] border-gray-50/15 duration-75 group transition dark:text-white dark:hover:bg-gray-700"
    aria-controls="dropdown-pages"
    data-collapse-toggle="dropdown-pages"
  >
    {project.name}
</Button>
{/each}


<style>

</style>