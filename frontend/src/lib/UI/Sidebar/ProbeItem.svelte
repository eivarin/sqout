<script lang="ts">
    import type { Probe } from "$lib/Types/Probe";
    export let probe: Probe;
    let state_class = probe.running ? "border-green-500" : "border-red-500";
</script>


<div class="h-16 w-full grid items-center">
    <a href="/probe/{probe.id}" class="round border-2 {state_class}  bg-info text-base-content w-56 justify-self-center h-full flex flex-col justify-around rounded-md shadow">
        <div class="">
            <div class="text-xl font-bold text-center">{probe.name}</div>
            <div class="text-sm pl-4">Type: {probe.type}</div>
        </div>
    </a>
</div>
