import type { Probe } from '$lib';
import type { LayoutServerLoad } from './$types';


export const load: LayoutServerLoad = async () => {
    const probesResp = await fetch('http://localhost:8080/probes');
    const probesJson: Probe[] = await probesResp.json();
    return {
        probes: probesJson
    };
}