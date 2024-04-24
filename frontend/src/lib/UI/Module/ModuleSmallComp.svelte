<script lang="ts">
    import type { Module } from '$lib';
    export let m: Module;
    export let destinyBase: string = "module";
</script>

<div class="h-28 lg:w-80 md:w-60 grid items-center">
    <a href="/{destinyBase}/{m.Id}" class="round border-2 bg-info text-base-content w-full justify-self-center h-full flex flex-col justify-center gap-2 rounded-md shadow px-4">
        <div class="text-xl font-bold text-center">{m.Exe.CommandName}</div>
        <div class="text-xs pl-4 line-clamp-2">Type: {m.Exe.Description}</div>
    </a>
</div>