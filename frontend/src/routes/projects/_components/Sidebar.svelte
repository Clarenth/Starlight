<script lang="ts">
  const projectsList = [
    { name: "project#1" },
    { name: "project#2" },
    { name: "project#3" },
  ]

</script>
<nav class="leftsidebar flex-col">
  <p>Projects list</p>
  
  {#each projectsList as project}
    <ul class="flex-center justify-around hover:bg-gray-50/15 transition">
      <li>
        <button type="button">{project.name}</button>
      </li>
    </ul>
  {/each}
</nav>


<style>

</style>