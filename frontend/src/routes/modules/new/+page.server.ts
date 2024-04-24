import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';
// import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async () => {

};

export const actions = {
	default: async ({ request, locals }) => {
        const data = await request.formData();
        const name = data.get('name') as string;
        const branch = data.get('branch') as string;
		const commit = data.get('commit') as string;
		const resp = await fetch(`${locals.url}/modules/`, {
			method: 'POST',
			body: JSON.stringify({
				Name: name,
                Branch: branch,
                Commit: commit
			})
		});
        const link = "/module/" + await resp.json()
        throw redirect(302 ,link);
	}
};
