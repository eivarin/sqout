<script lang="ts">
	import type { Module } from '$lib';
	export let data: {
		module: Module;
	};
	let m = data.module;
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
		name: '',
		description: '',
		heartbit: 0,
		module: m.Id,
		options: {}
	};
	formValues.options = Object.fromEntries(flags.map(({ name }) => [name, m.Exe.Flags[name].Default ?? '']));
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
			required
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
            required
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
			disabled={m.Exe.KeepAlive}
            required={!m.Exe.KeepAlive}
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
					<span class="label-text">{name}{f.Prefix != '' ? ` (${f.Prefix})` : ''}{f.Default != '' ? ` (Default is: ${f.Default})` : ''}:</span>
					<span class="label-text text-xs">{f.Description}</span>
				</div>
				{#if f.IsEmpty}
				<label class="label input cursor-pointer input-bordered">
					<span class="-webkit-input-placeholder">Toggle</span> 
					<input type="checkbox" class="checkbox" name={'options:' + name} bind:checked={formValues.options[name]}/>
				</label>
				{:else}
				<input
					type="text"
					name={'options:' + name}
					placeholder="Type here"
					class="input input-bordered w-full"
					bind:value={formValues.options[name]}
					required={f.Required}
				/>
				{/if}
			</label>
		{/each}
	</div>
	<input type="hidden" name="module" value={m.Id} />
	<label class="form-control grow px-4">
		<div class="label">
			<span class="label-text">Preview of the command:</span>
		</div>
		<textarea name="" id="" class="textarea input-bordered" disabled>
{m.Exe.CommandName} {Object.entries(formValues.options)
    .map(([key, value]) => {
		let f = m.Exe.Flags[key];
        if (f.Required && value == '')
            return `${f.Prefix} <${key}>`;
		else if (f.IsEmpty && value!="") return `${f.Prefix}`;
        else if (value == '') return ``;
        return `${f.Prefix} ${value}`;
    })
    .join(' ')}
		</textarea>
	</label>
	<div class="grow p-4">
		<button type="submit" class="btn btn-primary w-full">Create</button>
	</div>
</form>
