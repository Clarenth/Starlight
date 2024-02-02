<script lang="ts">
  import Checkmark from "$lib/images/check-dark-svgrepo-com.svg";
  import Edit from "$lib/images/edit-dark-svgrepo-com.svg"
</script>

<div class="flex flex-row shadow-lg gap-1 bg-[#303030] max-w-80 border border-black hover:cursor-pointer hover:border-white transition focus:outline-none focus:ring-2">
  <div class="flex flex-col flex-shrink-0 justify-between pl-1">
    <img 
      src={Edit}
      alt="Complete"
      class="w-6 rounded-full mt-1 hover:bg-gray-700 transition-all"
    />
    <img 
      src={Checkmark} 
      alt="Edit"
      class="w-6 rounded-full mb-1 hover:bg-gray-700 transition-all"
    />
  </div>
  <div class="flex-col justify-start items-start py-1 pr-1 max-w-72">
    <h3 class="">Investigate qOwnNote for testing of the new</h3>
    <hr class="bg-white">
    <p class="text-sm line-clamp-3 break-normal overflow-hidden">
      Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance.
    </p>
    <!-- <p>Due Date</p> -->
  </div>
</div>

<!-- <div class="flex flex-col shadow-lg bg-[#303030] border border-black hover:cursor-pointer hover:border-white transition focus:outline-none focus:ring-2">
  <div class="flex-col justify-start items-start gap-1 min-w-72 max-w-72">
    <h3 class="">Investigate qOwnNote for testing of the new </h3>
    <p class="text-sm line-clamp-3 break-normal overflow-hidden">
      Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance.
    </p>
    <p>Due Date</p> -->
  <!-- </div> -->
  <!-- <div class="flex flex-row flex-shrink-0 gap-2 justify-between">
    <img 
      src={Edit}
      alt="Complete"
      class="w-6 rounded-full mt-1 hover:bg-black transition-all"
    />
    <img 
      src={Checkmark} 
      alt="Edit"
      class="w-6 rounded-full mb-1 hover:bg-black transition-all"
    />
  </div> -->
<!-- </div> -->

<style>

</style>