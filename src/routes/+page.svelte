<script>
  import "../app.css";
  import NavPanel from "$lib/component/Nav_Panel.svelte";
  import { CircleX } from "@lucide/svelte";
  import DtaAPI from "$lib/data/Data_Api.js";


  const data = DtaAPI;
  let showModal = false;

  let Value = 1;

  let checkoutData = [];

  let showData = {};

  // Menghitung subtotal secara reaktif
  $: subtotal = (showData.price * Value) + showData.pajak;

  function Checkout(){

    const now = new Date();

// Ambil bagian-bagiannya
const tanggal = now.getDate();           // Tanggal (1–31)
const bulan = now.getMonth() + 1;        // Bulan (0–11), jadi ditambah 1
const tahun = now.getFullYear();         // Tahun (misalnya 2025)
const jam = now.getHours();              // Jam (0–23)
const menit = now.getMinutes();          // Menit (0–59)


let taggal = `${tanggal}/${bulan}/${tahun}`
let waktu = `${jam}:${menit}`

    checkoutData.push({
      data: showData,
      value: Value,
      taggal,
      waktu,
      subtotal
    });
    
    console.log(checkoutData);

    showModal = false;
    Value = 1;
  }

</script>

<NavPanel />
<div class="w-[100vw] h-[100vh] bg-gray-950 pt-15">
  <div class="w-full h-20 bg-yellow-500">
    
  </div>
  <div class="w-full h-full flex flex-wrap justify-evenly items-center overflow-scroll">

    {#each data as item}
    <div
    class="card bg-base-100 w-58 md:w-96 h-40 shadow-sm m-1 bg-cover bg-center"
    style={`background-image: url(${item.image_url})`}
  >
  <div class="absolute inset-0 bg-black opacity-40"></div> <!-- overlay -->
      <div class="card-body z-10">
        <div class="">
          <h2 class="card-title text-2xl ">{item.name}</h2>
        </div>
        <p>{item.description}</p>
        <div class="card-actions justify-end">
          <button class="btn bg-green-500 text-black" on:click={() => {
            showModal = true;
            showData = item;
            }}>Rp: <span class="">{item.price}</span></button>
        </div>
      </div>
    </div>
    {/each}


{#if showModal}
  <div class="modal modal-open" role="dialog">
    <div class="modal-box relative">
      <button class="absolute top-0 right-0 bg-transparent text-red-500" on:click={() => {showModal = false; Value = 1}}><CircleX size={35}/></button>
      <h3 class="font-bold text-3xl">{showData.name}</h3>
      <p class="py-4">
        Bum-bu <br>
        {showData.description}
      </p>
      <h1 class="">pajak : Rp<span class="text-green-500">{showData.pajak}</span></h1>
      <h1 class="">product : Rp<span class="text-green-500">{showData.price}</span></h1>
      <div class="w-full h-20 flex justify-center items-center ">
        <div class="p-2 bg-gray-800 rounded-full flex justify-center items-center">
          <button type="button" class="py-2 px-4 bg-red-500 rounded-full" on:click={() => {
            if (Value > 1) Value -= 1;
          }}>-</button>
          <div class="mx-2  py-2 px-5"><input class="w-15 text-center" type="number" name="" id="" 
            bind:value={Value}
            ></div>
          <button type="button" class="py-2 px-4 bg-green-500 rounded-full" on:click={() => Value += 1}>+</button>
        </div>
      </div>
      <div class="w-full h-20 flex justify-center items-center">
        <button class="w-60 h-12 rounded-md bg-green-500 flex" on:click={Checkout}><h1 class="text-2xl m-auto">Checkout</h1></button>
      </div>
      <h1 class="text-2xl">subtotal : Rp<span class="text-green-500">{subtotal}</span></h1>
    </div>
  </div>
{/if}
  </div>
</div>