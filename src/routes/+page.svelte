<script>
  import "../app.css";
  import NavPanel from "$lib/component/Nav_Panel.svelte";
  import { CircleX, Search } from "@lucide/svelte";
  import DtaAPI from "$lib/data/Data_Api.js";
  import Footer from "$lib/component/Footer.svelte";

  const pajak = 10

  const data = DtaAPI;
  let showModal = false;

  let Value = 1;

  let checkoutData = [];

  let showData = {};

  let subtotal = 0;

$: if (showData && showData.price && Value) {
  const totalMakanan = showData.price * Value;
  const totalPajak = (pajak / 100) * totalMakanan;
  const totalDiscount = (showData.discount / 100) * totalMakanan;

  subtotal = Math.round(totalMakanan + totalPajak - totalDiscount);
}

  function Checkout() {
  const now = new Date();

  const tanggal = now.getDate();
  const bulan = now.getMonth() + 1;
  const tahun = now.getFullYear();
  const jam = now.getHours();
  const menit = now.getMinutes();

  let taggal = `${tanggal}/${bulan}/${tahun}`;
  let waktu = `${jam}:${menit}`;

  // Tambahkan data dan trigger reactivity
  checkoutData = [
    ...checkoutData,
    {
      data: showData,
      value: Value,
      taggal,
      waktu,
      subtotal,
    }
  ];

  console.log(checkoutData); // sekarang harus terisi

  showModal = false;
  Value = 1;
}


</script>

<NavPanel />

<div class="w-[100vw] h-[100vh] bg-gray-950 pt-15">
  <div class="w-full h-20 flex items-center p-2">
    <div class="w-65 h-10 rounded-md flex justify-between items-center">
      <input type="text" name="" id="" class="bg-white w-50 h-10 text-black rounded-md p-2" placeholder="Search Menu"/>
      <Search class="p-2 rounded-md bg-yellow-500 text-black" size={40}/>
    </div>

    <div class="drawer drawer-end z-20">
      <input id="my-drawer-4" type="checkbox" class="drawer-toggle" />
      <div class="drawer-content">
        <!-- Page content here -->
        <label for="my-drawer-4" class="drawer-button btn btn-primary">Open drawer</label>
      </div>
      <div class="drawer-side">
        <label for="my-drawer-4" aria-label="close sidebar" class="drawer-overlay"></label>
        <ul class="menu bg-base-200 text-base-content min-h-full w-100 p-4 mt-15 border">

          <div class="w-full flex">
            <h1 class="m-auto text-3xl">CheckOut</h1>
          </div>
          <div class="w-full h-[70vh] flex flex-wrap py-2 justify-start items-start overflow-scroll">
            {#if checkoutData.length > 0}
            {#each checkoutData as item}
              <div class="w-full h-15 bg-black flex justify-evenly items-center rounded-full m-1">
                <div class="w-10 h-10 bg-red-500 rounded-full bg-cover bg-center"
                  style={`background-image: url(${item.data.image_url})`}
                ></div>
                <div class="flex flex-col">
                  <h1 class="">{item.data.name}</h1>
                  <p class="text-white/50">{item.taggal} : {item.waktu}</p>
                </div>
                <h1 class="">Rp:{item.subtotal}</h1>
                <h1 class="">{item.value}</h1>
                <CircleX size={30} class="text-red-500"/>
              </div>
            {/each}
          {/if}
          </div>

        </ul>
      </div>
    </div>


  </div>

<!-- list menu -->
  <div class="w-full h-[85vh] flex flex-wrap justify-evenly items-center overflow-auto">

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

<!-- popup heighlight -->
{#if showModal}
  <div class="modal modal-open backdrop-blur" role="dialog">
    <div class="modal-box relative">
      <button class="absolute top-0 right-0 bg-transparent text-red-500" on:click={() => {showModal = false; Value = 1}}><CircleX size={35}/></button>
      <div class="w-full h-60 bg-red-600 rounded-md my-3 bg-cover bg-center"
      style={`background-image: url(${showData.image_url})`}
      ></div>
      <h3 class="font-bold text-3xl">{showData.name}</h3>
      <p class="py-4">
        Bum-bu <br>
        {showData.description}
      </p>
      <h1 class="">pajak :<span class="text-green-500">{pajak}%</span></h1>
      <h1 class="">discount :<span class="text-green-500">{showData.discount}%</span></h1>
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
<Footer></Footer>