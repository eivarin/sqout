import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Module } from '$lib';
// import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, locals }) => {
	const url = `${locals.url}/modules/` + encodeURIComponent(params.moduleId);
	const resp = await fetch(url);
	const res: Module = await resp.json();
	return {
		module: res
	};
};

export const actions = {
	default: async ({ params, request, locals }) => {
		const fData = await request.formData();
		const options: { [key: string]: string} = {};
		for (const [key, value] of fData.entries()) {
			if (key.startsWith("options:")){
				options[key.substring(8)] = value as string;
			}
		}
		const name = fData.get('name') as string;
		const body = {
			Name: name,
			Description: fData.get('description') as string,
			Options: options,
			HeartbitInterval: Number(fData.get('heartbit') as string),
			ModuleName: params.moduleId
		};
		console.log(body);
		const url = `${locals.url}/probes/`;
		const resp = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(body)
		});
		if (await resp.status == 200){
			throw redirect(302, "/probe/" + name);
		}
	}
};