<script lang="ts">
  import CreateAccount from "./CreateAccount.svelte";
	import ProfileDropdown from "./ProfileDropdown.svelte";

  import logo from "$lib/images/logo.png"
  // let src = '$lib/images/logo.png';

  import { LoadAccountData, SaveUserToDB, PrintDB } from "$lib/wailsjs/go/repo/DB";

  let username = "testy1"
  let password = "123456"
  const newUser = () => {
    SaveUserToDB(username, password)
    PrintDB().then((result) => {
      console.log(result)
    })
  }
  
  const data = [
    {
      q: "John",
    },
  ];

  // const data:any = [];
</script>
<main id="main" class="flex flex-col justify-center items-center bg-[#252525] text-white text-lg py-16">
  <section id="img-logo" class="flex flex-col items-center justify-center h-auto text-white text-lg">
    <h1 id="Starlight title" placeholder="Starlight" class="font-bold text-2xl">Starlight</h1>
    <img 
      src={logo}
      alt="logo" 
      placeholder="Starlight logo"
      class="w-52"
    />
  </section>

  <section class="flex flex-col justify-center items-center bg-[#3a3a3a] shadow-lg p-12 mt-12 focus:outline-none focus:ring-2">
    <div id="profile-list">
      {#each data as { q }}
        <ProfileDropdown profile={{ name: q }} />
      {/each}
    </div>
    <div id="create-account">
      <CreateAccount />
    </div>
  </section>

  <section id="click-me" class="flex flex-col justify-center items-center bg-[#3a3a3a] shadow-lg p-12 mt-12 focus:outline-none focus:ring-2">
    <button id="greeting" on:click={newUser}>
      Greet
    </button>
  </section>
  <a href="/projects">Projects</a>
</main>

<style lang="postcss">

</style>