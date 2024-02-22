<script lang="ts">
  // Svelte
	import { onMount } from "svelte";
  // Wails
  import {} from "$lib/wailsjs/go/auth/auth";

  // Page Components
  import CreateAccount from "./_components/CreateAccount.svelte";
	import LoginDropdown from "./_components/LoginDropdown.svelte";
  
  // UI Components
  import { Button } from "$lib/components/ui/button"

  // Assets
  import logo from "$lib/images/logo.png"
  
  /* Leave as example
  import { LoadAccountData, SaveUserToDB, PrintDB } from "$lib/wailsjs/go/repo/DB";
  const newUser = () => {
    SaveUserToDB(username, password)
    PrintDB().then((result) => {
      console.log(result)
    })
  }
  */
  
  const data = [
    {
      q: "John",
    },
    {
      q: "Sally Sparrow",
    },
  ];

  onMount(() => {

  })
</script>
<main id="main" class="flex flex-col justify-center items-center bg-[#252525] text-white text-lg py-16">
  <section id="img-logo" class="flex flex-col items-center justify-center h-auto text-white text-lg">
    <h1 id="Starlight title" placeholder="Starlight" class="font-bold text-2xl">Starlight</h1>
    <img 
      src={logo}
      alt="logo" 
      placeholder="Starlight logo"
      class="w-40"
    />
  </section>

  <section class="flex flex-col justify-center items-center bg-[#3a3a3a] shadow-md shadow-black p-12 mt-12 focus:outline-none focus:ring-2">
    <div id="profile-list">
      {#each data as { q }}
        <LoginDropdown profileName={{ name: q }} />
      {/each}
    </div>
    <div id="create-account" class="">
      <CreateAccount />
    </div>
  </section>
</main>

<style lang="postcss">

</style>