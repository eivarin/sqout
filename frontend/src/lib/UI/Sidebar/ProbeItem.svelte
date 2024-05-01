<script lang="ts">
    import type { Probe } from "$lib/Types/Probe";
    export let probe: Probe;
    let state_class = probe.Alive ? "border-green-500" : "border-red-500";
</script>


<div class="h-16 w-full grid items-center">
    <a target="_self" href="/probe/{probe.Name}" class="round border-2 {state_class}  bg-info text-base-content w-56 justify-self-center h-full flex flex-col justify-around rounded-md shadow">
        <div class="">
            <div class="text-xl font-bold text-center">{probe.Name}</div>
            <div class="text-xs pl-4">{probe.Module}</div>
        </div>
    </a>
</div>
