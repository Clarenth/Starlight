<script>
  // Components
  import Sidebar from "./_components/Sidebar.svelte";
  import Notes from "./_components/notes/Notes.svelte";
  import Table from "./_components/project/Table.svelte";

  // UI Components
  import Button from "$lib/components/ui/button/button.svelte";

  function isActive() {

  }
</script>

<div class="flex">
  <Sidebar />
  <div class="px-2">
    <header class="flex flex-row sticky pt-3 underline">
      <Button class="justify-center items-center gap bg-inherit shadow-lg text-base hover:bg-[#3a3a3a] focus:outline-none focus:ring-2">
        Tasks
      </Button>
      <Button class="justify-center items-center bg-inherit shadow-lg text-base hover:bg-[#3a3a3a] focus:outline-none focus:ring-2">
        Docs
      </Button>
    </header>
    <main class="flex flex-row mt-2"> <!--h-full -->
      <section class="w-screen h-screen">
        <Table />
      </section>
      <section>
        <!-- <Notes /> -->
      </section>
    </main>
  </div>
</div>

<style lang="postcss">

</style>