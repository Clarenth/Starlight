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

<div class="ml-6 mb-4">
  {#key isOpen}
    <div class="" class:hidden={!isOpen} transition:slide={{ duration: 300 }}>
      <img alt="" src="" />
      <label for="password">Password</label>
      <input 
        id="password"  
        name="password" 
        type="password" 
        class="flex items-center h-12 px-4 bg-[#252525] hover:bg-[#1b1b1b] text-white mt-2 rounded focus:outline-none focus:ring-2" 
      />
    </div>
  {/key}
</div>

<style>
  button {
    border: none;
    background: none;
    color: inherit;
    cursor: pointer;
    margin: 0;
    padding-bottom: 0.5rem;
    padding-top: 0.5rem;
  }

  svg {
		flex-shrink: 0;
    transition: transform 0.2s ease-in;
  }

  [aria-expanded="true"] svg {
    transform: rotate(0.25turn);
  }
	
	p {
		margin: 0;
		margin-left: 1rem;
		color: #777;
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
