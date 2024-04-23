import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Module, Probe } from '$lib';
// import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params }) => {
	const urlProbe = 'http://localhost:8080/probes/' + encodeURIComponent(params.probeId);
	const respProbe = await fetch(urlProbe);
	const probe: Probe = await respProbe.json();
	const urlModule = 'http://localhost:8080/modules/' + encodeURIComponent(probe.Module);
	const respModule = await fetch(urlModule);
	const module: Module = await respModule.json();
	return {
		probe: probe,
		module: module
	};
};

export const actions = {
	Delete: async ({ params }) => {
		const url = 'http://localhost:8080/probes/' + encodeURIComponent(params.probeId);
		const resp = await fetch(url, {
			method: 'DELETE'
		});
		if (await resp.status == 200){
			throw redirect(302, "/");
		}
	}
};
