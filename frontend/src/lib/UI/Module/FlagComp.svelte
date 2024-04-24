<script lang="ts">
	import type { Flag } from '$lib';
	export let name: string;
	export let flag: Flag;
</script>

<div class="w-1/3 min-h-56 grow flex flex-col p-2 bg-base-300 rounded-lg shadow-lg">
	<div class="self-center max-w-[90%] text-wrap text-center pt-2">
		<p class="font-bold text-xl inline">{name}{flag.Prefix != '' ? ` (${flag.Prefix})` : ''}</p>
	</div>
    <div class="divider px-8 mt-2"/>
	<div class="flex w-full justify-around">
        <div><p class="font-bold inline">Type of Value:</p> {flag.Type}</div>
        <div><p class="font-bold inline">Required flag:</p> {flag.Required}</div>
	</div>
    <div class="flex w-full justify-around">
        <div><p class="font-bold inline">Default Value:</p> {flag.Default}</div>
        <div><p class="font-bold inline">Empty flag:</p> {flag.IsEmpty}</div>
	</div>
    <div class="divider px-8" />
    <div class="flex w-full justify-around text-center">
        <div><p class="font-bold inline">Description:</p> {flag.Description}</div>
    </div>
</div>
