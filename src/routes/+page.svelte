<script>
    import { UserRoundPlus, AlertTriangle, Trash2 } from "@lucide/svelte";
    import NavAdmin from "../components/NavAdmin.svelte";
  
    let showAlert = false;
    let showDelete = '';
  
    let obj = [];

    

async function fetchUsers() {
    try {
        const response = await fetch("http://127.0.0.1:3000/users"); // Panggil endpoint API
        const data = await response.json();

        if (data && typeof data === "object") { // Pastikan data adalah object
            obj = [data]; // ✅ Jadikan array supaya bisa ditampilkan dengan {#each}
        }
        
    } catch (error) {
        console.error("❌ Gagal mengambil data:", error);
    }
}

// Ambil data saat komponen dimuat
fetchUsers();


  
    let nameSearch = '';
  
    // Data yang difilter berdasarkan pencarian nama
    $: filteredObj = nameSearch
      ? obj.filter(item => item.Name.toLowerCase().includes(nameSearch.toLowerCase()))
      : obj;
  
    // Menyimpan ID dari checkbox yang dicentang
    let selectedIds = [];
  
    // Fungsi untuk menangani checkbox
    function toggleSelection(empId) {
      if (selectedIds.includes(empId)) {
        selectedIds = selectedIds.filter(id => id !== empId);
      } else {
        selectedIds = [...selectedIds, empId];
      }
    }
  
    // Fungsi saat ikon Trash diklik
    function handleDeleteClick() {
      let selectedEmployees = obj.filter(emp => selectedIds.includes(emp.Id));
  
      let selectedNames = selectedEmployees.map(emp => emp.Name);

      // jika ada data nnya, maka gabugkan semua array pisahkan deggan , (koma)
      if (selectedNames.length) {
        showDelete = selectedNames.join(", ");
        showAlert = true;
      }
    }
  </script>
  
  <NavAdmin title={"Staff"}>
    <div class="w-full flex justify-between items-center">
      <div class="flex justify-start items-center">
        <input type="checkbox" name="" id=""/>
        <input type="search" class="bg-white shadow-xl p-2 rounded-md text-black"
               placeholder="Search Name" bind:value={nameSearch} />
               <button  class="p-2 bg-red-950 rounded-md mx-2 text-red-500 cursor-pointer flex" on:click={handleDeleteClick}>
                   <Trash2 size={20} class="m-auto"/>
               </button>
      </div>
      <a href="/register" class="flex w-30 h-10 bg-blue-500 text-white p-2 rounded-xl">
        <UserRoundPlus /> Add Users
      </a>
    </div>
  
    <div class="relative my-3 overflow-x-auto shadow-md sm:rounded-lg">
      <table class="w-full text-sm text-left rtl:text-right text-gray-400">
        <thead class="uppercase bg-gray-800 text-white font-bold">
          <tr>
            <th class="px-6 py-3">#</th>
            <th class="px-6 py-3">User Name</th>
            <th class="px-6 py-3">Id</th>
            <th class="px-6 py-3">Role</th>
            <th class="px-6 py-3">Email</th>
            <th class="px-6 py-3">Bisnis-loc</th>
            <th class="px-6 py-3">Year</th>
            <th class="px-6 py-3">Date-Log</th>
            <th class="px-6 py-3">Action</th>
          </tr>
        </thead>
        <tbody>
          {#each filteredObj as obj}
            <tr class="odd:bg-white bg-gray-900 even:bg-gray-50 even:bg-gray-200 text-gray-600">
              <th class="px-6 py-4">
                <input type="checkbox" on:change={() => toggleSelection(obj.Id)} /> 

              </th>
              <td class="px-6 py-4 font-medium">{obj.Name}</td>
              <td class="px-6 py-4">{obj.Id}</td>
              <td class="px-6 py-4">{obj.Role}</td>
              <td class="px-6 py-4">{obj.Email}</td>
              <td class="px-6 py-4">{obj.Bis_Loc}</td>
              <td class="px-6 py-4">{obj.Year}</td>
              <td class="px-6 py-4">{obj.Date_Log}</td>
              <td class="px-6 py-4 flex">
                <button on:click={() => { showDelete = obj.Name; showAlert = true; }}
                        class="font-medium text-white bg-red-500 rounded-full py-1 px-2">
                  Delete
                </button>
                <a href="/views-users" class="font-medium text-white bg-green-500 rounded-full py-1 px-2">
                  View
                </a>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  
    <!-- Alert konfirmasi -->
    {#if showAlert}
      <div class="fixed inset-0 backdrop-blur-sm flex items-center justify-center p-4">
        <div class="bg-gray-950 text-red-800 rounded-lg shadow-lg p-6 max-w-sm w-full">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center">
              <AlertTriangle class="w-8 h-8 text-red-500 mr-2" />
              <h2 class="text-xl font-semibold">Warning</h2>
            </div>
          </div>
          <p class="text-gray-400 mb-6">
            Kamu yakin ingin menghapus <span class="text-white">{showDelete}</span> dari database?
          </p>
          <div class="flex justify-end space-x-4">
            <button on:click={() => showAlert = false} class="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg">
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
  