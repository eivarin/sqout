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
	Reload: async ({ params, locals }) => {
		const resp = await fetch(`${locals.url}/modules/`, {
			method: 'PUT',
			body: JSON.stringify({
				Name: params.moduleId
			})
		});
		return await resp.status;
	},
	Checkout: async ({ params, request, locals }) => {
		const data = await request.formData();
		const branch = data.get('branch');
		const commit = data.get('commit');
		console.log(`branch: ${branch}, commit: ${commit}`);
		const resp = await fetch(`${locals.url}/modules/`, {
			method: 'PUT',
			body: JSON.stringify({
				Name: params.moduleId,
				Branch: branch,
				Commit: commit
			})
		});
		return await resp.status;
	},
	Delete: async ({ params, locals }) => {
		const url = `${locals.url}/modules/` + encodeURIComponent(params.moduleId);
		const resp = await fetch(url, {
			method: 'DELETE'
		});
		if (await resp.status == 200){
			throw redirect(302, "/modules/");
		}
	}
};
