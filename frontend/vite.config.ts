import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
		host: "0.0.0.0",
		hmr: {
		  clientPort: 8080,
		},
		port: 8080, 
		watch: {
		  usePolling: true,
		},
		// https: true
	},
	plugins: [sveltekit()],
});