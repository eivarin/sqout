<script lang="ts">
    import { page } from '$app/stores';
	import type { Module, Probe } from '$lib';
	import GrafanaIcon from '$lib/UI/Icons/GrafanaIcon.svelte';
    export let data: { probe: Probe, module: Module };
    const p = data.probe
	const m = data.module
    let flags = m.Exe.FlagsOrder.map((name) => ({ name, f: m.Exe.Flags[name] }));
	let formValues: {
		name: string;
		description: string;
		heartbit: number;
		module: string;
		options: {
			[key: string]: string;
		};
	} = {
		name: p.Name ?? '',
		description: p.Description ?? '',
		heartbit: p.HeartbitInterval ?? 0,
		module: p.Module ?? '',
		options: {}
	};
	formValues.options = Object.fromEntries(flags.map(({ name }) => [name, p.Options[name] ?? '']));
</script>

<form method="POST" class="w-full">
	<label class="form-control grow px-4">
		<div class="label">
			<span class="label-text">Name of the Probe:</span>
		</div>
		<input
			type="text"
			name="name"
			placeholder="Type here"
			class="input input-bordered w-full"
			value={formValues.name}
			readonly
		/>
	</label>
	<label class="form-control grow px-4">
		<div class="label">
			<span class="label-text">Description:</span>
		</div>
		<input
			type="text"
			name="description"
			placeholder="Type here"
			class="input input-bordered w-full"
			value={formValues.description}
            readonly
		/>
	</label>
	<label class="form-control grow px-4">
		<div class="label">
			<span class="label-text"
				>Heartbit of the probe{m.Exe.KeepAlive ? '(Disabled due to keep alive)' : ''}:</span
			>
		</div>
		<input
			type="number"
			name="heartbit"
			placeholder="Type here"
			class="input input-bordered w-full"
			value={formValues.heartbit}
			readonly
		/>
	</label>
	<div class="grow px-4">
		<div class="label">
			<span class="label-text">Options:</span>
		</div>
	</div>
	<div class="px-6">
		{#each flags as { name, f }}
			<label class="form-control grow px-4">
				<div class="label">
					<span class="label-text">{name}{f.Prefix != '' ? ` (${f.Prefix})` : ''}:</span>
					<span class="label-text text-xs">{f.Description}</span>
				</div>
				<input
					type="text"
					name={'options:' + name}
					placeholder="Type here"
					class="input input-bordered w-full"
					bind:value={formValues.options[name]}
					readonly
				/>
			</label>
		{/each}
	</div>
	<input type="hidden" name="module" value={m.Id} />
	<label class="form-control grow px-4">
		<div class="label">
			<span class="label-text">Preview of the command:</span>
		</div>
		<textarea name="" id="" class="textarea input-bordered" readonly>
{m.Exe.CommandName} {Object.entries(formValues.options)
    .map(([key, value]) => {
        if (m.Exe.Flags[key].Required && value == '')
            return `${m.Exe.Flags[key].Prefix} <${key}>`;
        else if (value == '') return ``;
        return `${m.Exe.Flags[key].Prefix} ${value}`;
    })
    .join(' ')}
		</textarea>
	</label>
	<div class="grow p-4 flex gap-4">
		<!-- <button type="submit" class="btn btn-info grow">Edit</button> -->
		<button type="submit" formaction="?/Delete" class="btn btn-error grow">Delete</button>
	</div>
	<div class="grow pb-4 px-4 flex gap-4">
		<!-- <button type="submit" class="btn btn-info grow">Edit</button> -->
		<a href="/dashboard/{$page.params.probeId}" class="grow bg-[#F2F4F9] text-black flex justify-center items-center gap-1 h-[46px] rounded-lg">
			<GrafanaIcon />
			Grafana
		</a>
	</div>
</form>