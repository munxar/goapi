// src/routes/about.server.ts

// you can make some calls to an api, or load markdown from the filesystem to build a documentation site or the lile. 
// this will happen at build time, so don't call your own go api!
export async function load() {
	return {
		message: 'Hello from the server at build time'
	};
}
