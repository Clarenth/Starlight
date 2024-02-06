<script lang="ts">

  // Components
	import Button from "$lib/components/ui/button/button.svelte";
  import NewTask from "./NewTask.svelte";
  import Task from "./Task.svelte";
  import EditSectionTitle from "./section/EditSectionTitle.svelte";

  // UI Components
  import Modal from "$lib/components/ui/modal/modal.svelte"
  import Textarea from "$lib/components/ui/textarea/textarea.svelte";
	import Input from "$lib/components/ui/input/input.svelte";


  // JavaScript
  let title: string;
  let description: string;
  let showModal: Boolean = false;
  let dialog: HTMLDialogElement;

  function toggleEditTitle() {
    alert("TODO")
  }
</script>

<div class="flex-shrink-0 overflow-x-scroll pr-1">
</div>

<Button
  on:click={toggleEditTitle}
  class="w-full bg-[#3a3a3a] justify-center items-center shadow-lg rounded-b-none focus:outline-none focus:ring-2"
>
  <header>Title of Section</header>
</Button>
<div class="space-y-0.5 overflow-y-scroll">
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
    class="w-full bg-[#3a3a3a] justify-center items-center shadow-lg rounded-t-none focus:outline-none focus:ring-2"
  >
    Add Task
  </Button>
  <Modal bind:showModal bind:dialog>
    <h2 slot="header">
      Add Task
      <small class="">add a new task to this section</small>
    </h2>
    <form action="post" class="border border-none text-white">
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
      <p>Due Date</p>
    </form>
    <div class="flex flex-row justify-end gap-2 pt-1">
      <Button 
        on:click={() => dialog.close()}
        class="bg-[#3a3a3a] justify-center items-center gap shadow-lg text-base focus:outline-none focus:ring-2"
        >
        Save
      </Button>
      <Button 
        on:click={() => dialog.close()}
        class="bg-red-800 hover:bg-red-600"
      >
          Cancel
      </Button>
    </div>
  </Modal>
</footer>
  
<style>

</style>