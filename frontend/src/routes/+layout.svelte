<script>
	import '../app.css';
    import { Body } from '$lib/UI/Sidebar';
</script>

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link href="https://fonts.googleapis.com/css2?family=Comfortaa:wght@300..700&display=swap" rel="stylesheet">
    <title>sqout</title>
</svelte:head>

<style lang="postcss">
    :global(body) {
        font-family: 'Comfortaa', sans-serif;
    }
</style>

<div class="h-screen w-screen flex bg-base-100">
    <Body />
    <div class="max-h-screen max-w-[calc(100vw-18rem)] grow overflow-auto">
        <slot />
    </div>
</div>
