<script lang="ts">
  // Svelte
  import { slide } from "svelte/transition";
  // Wails
  import { Login } from "$lib/wailsjs/go/auth/auth"
  // UI Components
	import Button from "$lib/components/ui/button/button.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import { enhance } from "$app/forms";
	import { goto } from "$app/navigation";
  
  // JavaScript
  export let profile;

  const { name } = profile;
  let password: string;
  let isOpen = false;
  
  const toggle = () => (isOpen = !isOpen);

  const handleLogin = (password: string) => {
    try {
      Login(name, password)
        .then((result: any) => {
          const data = result;
          console.log(data)
          return data;
        })
        .catch(error => console.log(error))
    } catch (error) {
      console.log(error)
      return error
    }
  }
</script>

<button class="flex w-72 border-b border-separate border-white" on:click={toggle} aria-expanded={isOpen}
  ><svg
    class="mr-1"
    width="20"
    height="20"
    fill="none"
    stroke-linecap="round"
    stroke-linejoin="round"
    stroke-width="2"
    viewBox="0 0 24 24"
    stroke="currentColor"><path d="M9 5l7 7-7 7" /></svg
  >
  {name}
</button>

<div class="mt-4 mb-4 shadow-xl">
  {#key isOpen}
    <form 
      method="post"
      on:submit|preventDefault={() => handleLogin(password)} 
      class="flex flex-col items-center justify-center bg-[#474747] text-lg border border-separate border-[#252525]" class:hidden={!isOpen} transition:slide={{ duration: 300 }}
    >
      <label class="mt-2" for="password" hidden></label>
      <Input 
        id="password"  
        name="password" 
        type="password" 
        placeholder="Password"
        bind:value={password}
        required
        class="flex items-center h-12 px-4 bg-[#252525] rounded-none text-white text-lg mt-2 mb-2 hover:bg-[#1b1b1b] focus:bg-[#1b1b1b] focus:outline-none focus:ring-2" 
      />
      <Button 
        type="submit"
        variant="default"
        class="flex cursor-pointer items-center h-12 px-4 bg-[#252525] rounded-none text-white text-lg mt-2 mb-2 hover:bg-[#1b1b1b] focus:bg-[#1b1b1b] focus:outline-none focus:ring-2" 
        on:click={(e) => {
          
        }}
      >
        Login
      </Button>
    </form>
  {/key}
</div>

<style>
  svg {
		flex-shrink: 0;
    transition: transform 0.2s ease-in;
  }

  [aria-expanded="true"] svg {
    transform: rotate(0.25turn);
  }
  
	.flex {
		display: flex;
	}
	.mr-1 {
		margin-right: 0.25rem;
	}
	.ml-6 {
		margin-left: 1.5rem;
	}
	.mb-4 {
		margin-bottom: 1rem;
	}
	.hidden {
		display: none;
	}
</style>
