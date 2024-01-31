<script lang="ts">
  const projectsList = [
    { name: "project#1" },
    { name: "project#2" },
    { name: "project#3" },
  ]

</script>
<nav class="leftsidebar">
  <button>
    <a href="/profile">Profile</a>
  </button>
  
  {#each projectsList as project}
    <ul>
      <li>
        <button type="button">{project.name}</button>
      </li>
    </ul>
  {/each}
</nav>


<style>

</style>