<script lang="ts">
  import Checkmark from "$lib/images/check-dark-svgrepo-com.svg";
  import Edit from "$lib/images/edit-dark-svgrepo-com.svg"

  // UI Components
  import Button from "$lib/components/ui/button/button.svelte";
  import Input from "$lib/components/ui/input/input.svelte";
	import Textarea from "$lib/components/ui/textarea/textarea.svelte";

  let title: string;
  let description: string;
  let showModal: Boolean = false;
  let dialog: HTMLDialogElement;
</script>

<div class="flex flex-col shadow-lg gap-1 bg-[#303030] max-w-80 border border-black hover:cursor-pointer hover:border-white transition focus:outline-none focus:ring-2">
  <div class="flex flex-col flex-shrink-0 justify-between pl-1">
    <img 
      src={Edit}
      alt="Complete"
      class="w-6 rounded-full mt-1 hover:bg-gray-700 transition-all"
    />
    <img 
      src={Checkmark} 
      alt="Edit"
      class="w-6 rounded-full mb-1 hover:bg-gray-700 transition-all"
    />
  </div>
  <form action="post" class="flex-col justify-start items-start py-1 pr-1 max-w-72">
    <Textarea
      id="title"
      placeholder="Title" 
      value={title}
      class="bg-[#252525] border-none"
    />
    <hr class="bg-white">
    <Textarea
      id="description"
      placeholder="Description" 
      value={title}
      class="bg-[#252525] text-sm break-normal border-none"
    />
    <!-- <p>Due Date</p> -->
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
</div>

<style>

</style>