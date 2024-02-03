<script>
	import Button from "$lib/components/ui/button/button.svelte";
import NewTask from "./NewTask.svelte";
  import Task from "./Task.svelte";
</script>

<div class="flex-col flex-shrink-0 overflow-x-scroll">
  <header class="">
    <h1>Title of Sections</h1>
  </header>
  <div class="">
    <Task />
    <!-- <Task />
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
    <Task /> -->
  </div>
  <footer class="flex justify-center hover:visible">
    <Button class="w-full justify-center items-center gap bg-inherit shadow-lg text-base hover:bg-gray-50/5 focus:outline-none focus:ring-2">
      Tasks
    </Button>
  </footer>
</div>
  
<style>

</style>