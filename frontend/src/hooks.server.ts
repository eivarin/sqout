import type { Handle } from '@sveltejs/kit';



export const handle: Handle = async ({ event, resolve }) => {
	// Attach id to locals
	event.locals.url = process.env.BACKEND_URL ?? 'http://localhost:8080';
    const response = await resolve(event);
	return response;
};