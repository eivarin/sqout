// import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Module } from '$lib';
// import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({locals}) => {
    const modules = await fetch(`${locals.url}/modules/`);
    const modulesJson: Module[] = await modules.json();
    return {
        modules: modulesJson
    };
};