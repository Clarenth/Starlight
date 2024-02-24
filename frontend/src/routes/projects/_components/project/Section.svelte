<script lang="ts">
  // Svelte
  import { flip } from "svelte/animate";
  import { dndzone } from "svelte-dnd-action";

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

  let editSection = false;
  let addNewTask = false;

  let items = [
    {id: 1, name: "Hello test1", content: "from Seank"},
    {id: 2, name: "Investigate qOwnNote for testing of the new", content: "Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance."},
    {id: 3, name: "Linkin Survey results", content: "They show that JavaScript is a great high-concurrency language."},
    {id: 4, name: "This app have been built using", content: "SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE SVELTE ."},
    {id: 5, name: "The source of the mysterious ticking", content: "Snape, Snape, Severus Snape. Snape, Snape, Severus Snape. Snape, Snape, Severus Snape. Snape, Snape, Severus Snape. Snape, Snape, Severus Snape."},
    {id: 6, name: "Find a new base.", content: "We need a new base to attack Blue base."},
    {id: 7, name: "Why did the toys cross the road?", content: "To get to the toy store on the other side!"},
    {id: 8, name: "HAM", content: "HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM HAM ."},
    {id: 9, name: "Beep Beep", content: "Nurgh!."},
    {id: 10, name: "Create new greeting for the cryo patients", content: "Welcome to the world of tomorrow!!!"},
    {id: 11, name: "Investigate qOwnNote for testing of the new", content: "Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance."},
    {id: 12, name: "Investigate qOwnNote for testing of the new", content: "Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance."},
    {id: 13, name: "Investigate qOwnNote for testing of the new", content: "Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance."},
  ]

  const flipDurationMs = 100;
  function handleDndConsider(e:any ) {
      items = e.detail.items;
  }
  function handleDndFinalize(e: any) {
      items = e.detail.items;
  }
</script>

<div class="flex-shrink-0 overflow-x-scroll pr-1">
</div>

<!-- TODO
<p>Implement save to db functionality on create new task</p>
<p>Expande on the reuseable Modal for other components</p>
<p>Implement props for Task and Section, allowing for data passing when rendering via for-each in the Table</p>
<p>Look into moving Sidebar, Header, Footer into their own components for maintainability/expandability</p>
<p>Implement Authentication on the Login screen. Implement the data flow for Create Profile</p>
<p>Figure out moving components around the screen using svelte-dnd.</p>
<p>Can I hold everything in an internal array for each component state? Use svelte store based on which project is selected (project ID)?</p> -->

{#if editSection == false}
  <Button
    on:click={() => editSection = true}
    class="w-full bg-[#3a3a3a] justify-center items-center border border-gray-50/15 shadow-lg rounded-b-none focus:outline-none focus:ring-2"
  >
  <span>Title of Section</span>
  </Button>
  {:else}
    <Input
      type="text"
      placeholder="Title of Section"
      class="h-10 w-full bg-[#3a3a3a] justify-center items-center text-center border border-gray-50/15 shadow-lg rounded-t-md rounded-b-none focus:outline-none"
      autofocus
    />
    <div class="flex flex-row justify-end">
      <Button type="submit" class="bg-red-700 rounded-none hover:bg-red-900" on:click={() => editSection = false}>
        Cancel
      </Button>
      <Button type="submit" class="bg-[#3a3a3a] shadow-lg border border-gray-50/15 rounded-none focus:outline-none focus:ring-2">
        Save
      </Button>
    </div>
{/if}

<div use:dndzone="{{items, flipDurationMs}}" on:consider="{handleDndConsider}" on:finalize="{handleDndFinalize}" class="space-y-1 overflow-y-scroll">
  {#each items as task(task.id)}
  <div class="border-" animate:flip="{{duration: flipDurationMs}}">
      <Task title={task.name} content={task.content} />
    </div>
  {/each}
</div>

{#if addNewTask == false}
  <Button 
    on:click={() => (addNewTask = true)}
    class="w-full bg-[#3a3a3a] justify-center items-center border border-gray-50/15 shadow-lg rounded-t-none focus:outline-none focus:ring-2"
  >
    Add Task
  </Button>
  {:else}
    <form action="post" class="border border-black text-white">
      <Input
        id="title"
        placeholder="Title" 
        value={title}
        autofocus
        class="w-full bg-[#3a3a3a] text-base break-normal border border-gray-50/15 justify-center items-center text-center shadow-lg rounded-b-none focus:outline-none"
      />
      <hr class="bg-white">
      <Textarea
        id="description"
        placeholder="Task Description" 
        value={description}
        class="h-44 bg-[#3a3a3a] text-base border border-gray-50/15 rounded-t-none focus:outline-none"
      />
      <!-- <p>Due Date</p> -->
    </form>
    <div class="flex flex-row justify-end">
      <Button type="submit" class="bg-red-700 rounded-none hover:bg-red-900" on:click={() => addNewTask = false}>
        Cancel
      </Button>
      <Button type="submit" class="bg-[#3a3a3a] shadow-lg border border-gray-50/15 rounded-none focus:outline-none focus:ring-2">
        Save
      </Button>
    </div>
{/if}

<footer class="flex justify-center hover:visible">
  <!-- <Button 
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
      <Input
        id="title"
        placeholder="Title" 
        value={title}
        autofocus
        class="w-full bg-[#252525] text-base break-normal border-none"
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
  </Modal> -->
</footer>
  
<style>

</style>