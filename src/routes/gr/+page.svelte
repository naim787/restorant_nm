<script>
    import { tweened } from "svelte/motion";
    import { cubicInOut } from "svelte/easing";
    import "../../app.css"

    let progress = tweened(0, { duration: 3000, easing: cubicInOut });
  
    function animateGradient() {
      progress.set(1, { duration: 3000 }).then(() => {
        progress.set(0, { duration: 3000 }).then(animateGradient);
      });
    }
  
    animateGradient();

    $: gradientClass = $progress < 0.5
    ? "from-cyan-500 to-blue-500"
    : "from-blue-500 to-cyan-500";
  </script>
  
  <h1 
    class="text-5xl font-bold text-center bg-clip-text text-transparent transition-all duration-1000 bg-gradient-to-r {gradientClass}"
   
  >
    Animated Gradient Text
  </h1>
  