import type { Probe } from '$lib';
import type { LayoutServerLoad } from './$types';


export const load: LayoutServerLoad = async ({locals}) => {
    const probesResp = await fetch(`${locals.url}/probes`);
    const probesJson: Probe[] = await probesResp.json();
    return {
        probes: probesJson
    };
}