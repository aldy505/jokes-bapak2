import type { Handle } from "@sveltejs/kit";

export const handle: Handle = ({ event, resolve }) => {
    return resolve(event, {
        transformPageChunk: ({ html }) => html.replace('%lang%', "en")
    });
}