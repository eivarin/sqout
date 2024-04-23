// import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Module } from '$lib';
// import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params }) => {
	const url = 'http://localhost:8080/modules/' + encodeURIComponent(params.moduleId);
	const resp = await fetch(url);
	const res: Module = await resp.json();
	return {
		module: res
	};
};