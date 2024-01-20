<script lang="ts">
  const projectsList = [
    { name: "project#1" },
    { name: "project#2" },
    { name: "project#3" },
  ]

</script>

<h1>Sidebar</h1>

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

<style>

</style>