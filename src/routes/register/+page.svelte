<script>
    import { Eye, EyeOff } from "@lucide/svelte";
    import "../../app.css";
  import { goto } from "$app/navigation";
  
    let pass = false;
    let passValue = "";
    let username = "";
    let email = "";
    let bisnis_location = "";
    let birthDay = "";
    let nik = "";
  
    async function handleSubmit(event) {
      event.preventDefault(); // Mencegah reload halaman
  
      const formData = {
        Name,
        Id: "1255",
        Role: "nn",
        Email,
        bis_loc: bisnis_location,
        Year: birthDay,
        DateLog: "2025-03-18",
        Nik,
        password: passValue
      };
  
      try {
        const response = await fetch("/register/create-users", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(formData),
        });
  
        const result = await response.json();
        if (result.success) {
          alert("User berhasil dibuat!");
          goto("/")
        } else {
          alert("Gagal: " + result.error);
        }
      } catch (error) {
        console.error("Fetch error:", error);
      }
    }
  </script>
  
  <div class="w-[100vw] h-[100vh] flex justify-center items-center flex-col bg-blue-50">
      <form on:submit={handleSubmit} class="w-1/4 h-2/3 bg-white rounded-md shadow-xl flex flex-col justify-start items-center py-2">
          <div class="w-full h-auto flex">
              <h1 class="text-2xl m-auto font-bold">Register</h1>
          </div>
          <div class="">
              <p class="text-gray-700">* User Name</p>
              <input type="text" class="bg-gray-200 w-75 h-10 rounded-md p-2 mb-2" bind:value={username} placeholder="..." />
  
              <p class="text-gray-700">* Email</p>
              <input type="email" class="bg-gray-200 w-75 h-10 rounded-md p-2 mb-2" bind:value={email} placeholder="..." />
  
              <p class="text-gray-700">* Brith Day</p>
              <input type="date" class="bg-gray-200 w-75 h-10 rounded-md p-2 mb-2" bind:value={birthDay} />

              <p class="text-gray-700">* Bisnis Location</p>
             <select name="bis_loc" id="" bind:value={bisnis_location}>
                <option value="komo">komo</option>
                <option value="paniki">paniki</option>
                <option value="wasian">wasian</option>
             </select>
  
              <p class="text-gray-700">* Nik</p>
              <input type="text" class="bg-gray-200 w-75 h-10 rounded-md p-2 mb-2" bind:value={nik} placeholder="..." />
  
              <p class="text-gray-700">* Password</p>
              {#if pass}
                  <input type="password" class="bg-gray-200 w-75 h-10 rounded-md p-2" bind:value={passValue} placeholder="..." />
              {:else}
                  <input type="text" class="bg-gray-200 w-75 h-10 rounded-md p-2" bind:value={passValue} placeholder="..." />
              {/if}
              <button type="button" on:click={() => pass = !pass}>
                  {#if pass}
                    <EyeOff size={20} class="m-2"/>
                  {:else}
                    <Eye size={20} class="m-2"/>
                  {/if}
              </button>
          </div>
          <div class="mt-3">
              <button type="submit" class="py-2 px-15 bg-blue-500 rounded-md text-white">Create</button>
          </div>
      </form>
  </div>
  