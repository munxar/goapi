// frontend/svelte.config.js
import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		paths: {
			base: '/admin'
		},
		trailingSlash: 'ignore',
		adapter: adapter(),
		prerender: {
			default: true
		}
	}
};

export default config;