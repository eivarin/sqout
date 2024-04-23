<script lang="ts">
	import type { Module } from '$lib';
	import { enhance } from '$app/forms';
	export let data: {
		module: Module;
	};
	let m = data.module;
</script>

<!-- DESCRIBE EVERY FIELD OF THE data.module -->

<p>{m.Id}</p>
<p>{m.Path}</p>
<p>{m.IsRepo}</p>
<p>{m.GitInfo.Branch}</p>
<p>{m.GitInfo.Commit}</p>
<p>{m.Exe.CommandName}</p>
<p>{m.Exe.Description}</p>
<p>{m.Exe.KeepAlive}</p>
<p>{m.Exe.FlagsOrder}</p>
{#each m.Exe.FlagsOrder as flagname}
	<p>{flagname}</p>
	<div class="pl-4">
		<p>{m.Exe.Flags[flagname].Description}</p>
		<p>{m.Exe.Flags[flagname].Type}</p>
		<p>{m.Exe.Flags[flagname].Required}</p>
		<p>{m.Exe.Flags[flagname].Prefix}</p>
	</div>
{/each}
<form method="POST" class="w-full">
	<button type="submit" class="btn" name="action" formaction="?/Reload">Reload</button>
	<button type="submit" class="btn" name="action" formaction="?/Delete">Delete</button>
	<div class="flex w-full gap-4 px-4">
		<label class="form-control grow">
			<div class="label">
				<span class="label-text">Branch</span>
			</div>
			<input type="text" name="branch" placeholder="Type here" class="input input-bordered w-full " value={m.GitInfo.Branch} />
		</label>
		<label class="form-control grow">
			<div class="label">
				<span class="label-text">Commit</span>
			</div>
			<input type="text" name="commit" placeholder="Type here" class="input input-bordered w-full" value={m.GitInfo.Commit} />
		</label>
		<button type="submit" class="btn self-end w-1/12" name="action" formaction="?/Checkout">Checkout</button>
	</div>
</form>
