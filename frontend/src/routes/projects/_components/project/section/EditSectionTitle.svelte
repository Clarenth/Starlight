<script lang="ts">
	import Button from "$lib/components/ui/button/button.svelte";
import Input from "$lib/components/ui/input/input.svelte";
</script>


<Input
  type="text"
  placeholder="Title of Section"
  class="h-10 w-full bg-[#3a3a3a] justify-center items-center text-center text-base shadow-lg rounded-t focus:outline-none focus:ring-2"
  />

<Button
  type="submit"
  class="bg-red-900"
>
  Cancel
</Button>
<Button
  type="submit"
  class="bg-inherit relative"
>
  Save
</Button>


<style>
  
</style>