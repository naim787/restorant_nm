<script>
    import NavAdmin from "../../components/NavAdmin.svelte";

    let Menu = [
        {
            name: "Sate icikwir",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn1.jpg",
            price: 15000,
        },
        {
            name: "Sate icikwir",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn1.jpg",
            price: 15000,
        },
        {
            name: "Tuna bakar",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn2.jpg",
            price: 15000,
        },
        {
            name: "Tuna bakar",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn2.jpg",
            price: 15000,
        },
        {
            name: "Tuna bakar",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn2.jpg",
            price: 15000,
        },
        {
            name: "Tindarug Bakar",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn3.jpg",
            price: 15000,
        },
        {
            name: "Tindarug Bakar",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn3.jpg",
            price: 15000,
        },
        {
            name: "Tindarug Bakar",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn3.jpg",
            price: 15000,
        },
        {
            name: "Tindarug Bakar",
            id: "330",
            ket: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur commodi numquam incidunt eos qui lorem:15",
            url: "./mkn3.jpg",
            price: 15000,
        },

    ]

    let searchItem = "";
    $: filteredObj = searchItem 
      ? Menu.filter(item => item.name.toLowerCase().includes(searchItem.toLowerCase())) 
      : Menu;
</script>

<NavAdmin title={"Menu"}>
    <div class="w-full flex justify-between">
        <input type="text" name="" id="" bind:value={searchItem} class="bg-white p-2 shadow rounded-md mb-2 border text-black" placeholder="Search Item"/>
    </div>
    <div class="w-[88vw] h-[89vh] flex justify-evenly items-start flex-wrap gap-2 overflow-scroll">

        {#each filteredObj as D}
            <div class="w-70 h-90 rounded bg-gray-900 overflow-hidden relative">
                <div style={`background-image: url('${D.url}')`} class="w-full h-1/2 bg-cover bg-center">
                </div>
                <div class="flex flex-col p-2 text-white">
                    <h1 class="text-2xl font-bold">{D.name}</h1>
                    <p class="text-1xl font-thin text-gray-300">{D.ket}</p>
                    <div class="w-full flex justify-between items-center">
                        <h1 class="font-bold text-2xl">Rp : <span class="text-green-500">15000</span></h1>   
                        <button class="py-2 px-3 bg-blue-500 rounded-md">Pesan</button>
                    </div>
                </div>
            </div>
        {/each}

    </div>
</NavAdmin>