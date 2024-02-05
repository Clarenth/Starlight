<script lang="ts">

  // Components
	import Button from "$lib/components/ui/button/button.svelte";
  import NewTask from "./NewTask.svelte";
  import Task from "./Task.svelte";

  // UI Components
  import Modal from "$lib/components/ui/modal/modal.svelte"
  import Textarea from "$lib/components/ui/textarea/textarea.svelte";

  // JavaScript
  let title: string;
  let description: string;
  let showModal: Boolean = false;
</script>

<div class="flex-shrink-0 overflow-x-scroll pr-1">
</div>
<header class="">
  <h1>Title of Sections</h1>
</header>
<div class="overflow-y-scroll">
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
  <Task />
</div>
<footer class="flex justify-center hover:visible">
  <Button 
    on:click={() => (showModal = true)}
    class="w-full justify-center items-center gap bg-inherit shadow-lg hover:bg-[#303030] focus:outline-none focus:ring-2"
  >
    Add Task
  </Button>
  <Modal bind:showModal>
    <h2 slot="header">
      Add Task
      <small class="">add a new task to this section</small>
    </h2>
    <form action="post" class=" text-white border border-gray/10">
      <Textarea
        id="title"
        placeholder="Title" 
        value={title}
        autofocus
        class="bg-[#252525] text-sm break-normal border-none"
      />
      <hr class="bg-white">
      <Textarea
        id="description"
        placeholder="Description" 
        value={title}
        class="bg-[#252525] text-sm break-normal border-none "
      />
      <!-- <p>Due Date</p> -->
    </form>
  </Modal>
</footer>
  
<style>

</style>