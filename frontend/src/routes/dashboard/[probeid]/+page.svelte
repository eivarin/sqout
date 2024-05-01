<script lang="ts">
    import { page } from '$app/stores';
    console.log($page.url)
    const grafanaUrl = `http://${$page.url.hostname}:3000/d/${$page.params.probeid}`
</script>

<div class="flex flex-col h-full">
    <div class="w-full flex justify-center text-4xl my-4">
        Grafana Dashboard
    </div>
    
    <iframe src={grafanaUrl} class="w-full h-full"></iframe>
</div>