<script>
    import { UserRoundPlus , AlertTriangle } from "@lucide/svelte";
    import NavAdmin from "../../components/NavAdmin.svelte";
  
    let showAlert = false;
    let showDelete = '';
  
    let obj = [
        {
            Name_Product: "Super Sol",
            Id:"266",
        	Unit:"5kg",
        	Stock:5,
            Purchase_Price:5000,
        	Urchase_Date:"2024-03-13",
        	Time:"06:00"
        }
    ];
  
    let nameSearch = '';
  
    // Data yang difilter berdasarkan nameSearch
    $: filteredObj = nameSearch 
        ? obj.filter(item => item.Name_Product.toLowerCase().includes(nameSearch.toLowerCase())) 
        : obj;
  </script>
  
  <NavAdmin title={"Warehouse"}>
    <div class="w-full flex justify-between items-center">
      <div>
        <input type="search" class="bg-white shadow-xl p-2 rounded-md text-black" 
               placeholder="Search Name" bind:value={nameSearch} />
      </div>
      <a href="/register" class="flex w-30 h-10 bg-blue-500 text-white p-2 rounded-xl">
        <UserRoundPlus /> Add Users
      </a>
    </div>
  
    <div class="relative my-3 overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left rtl:text-right text-gray-400">
        <thead class="uppercase bg-gray-800 text-white font-bold">
          <tr>
            <th class="px-6 py-3">No</th>
            <th class="px-6 py-3">name product</th>
            <th class="px-6 py-3">Id</th>
            <th class="px-6 py-3">unit</th>
            <th class="px-6 py-3">stock</th>
            <th class="px-6 py-3">Purchase Price</th>
            <th class="px-6 py-3">Urchase Date</th>
            <th class="px-6 py-3">Time</th>
            <th class="px-6 py-3">Action</th>
          </tr>
        </thead>
        <tbody>
          {#each filteredObj as obj, i}
            <tr class="odd:bg-white bg-gray-900 even:bg-gray-50 even:bg-gray-200 text-gray-600">
              <th class="px-6 py-4">{i + 1}</th>
              <td class="px-6 py-4 font-medium">{obj.Name_Product}</td>
              <td class="px-6 py-4">{obj.Id}</td>
              <td class="px-6 py-4">{obj.Unit}</td>
              <td class="px-6 py-4">{obj.Stock}</td>
              <td class="px-6 py-4">{obj.Purchase_Price}</td>
              <td class="px-6 py-4">{obj.Urchase_Date}</td>
              <td class="px-6 py-4">{obj.Time}</td>
              <td class="px-6 py-4 flex">
                <button on:click={() => {showDelete = obj.Name_Product; showAlert = !showAlert}}
                        class="font-medium text-white bg-red-500 rounded-full py-1 px-2">
                  Delete
                </button>
                <a href="/views-users" class="font-medium text-white bg-green-500 rounded-full py-1 px-2">
                  Views
                </a>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  
    <!-- alert confirm -->
    {#if showAlert}
      <div class="fixed inset-0 backdrop-blur-sm flex items-center justify-center p-4">
        <div class="bg-gray-950 text-red-800 rounded-lg shadow-lg p-6 max-w-sm w-full">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center">
              <AlertTriangle class="w-8 h-8 text-red-500 mr-2" />
              <h2 class="text-xl font-semibold">Warning</h2>
            </div>
          </div>
          <p class="text-white mb-6">
            Kamu yakin ingin menghapus <span>{showDelete}</span> dari database?
          </p>
          <div class="flex justify-end space-x-4">
            <button on:click={() => showAlert = !showAlert} class="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg">
              Cancel
            </button>
            <button class="px-4 py-2 bg-red-500 text-white hover:bg-red-900 rounded-lg">
              Delete
            </button>
          </div>
        </div>
      </div>
    {/if}
  </NavAdmin>
  