<script lang="ts">
    import { page } from '$app/stores';
    $: pId = $page.params.probeId;
</script>

{pId}