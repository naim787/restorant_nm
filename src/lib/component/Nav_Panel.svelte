<script>
  import "../../app.css";
  import { AlignJustify, Moon, Sun } from "@lucide/svelte";
</script>


<div class="drawer fixed z-50 bg-coffe" data-theme="coffe" >
    <input id="my-drawer-3" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content flex flex-col">
      <!-- Navbar -->
      <div class="navbar bg-base-300 w-full">
        <div class="flex-none">
          <label for="my-drawer-3" aria-label="open sidebar" class="btn btn-square btn-ghost">
            <AlignJustify />
          </label>
        </div>
        <div class="mx-2 flex-1 px-2">Navbar Title</div>
        <div class="hidden flex-none lg:block">
          <ul class="menu menu-horizontal">
            <!-- Navbar menu content here -->
            <label class="swap swap-rotate">
              <!-- this hidden checkbox controls the state -->
              <input type="checkbox" class="theme-controller" value="synthwave" />
            
              <!-- sun icon -->
              <Sun class="swap-off" />
            
              <!-- moon icon -->
              <Moon class="swap-on" />
            </label>
            <li><a href="/">Navbar Item 1</a></li>
            <li><a href="/">Navbar Item 2</a></li>
          </ul>
        </div>
      </div>
      <!-- Page content here -->
      <!-- Content -->
    </div>
    <div class="drawer-side">
      <label for="my-drawer-3" aria-label="close sidebar" class="drawer-overlay"></label>
      <ul class="menu bg-base-200 min-h-full w-80 p-4">
        <!-- Sidebar content here -->
        <li><a href="/">Sidebar Item 1</a></li>
        <li><a href="/">Sidebar Item 2</a></li>
      </ul>
    </div>
  </div>