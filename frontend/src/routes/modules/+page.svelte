<script lang="ts">
    import type { Module } from '$lib';
	import ModuleSmallComp from '$lib/UI/Module/ModuleSmallComp.svelte';
	export let data: {
        modules: Module[];
    };
    let ms = data.modules
</script>

<div class="flex flex-col w-full h-full p-4 gap-4">
    <div class="flex gap-4 flex-wrap w-full lg:w-[83rem] md:w-[31rem] grow self-center">
        {#each ms as m}
            <ModuleSmallComp m={m} />
        {/each}
    </div>
    <div class="w-full">
        <a href="/modules/new" class="btn btn-primary w-full">
            Add New Module
        </a>
    </div>
</div>
