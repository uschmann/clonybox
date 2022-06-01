
export default function parseSpotifyUrl(url) {
    const regEx = /^(?:spotify:|(?:https?:\/\/(?:open|play)\.spotify\.com\/))(?:embed)?\/?(album|track|playlist)(?::|\/)((?:[0-9a-zA-Z]){22})/;
    const match = url.match(regEx);
    if(match){
        return {
            type: match[1],
            id: match[2]
        }
    }

    return null;
}