<script lang="ts">
  // Wails
  import { DeleteAccount } from "$lib/wailsjs/go/account/account"

  // Shadcn-Svelte
	import Button from "$lib/components/ui/button/button.svelte";

  export let profileID;
  const id:string = profileID;

  async function handleDeleteAccount(id: string) {
    try {
      DeleteAccount(id)
    } catch (error) {
      console.log("Account was not deleted")
    }
  } 
</script>

<Button
  type="button"
  variant="destructive"
  class="shadow-lg rounded-none"
  on:click={() => handleDeleteAccount(id)}
>
  Delete
</Button>

<style>

</style>