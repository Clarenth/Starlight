<script lang="ts">
  // Svelte
	import { onMount } from "svelte";
  // Wails
  import {} from "$lib/wailsjs/go/auth/auth";
	import { GetAccounts, GetOneAccount } from "$lib/wailsjs/go/account/account";

  // Page Components
  import CreateAccount from "./_components/CreateAccount.svelte";
	import LoginDropdown from "./_components/LoginDropdown.svelte";
  
  // UI Components
  import { Button } from "$lib/components/ui/button"

  // Assets
  import logo from "$lib/images/logo.png"
	import DeleteAccount from "./_components/DeleteAccount.svelte";
	import DeleteAllData from "./_components/DeleteAllData.svelte";
	// import { account } from "$lib/stores/account";
  
  /* Leave as example
  import { LoadAccountData, SaveUserToDB, PrintDB } from "$lib/wailsjs/go/repo/DB";
  const newUser = () => {
    SaveUserToDB(username, password)
    PrintDB().then((result) => {
      console.log(result)
    })
  }
  */

  let accounts:any = [];
  let oneAccount: any;
  async function getOneAccount() {
    oneAccount = await GetOneAccount()
    console.log(oneAccount)
    return oneAccount;
  }

  async function getManyAccounts() {
    accounts = await GetAccounts()
    console.log(accounts)
    return accounts
  }

  async function consoleLogAccounts() {
    console.log(accounts)
  }

  onMount(() => {
    // getOneAccount()
    getManyAccounts()
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

  <section class=" bg-[#3a3a3a] shadow-md shadow-black p-9 mt-12 focus:outline-none focus:ring-2">
    <div class="flex flex-col justify-center items-center">
      <button on:click={getOneAccount}>Click Me</button>
      <button on:click={consoleLogAccounts}>console.log Accounts</button>
      <!-- TODO in future, make sure the screen does not rerender when a new user is made -->
      <div id="profile-list">
        {#each accounts as profile}
          <LoginDropdown profiles={{ id: profile.id, username: profile.username, created_at: profile.created_at, updated_at: profile.updated_at }} />
        {/each}
      </div>
      <div id="create-account" class="">
        <CreateAccount />
      </div>
    </div>
    <div id="delete-account" class="flex mt-6 items-end justify-end">
      <!-- Here we need to ask for the user's OS password (Linux, Windows) to delete the data-->
      <DeleteAllData />
    </div>
  </section>
</main>

<style lang="postcss">

</style>