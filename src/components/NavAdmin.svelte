<script>
    import "../app.css"
    export let title
    import { Building, Home, FileText, Folder, BarChart2, Settings, Users, User, Plus, MoreVertical,UserCog,TableProperties,PackageCheck, Armchair, CirclePercent, OctagonX, FishSymbol} from '@lucide/svelte';

    let DarkMode = false;
    $: wrapMode = `w-14 h-6 rounded-full flex justify-start items-center p-1 shadow relative  ${DarkMode ? "bg-gray-800" : "bg-gray-400"}`
    $: oMode = `w-5 h-5 rounded-full absolute ${DarkMode ? "right-1 bg-white" : "left-1 bg-black"}`
</script>

<div class="flex w-screen h-screen text-gray-700">
    <!-- Sidebar -->
    <div class="flex flex-col items-start w-30 pb-4 overflow-auto border-r border-gray-300 p-1"
    class:bg-gray-950={DarkMode}
    class:text-white={DarkMode}
    >
        <a class="flex items-center justify-center flex-shrink-0 w-full h-16 rounded-tl-md " href="/"
        class:bg-green-500={DarkMode}
        class:bg-gray-300={!DarkMode}
        >
            <Building class="w-8 h-8" />
        </a>
        <a class="flex items-center justify-start flex-shrink-0 w-full h-10 mt-4 rounded hover:bg-gray-300 hover:text-black" href="/">
            <Users class="w-5 h-5 mx-1" />
            <span>Staff</span>
        </a>
        <a class="flex items-center justify-start flex-shrink-0 w-full h-10 mt-4 rounded hover:bg-gray-300 hover:text-black" href="/menu">
            <TableProperties class="w-5 h-5 mx-1"/>
            <span>Menu</span>
        </a>
        <a class="flex items-center justify-start flex-shrink-0 w-full h-10 mt-4 rounded hover:bg-gray-300 hover:text-black" href="/warehouse">
            <PackageCheck class="w-5 h-5 mx-1"/>
            <span>Warehouse</span>
        </a>
        <a class="flex items-center justify-start flex-shrink-0 w-full h-10 mt-4 rounded hover:bg-gray-300 hover:text-black" href="/">
            <FishSymbol  class="w-5 h-5 mx-1"/>
            <span>Meat</span>
        </a>
        <a class="flex items-center justify-start flex-shrink-0 w-full h-10 mt-4 rounded hover:bg-gray-300 hover:text-black" href="/tabel">
            <Armchair class="w-5 h-5 mx-1"  />
            <span>Tabel</span>
        </a>
        <a class="flex items-center justify-start flex-shrink-0 w-full h-10 mt-4 rounded hover:bg-gray-300 hover:text-black" href="/">
            <CirclePercent class="w-5 h-5 mx-1" />
            <span>Discount</span>
        </a>
        <a class="flex items-center justify-start flex-shrink-0 w-full h-10 mt-4 rounded hover:bg-gray-300 hover:text-black" href="/">
            <OctagonX class="w-5 h-5 mx-1"/>
            <span>Reservasi</span>
        </a>

        <a class="flex items-center justify-center flex-shrink-0 w-10 h-10 mt-4 mt-auto rounded hover:bg-gray-300" href="/">
            <User class="w-5 h-5 mx-1" />
        </a>
    </div>


    <!-- Main Content -->
    <div class="flex flex-col flex-grow"
    class:bg-gray-950={DarkMode}
    class:text-white={DarkMode}
    >
        <div class="flex items-center flex-shrink-0 h-16 px-8 border-b border-gray-300">
            <h1 class="text-lg font-medium">{title}</h1>
            <button class="flex items-center justify-center h-10 px-4 ml-auto text-sm font-medium rounded hover:bg-gray-300 hover:text-black">
                Action 1
            </button>
            <button class="flex items-center justify-center h-10 px-4 ml-2 text-sm font-medium bg-gray-200 rounded hover:bg-gray-300 text-black">
                Action 2
            </button>
            <button class="relative ml-2 text-sm focus:outline-none group">
                <div class="flex items-center justify-between w-10 h-10 rounded hover:bg-gray-300 hover:text-black">
                    <MoreVertical class="w-5 h-5 mx-auto" />
                </div>
                <div class="absolute right-0 flex-col items-start hidden w-40 pb-1 bg-white border border-gray-300 shadow-lg group-focus:flex z-10 p-2 text-black">
                    <a class="w-full px-4 py-2 text-left hover:bg-gray-300" href="/">Menu Item 1</a>
                    <a class="w-full px-4 py-2 text-left hover:bg-gray-300" href="/">Menu Item 2</a>
                    <a class="w-full px-4 py-2 text-left hover:bg-gray-300" href="/">Menu Item 3</a>
                    <!-- <div on:click={() => DarkMode = !DarkMode} class={wrapMode}>
                        <div class={oMode}></div>
                    </div> -->
                </div>
            </button>
        </div>

        <div class="flex-grow p-6 overflow-auto"
        class:bg-blue-50={!DarkMode}
        class:bg-gray-800={DarkMode}
        >
            <div class="rounded-lg">
                <!-- ini isi content -->
                 <slot></slot>
            </div>
        </div>
    </div>
</div>

<a class="fixed flex items-center justify-center h-8 pr-2 pl-1 bg-blue-600 rounded-full bottom-0 right-0 mr-4 mb-4 shadow-lg text-blue-100 hover:bg-blue-600" href="https://twitter.com/lofiui" target="_top">
    <div class="flex items-center justify-center h-6 w-6 bg-blue-500 rounded-full">
        <User class="w-4 h-4 text-white" />
    </div>
    <span class="text-sm ml-1 leading-none">@lofiui</span>
</a>
