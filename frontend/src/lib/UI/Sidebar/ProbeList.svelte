<script lang="ts">
    import ProbeItem from "./ProbeItem.svelte";
    import type { Probe } from "$lib/Types/Probe";
    //generate a list of random probes
    export let probes: Probe[]
</script>


<div class="flex flex-col px-3 gap-3">
    {#each probes as probe}
        <ProbeItem probe={probe}/>
    {/each}
</div>