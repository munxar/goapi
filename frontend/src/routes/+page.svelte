<!-- src/routes/+page.svelte -->
<script lang="ts">
	const apiUrl = (path: string) => `${import.meta.env.VITE_API_URL || ''}${path}`;

	const getVersion = async () => {
		const url = apiUrl('/');
		const res = await fetch(url);
		if (!res.ok) {
			throw `Error while fetching data from ${url} (${res.status} ${res.statusText}).`;
		}
		const { version } = await res.json();
		return version;
	};
</script>

<h1>Home</h1>
{#await getVersion()}
	loading...
{:then version}
	Version from Server: {version}
{:catch err}
	{err}
{/await}
