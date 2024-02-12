<script lang="ts">
  // Svelte
  import { slide } from "svelte/transition";
  // Wails
  import { CreateAccount } from "$lib/wailsjs/go/auth/auth"
  // UI Components
  import Button from "$lib/components/ui/button/button.svelte";
	import Input from "$lib/components/ui/input/input.svelte";

  let isOpen = false;
  const toggle = () => (isOpen = !isOpen);

  let username: String;
  let password: String;
  let confirmPassword: String;
</script>

<div class="border-b border-separate border-white mb-4">

  <button class="flex w-72" on:click={toggle} aria-expanded={isOpen}
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
  <h1>Create Account</h1>
</button>
</div>

<div class="shadow-xl">
  {#key isOpen}
    <form on:submit|preventDefault action="#" method="post" class="flex flex-col items-center justify-center bg-[#474747] text-lg border border-separate border-[#252525]" class:hidden={!isOpen} transition:slide={{ duration: 300 }}>
      <label class="mt-2" for="profile-name">Profile Name</label>
      <Input 
        id="profile-name" 
        name="profile-name" 
        placeholder="Profile Name"
        type="text" 
        class="flex items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] text-white mb-2 border border-separate border-[#252525] focus:bg-[#1b1b1b] focus:outline-none focus:ring-2" 
      />

      <label class="mt-2" for="password">Password</label>
      <Input 
          id="password"  
          name="password" 
          type="password" 
          placeholder="Password"
          class="flex items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] text-white mb-2 border border-separate border-[#252525] focus:bg-[#1b1b1b] focus:outline-none focus:ring-2" 
      />
      
      <label class="mt-2" for="confirm-password">Confirm Password</label>
      <Input 
          id="confirm-password"  
          name="confirm-password" 
          type="password" 
          placeholder="Confirm Password"
          value={confirmPassword}
          class="flex items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] text-white mb-2 border border-separate border-[#252525] focus:bg-[#1b1b1b] focus:outline-none focus:ring-2" 
      />
      <Button 
        href="/projects"
        variant="default"
        class="flex cursor-pointer items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] rounded-none text-white text-lg mt-2 mb-2 focus:bg-[#1b1b1b] focus:outline-none focus:ring-2"
        on:click={(e) => {
          
        }}
      >
        Create Profile
      </Button>
      <!-- <button type="submit" class="flex cursor-pointer items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] text-white mt-2 mb-2 focus:bg-[#1b1b1b] focus:outline-none focus:ring-2">
        <a href="/projects">Create Account</a>
      </button> -->
      <!-- <a 
        href="/projects"
        class="flex cursor-pointer items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] rounded-none text-white text-lg mt-2 mb-2 focus:bg-[#1b1b1b] focus:outline-none focus:ring-2"
        >
        Create Account
      </a> -->
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