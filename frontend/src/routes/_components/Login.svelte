<script>

</script>
<!-- <h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p> -->
  <form on:submit|preventDefault action="#" method="post" class="flex flex-col bg-[#3a3a3a] rounded shadow-lg p-12 mt-12">
      <label class="font-semibold" for="username">Username</label>
      <input 
        id="username"
        name="username" 
        type="text" 
        class="flex items-center h-12 px-4 w-64 bg-[#252525] hover:bg-[#1b1b1b] mt-2 rounded focus:outline-none focus:ring-2" 
      />
      <!-- error message -->
        <span class="text-red-500 text-sm"></span>
      <!-- -->
      <label class="font-semibold mt-3" for="password">Password</label>
      <input 
        id="password"  
        name="password" 
        type="password" 
        class="flex items-center h-12 px-4 w-64 bg-[#252525] hover:bg-[#1b1b1b] mt-2 rounded focus:outline-none focus:ring-2" 
        />
      <!-- error message -->
        <span class="text-red-500 text-sm"></span>
      <!-- -->
      <input 
        id="remember"
        name="remember-me"
        type="checkbox"
        class="flex items-center"
      />
      <label for="remember">Remember me?</label>
      <button class="flex items-center justify-center h-12 px-6 w-64 bg-[#252525] hover:bg-[#1b1b1b] mt-8 rounded font-semibold text-white">Login</button>
      <div class="flex mt-6 justify-center text-xs">
          <a class="text-blue-400 hover:text-blue-500" href="#">Forgot Password</a>
          <span class="mx-2 text-gray-300">/</span>
          <a class="text-blue-400 hover:text-blue-500" href="#">Sign Up</a>
      </div>
  </form>

<style lang="postcss">

</style>