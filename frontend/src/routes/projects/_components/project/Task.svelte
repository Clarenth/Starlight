<script lang="ts">
  // UI Components
  // Assets
  import Checkmark from "$lib/images/check-dark-svgrepo-com.svg";
  import Edit from "$lib/images/edit-dark-svgrepo-com.svg"
	import { string } from "zod";

  // JavaScript
  export let title: string;
  export let content: string;
</script>

<div 
  draggable="true"
  class="flex flex-row shadow-lg gap-1 bg-[#303030] max-w-80 border border-gray-50/15 hover:cursor-pointer hover:border-white transition focus:outline-none focus:ring-2" >
  <div class="flex flex-col flex-shrink-0 justify-between pl-1">
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
  </div>
  <div class="flex-col justify-start items-start py-1 pr-1 max-w-72">
    <h3 class="">
      {title}
      <!-- Investigate qOwnNote for testing of the new -->
    </h3>
    <hr class="bg-white">
    <small class="line-clamp-3 break-normal overflow-hidden">
      {content}
      <!-- Officially recognized by the Duden dictionary as the longest word in German, Kraftfahrzeughaftpflichtversicherung is a 36 letter word for motor vehicle liability insurance. -->
    </small>
    <!-- <p>Due Date</p> -->
  </div>
</div>

<style>

</style>