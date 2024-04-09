<script lang="ts">
    import ProbeItem from "./ProbeItem.svelte";
    import type { Probe } from "$lib/Types/Probe";
    //generate a list of random probes
    let probes: Probe[] = [
        {
            id: "1",
            name: "Probe 1",
            description: "This is a probe",
            running: true,
            type: "http",
        },
        {
            id: "2",
            name: "Probe 2",
            description: "This is a probe",
            running: false,
            type: "http",
        },
        {
            id: "3",
            name: "Probe 3",
            description: "This is a probe",
            running: true,
            type: "http",
        },
        {
            id: "4",
            name: "Probe 4",
            description: "This is a probe",
            running: false,
            type: "http",
        },
        {
            id: "5",
            name: "Probe 5",
            description: "This is a probe",
            running: true,
            type: "http",
        },
        {
            id: "6",
            name: "Probe 6",
            description: "This is a probe",
            running: false,
            type: "http",
        },
        {
            id: "7",
            name: "Probe 7",
            description: "This is a probe",
            running: true,
            type: "http",
        },
        {
            id: "8",
            name: "Probe 8",
            description: "This is a probe",
            running: false,
            type: "http",
        },
        {
            id: "9",
            name: "Probe 9",
            description: "This is a probe",
            running: true,
            type: "http",
        },
        {
            id: "10",
            name: "Probe 10",
            description: "This is a probe",
            running: false,
            type: "http",
        },
    ]
</script>


<div class="flex flex-col px-3 gap-3">
    {#each probes as probe}
        <ProbeItem probe={probe}/>
    {/each}
</div>