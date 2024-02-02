<script lang="ts">
  import Checkmark from "$lib/images/check-dark-svgrepo-com.svg";
  import Edit from "$lib/images/edit-dark-svgrepo-com.svg"
</script>

<div class="flex flex-row gap-2 px-1 shadow-lg bg-[#303030] border border-black hover:cursor-pointer hover:border-white transition focus:outline-none focus:ring-2">
  <div class="flex flex-col gap-2 justify-between">
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
  <div class="flex flex-col justify-start items-start gap-1 min-w-72 max-w-72">
    <h3 class="">Investigate qOwnNotes for self-hosting</h3>
    <hr class="w-full opacity-50">
    <p class="line-clamp-3 break-normal overflow-hidden">
      Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance.
    </p>
    <!-- <p>Due Date</p> -->
  </div>
</div>

<style>

</style>