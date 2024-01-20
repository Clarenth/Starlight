<script>
	import Notes from "./_components/notes/Notes.svelte";
  import ProjectTable from "./_components/project/ProjectTable.svelte";
  import Sidebar from "./_components/Sidebar.svelte";
</script>

<h1>Projects Screen</h1>
<button 
  type="button"
>
  <a href="/">Signin Page</a>
</button>

<button>
  <a href="/account">Account</a>
</button>

<button>
  <a href="/notes">Notes</a>
</button>

<main>
  <Sidebar />
  <ProjectTable />
  <Notes />
</main>

<style>

</style>