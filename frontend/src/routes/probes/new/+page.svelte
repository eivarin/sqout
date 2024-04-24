<script lang="ts">
    import type { Module } from '$lib';
	import ModuleSmallComp from '$lib/UI/Module/ModuleSmallComp.svelte';
	export let data: {
        modules: Module[];
    };
    let ms = data.modules
</script>

<div class="flex flex-col w-full h-full p-4 gap-4">
    <p class="self-center text-4xl font-bold pb-2">Pick a module for the probe:</p>
    <div class="self-center flex gap-4 lg:w-[83rem] md:w-[31rem] flex-wrap">
        {#each ms as m}
            <ModuleSmallComp m={m} destinyBase="probes/new"/>
        {/each}
    </div>
</div>
