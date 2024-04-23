import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
		host: "0.0.0.0",
		hmr: {
		  clientPort: 5000,
		},
		port: 5000, 
		watch: {
		  usePolling: true,
		},
		// https: true
	},
	plugins: [sveltekit()],
});