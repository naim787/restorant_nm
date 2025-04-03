<script>
  import NavAdmin from "../../components/NavAdmin.svelte";

  let value = [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
</script>


<NavAdmin title={"Tabel"}>
    <div class="w-full flex justify-between items-center">
        <div class="flex">
            <a href="/" class="w-30 h-10 bg-yellow-500 text-white p-2 rounded-xl">Add room</a>
            <div class="w-[40vw] flex justify-evenly items-center overflow-auto">
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">Lumba"</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">Cakalag</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">Hiu</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">paus</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
                <p class="py-1 px-3 rounded-md bg-gray-500 inline-block text-white mx-2">dugog</p>
            </div>
        </div>
        <a href="/" class="flex w-30 h-10 bg-blue-500 text-white p-2 rounded-xl">Add Teibel</a>
    </div>
    <div class="w-[88vw] h-[90vh] flex flex-wrap justify-evenly items-start overflow-scroll gap-2">
        {#each value as v, i}
            <div class="w-42 h-42 bg-green-500 rounded-md flex">
                <h1 class="m-auto text-[2vw] font-bold">{i}</h1>
            </div>   
        {/each}
    </div>
</NavAdmin>