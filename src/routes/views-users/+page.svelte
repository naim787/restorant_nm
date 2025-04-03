<script>
    import { Ellipsis, UtensilsCrossed } from "@lucide/svelte";
    import "../../app.css";
</script>

<div class="w-[100vw] h-[100vh] flex justify-center items-center bg-blue-50">
    <div class="w-full max-w-sm bg-white border rounded-md shadow-sm dark:bg-gray-800 dark:border-gray-700 relative overflow-hidden">
        <div class="flex justify-end px-4 pt-4 relative">
            <button 
                class="text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg p-1.5 ">
                <Ellipsis class="z-10" />
            </button>
        </div>
        
        <div class="flex flex-col items-center pb-10">
            <div class="w-full h-32 absolute top-0 bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500"></div>
            <div class="z-10 w-24 h-24 mb-3 rounded-full shadow-lg flex justify-center items-center bg-green-500">
                <UtensilsCrossed size={40}/>
            </div>
            <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">Naim Abdullah</h5>
            <div class="flex flex-col">
                <span class="text-sm text-gray-400">id: 155</span>
                <span class="text-sm text-gray-400">Role: Cleaning Service</span>
                <span class="text-sm text-gray-400">Email: naim@gmail.com</span>
                <span class="text-sm text-gray-400">Year: 2004-09-18</span>
                <span class="text-sm text-gray-400">Date-Log: 2024-09-18</span>
                <span class="text-sm text-gray-400">Nik: 18393782598472</span>
                <span class="text-sm text-gray-400">Password: 18393782598472</span>
            </div>
            <div class="mt-9">
                <a href="/" class="py-2 px-4 rounded-md bg-green-500 mx-4">Edit</a>
                <a href="/" class="py-2 px-4 rounded-md bg-red-500">Delete</a>
            </div>
        </div>
    </div>
</div>

