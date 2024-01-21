<script lang="ts">
  import CreateAccount from "./CreateAccount.svelte";
  import Login from "./Login.svelte"
	import ProfileDropdown from "./ProfileDropdown.svelte";

  const data = [
    {
      q: "John",
    },
    {
      q: "Smith",
    }
  ];

  // const data:any = [];
</script>

<main class="flex flex-col items-center justify-center h-screen bg-[#2c2c2c] text-white">
  <h1 class="font-bold text-2xl">Starlight</h1>
  <img alt="" src=""/>
  
  <div class="flex flex-col justify-center items-center bg-[#3a3a3a] rounded shadow-lg p-12 mt-12">
    <section id="profile-list">
      {#each data as { q }}
        <ProfileDropdown profile={{ name: q }} />
      {/each}
    </section>
    <section id="create-account">
      <CreateAccount />
    </section>
  </div>
  
  <ul>
    <li>
      <a href="/projects">Projects</a>
    </li>
  </ul>
</main>

<style lang="postcss">

</style>