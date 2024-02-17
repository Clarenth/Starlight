<script lang="ts">
  // Libraries

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
  ];

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

  <section class="flex flex-col justify-center items-center bg-[#3a3a3a] shadow-lg p-12 mt-12 focus:outline-none focus:ring-2">
    <div id="profile-list">
      {#each data as { q }}
        <LoginDropdown profile={{ name: q }} />
      {/each}
    </div>
    <div id="create-account">
      <CreateAccount />
    </div>
  </section>

  <Button href="/projects" class="flex flex-col justify-center items-center bg-[#3a3a3a] shadow-lg mt-12 text-lg focus:outline-none focus:ring-2">
    Login Test Button
  </Button>
</main>

<style lang="postcss">

</style>