<script lang="ts">
  import Checkmark from "$lib/images/check-dark-svgrepo-com.svg";
  import Edit from "$lib/images/edit-dark-svgrepo-com.svg"
</script>

<div class="flex flex-row mt-3 gap-2 shadow-lg bg-[#303030] border border-black hover:border-white transition focus:outline-none focus:ring-2">
  <div class="flex flex-col ml-1 justify-between">
    <img 
      src={Edit}
      alt="Finished"
      class="w-6 rounded-full mt-1 hover:bg-black transition-all"
    />
    <img 
      src={Checkmark} 
      alt="Edit"
      class="w-6 rounded-full mb-1 hover:bg-black transition-all"
    />
  </div>
  <div class="flex flex-col justify-start items-start gap min-w-60 max-w-60 text-wrap break-all">
    <h3>Investigate qOwnNotes for self-hosting</h3>
    <p class="text-wrap break-all">Description</p>
    <p>Due Date</p>
  </div>
</div>

<style>

</style>