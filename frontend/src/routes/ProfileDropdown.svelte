<script>
  import { slide } from "svelte/transition";
  export let profile;
  const { name } = profile;
  let isOpen = false;
  const toggle = () => (isOpen = !isOpen);
</script>

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
  {name}
</button>

<div class="mt-4 mb-4 shadow-xl">
  {#key isOpen}
    <div class="flex flex-col items-center justify-center bg-[#474747]" class:hidden={!isOpen} transition:slide={{ duration: 300 }}>
      <img alt="" src="" />
      <input 
        id="password"  
        name="password" 
        type="password" 
        placeholder="password"
        class="flex items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] text-white mt-2 rounded focus:outline-none focus:ring-2" 
      />
      <button type="submit" class="flex cursor-pointer items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] text-white mt-2 mb-2 rounded focus:outline-none focus:ring-2">
        Login
      </button>
    </div>
    <!-- <button class="flex items-center justify-center h-12 px-6 w-64 bg-[#252525] hover:bg-[#1b1b1b] mt-8 rounded font-semibold text-white">Login</button> -->

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
