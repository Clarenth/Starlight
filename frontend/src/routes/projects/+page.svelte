<script>
  // Svelte
  // Stores
  import { account } from "$lib/stores/account";
  // Components
  import Sidebar from "./_components/Sidebar.svelte";
  import Notes from "./_components/notes/Notes.svelte";
  import Table from "./_components/project/Table.svelte";

  // UI Components
  import Button from "$lib/components/ui/button/button.svelte";
	import { goto } from "$app/navigation";

  if($account == null || undefined) {
    goto("/")
  }
  function isActive() {

  }
</script>

<div class="antialiased h-screen flex  light:bg-white">
  <!-- <nav class="bg-white border-b border-gray-200 px-4 py-2.5 dark:bg-gray-800 dark:border-gray-700 fixed left-0 right-0 top-0 z-50">
    <div class="flex flex-wrap justify-between items-center">
      <div class="flex justify-start items-center">
        <Button class="justify-center items-center gap bg-inherit shadow-lg text-base hover:bg-[#3a3a3a] focus:outline-none focus:ring-2">
          Tasks
        </Button>
        <Button class="justify-center items-center bg-inherit shadow-lg text-base hover:bg-[#3a3a3a] focus:outline-none focus:ring-2">
          Docs
        </Button>
      </div>
      <div class="flex items-center lg:order-2">

      </div>
    </div>
  </nav> -->

  <!-- Sidebar -->
  <aside
    class="flex-none block top-0 left-0 z-40 w-36 pt-16 transition-transform -translate-x-full bg-default  md:translate-x-0"
    aria-label="Sidenav"
    id="drawer-navigation"
  >
    <div class="overflow-y-auto px-3 h-full">
      <Button
        type="button"
        aria-controls="dropdown-pages"
        data-collapse-toggle="dropdown-pages"
        class="bg-[#3a3a3a] flex items-center py-6 mb-10 w-full text-base font-medium border border-gray-50/15 rounded-lg duration-75 group transition dark:text-white dark:hover:bg-gray-700"
      >
        Profile
      </Button>
      <ul class="space-y-2">
        <li>
          <Button
            type="button"
            class="bg-[#3a3a3a] flex items-center p-2 w-full text-base font-medium border border-gray-50/15 rounded-lg duration-75 group transition dark:text-white dark:hover:bg-gray-700"
            aria-controls="dropdown-pages"
            data-collapse-toggle="dropdown-pages"
          >
            Project#1
          </Button>
          <ul id="dropdown-pages" class="hidden py-2 space-y-2">
          </ul>
        </li>
      <!-- <ul
        class="pt-5 mt-5 space-y-2 border-t border-pink-600 dark:border-gray-700"
      >
        <li>
          <a
            href="#"
            class="flex items-center p-2 text-base font-medium text-gray-900 rounded-lg transition duration-75 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-white group"
          >
            <svg
              aria-hidden="true"
              class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="M9 2a1 1 0 000 2h2a1 1 0 100-2H9z"></path>
              <path
                fill-rule="evenodd"
                d="M4 5a2 2 0 012-2 3 3 0 003 3h2a3 3 0 003-3 2 2 0 012 2v11a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm3 4a1 1 0 000 2h.01a1 1 0 100-2H7zm3 0a1 1 0 000 2h3a1 1 0 100-2h-3zm-3 4a1 1 0 100 2h.01a1 1 0 100-2H7zm3 0a1 1 0 100 2h3a1 1 0 100-2h-3z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span class="ml-3">Docs</span>
          </a>
        </li>
        <li>
          <a
            href="#"
            class="flex items-center p-2 text-base font-medium text-gray-900 rounded-lg transition duration-75 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-white group"
          >
            <svg
              aria-hidden="true"
              class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z"
              ></path>
            </svg>
            <span class="ml-3">Components</span>
          </a>
        </li>
        <li>
          <a
            href="#"
            class="flex items-center p-2 text-base font-medium text-gray-900 rounded-lg transition duration-75 hover:bg-gray-100 dark:hover:bg-gray-700 dark:text-white group"
          >
            <svg
              aria-hidden="true"
              class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-2 0c0 .993-.241 1.929-.668 2.754l-1.524-1.525a3.997 3.997 0 00.078-2.183l1.562-1.562C15.802 8.249 16 9.1 16 10zm-5.165 3.913l1.58 1.58A5.98 5.98 0 0110 16a5.976 5.976 0 01-2.516-.552l1.562-1.562a4.006 4.006 0 001.789.027zm-4.677-2.796a4.002 4.002 0 01-.041-2.08l-.08.08-1.53-1.533A5.98 5.98 0 004 10c0 .954.223 1.856.619 2.657l1.54-1.54zm1.088-6.45A5.974 5.974 0 0110 4c.954 0 1.856.223 2.657.619l-1.54 1.54a4.002 4.002 0 00-2.346.033L7.246 4.668zM12 10a2 2 0 11-4 0 2 2 0 014 0z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span class="ml-3">Help</span>
          </a>
        </li>
      </ul> -->
    </div>
  </aside>
    
  <!-- Project Table -->
  <main class="flex-1 flex flex-col overflow-hidden"> <!--p-4 md:ml-64 h-auto -->
    <!--Header-->
    <nav class="flex px-6 py-2 items-center flex-none">
      <div class="flex gap-2 flex-row">
        <Button class="bg-[#3a3a3a] justify-center items-center gap border border-gray-50/15 shadow-lg text-base focus:outline-none focus:ring-2">
          Tasks
        </Button>
        <Button class="bg-[#3a3a3a] justify-center items-center border border-gray-50/15 shadow-lg text-base focus:outline-none focus:ring-2">
          Docs
        </Button>
      </div>
    </nav>
    
    <h1 class="flex justify-center text-lg">Title of Project</h1>
        
    <!-- Project Table -->
    <div class="px-6 py-1 flex-1 overflow-scroll">
      <Table />
    </div>
  </main>
</div>

<style>

</style>