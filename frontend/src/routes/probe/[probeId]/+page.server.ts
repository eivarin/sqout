import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Module, Probe } from '$lib';
// import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, locals }) => {
	const urlProbe = `${locals.url}/probes/` + encodeURIComponent(params.probeId);
	const respProbe = await fetch(urlProbe);
	const probe: Probe = await respProbe.json();
	const urlModule = `${locals.url}/modules/` + encodeURIComponent(probe.Module);
	const respModule = await fetch(urlModule);
	const module: Module = await respModule.json();
	return {
		probe: probe,
		module: module
	};
};

export const actions = {
	Delete: async ({ params, locals }) => {
		const url = `${locals.url}/probes/` + encodeURIComponent(params.probeId);
		const resp = await fetch(url, {
			method: 'DELETE'
		});
		if (await resp.status == 200){
			throw redirect(302, "/");
		}
	}
};
