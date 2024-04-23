<script lang="ts">
    import type { Module } from '$lib';
	export let data: {
        modules: Module[];
    };
    let ms = data.modules
</script>

<!-- <div class="w-full p-4">
    <a href="/modules/new" class="btn w-full">
        new
    </a>
</div> -->

<div class="flex gap-4 flex-wrap w-full h-full p-4">
	{#each ms as m}
        <div class="h-20 w-56 grid items-center">
            <a href="/probes/new/{m.Id}" class="round border-2 bg-info text-base-content w-full justify-self-center h-full flex flex-col justify-around rounded-md shadow px-4">
                <div class="">
                    <div class="text-xl font-bold text-center">{m.Exe.CommandName}</div>
                    <div class="text-xs pl-4">Type: {m.Exe.Description}</div>
                </div>
            </a>
        </div>
	{/each}
</div>