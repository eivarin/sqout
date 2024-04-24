<script lang="ts">
	import type { Module } from '$lib';
	import { enhance } from '$app/forms';
	import FlagComp from '$lib/UI/Module/FlagComp.svelte';
	export let data: {
		module: Module;
	};
	let m = data.module;
</script>

<!-- DESCRIBE EVERY FIELD OF THE data.module -->
<div class="w-full px-4 flex flex-col py-3">
	<p class="self-center text-4xl font-bold pb-2">{m.Id}</p>
	<div class="divider px-8 text-3xl">Description:</div>
	<textarea class="textarea textarea-bordered my-2" readonly>{m.Exe.Description}</textarea>
	<div class="divider px-8 text-3xl">Flags Order:</div>
	<div class="breadcrumbs input self-center">
		<ul>
			<li>
				<div class="px-2 rounded bg-primary text-primary-content">{m.Exe.CommandName}</div>
			</li>
			{#each m.Exe.FlagsOrder as f}
				<li>
					<div class="px-2 rounded bg-primary text-primary-content">{f}</div>
				</li>
			{/each}
		</ul>
	</div>
	<div class="divider px-8 text-3xl">Flags:</div>
	<div class="flex flex-wrap gap-4 justify-around p-4">
		{#each m.Exe.FlagsOrder as flagname}
			<FlagComp name={flagname} flag={m.Exe.Flags[flagname]} />
		{/each}
	</div>
	<form method="POST" class="w-full">
		{#if m.IsRepo}
			<div class="divider px-8 text-3xl">Git Info:</div>
			<div class="flex w-full gap-4 pb-4">
				<label class="form-control grow">
					<div class="label">
						<span class="label-text">Branch</span>
					</div>
					<input
					type="text"
					name="branch"
					placeholder="Type here"
					class="input input-bordered w-full"
					value={m.GitInfo.Branch}
					/>
				</label>
				<label class="form-control grow">
					<div class="label">
						<span class="label-text">Commit</span>
					</div>
					<input
					type="text"
					name="commit"
					placeholder="Type here"
					class="input input-bordered w-full"
					value={m.GitInfo.Commit}
					/>
				</label>
				<button
				type="submit"
				class="btn btn-info self-end w-1/12"
				name="action"
				formaction="?/Checkout">Checkout</button
				>
			</div>
		{/if}
		<div class="divider px-8 text-3xl">Alter:</div>
		<div class="grow py-4 flex gap-4">
			<button type="submit" class="btn btn-info grow" name="action" formaction="?/Reload"
			>Reload</button
			>
			<button type="submit" class="btn btn-error grow" name="action" formaction="?/Delete"
				>Delete</button
			>
		</div>
	</form>
</div>
