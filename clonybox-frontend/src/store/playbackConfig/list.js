

export default {
    namespaced: true,
    state: {
        playbackConfigs: null,
        isLoading: false
    },
    mutations: {
        setPlaybackConfigs(state, playbackConfigs) {
            state.playbackConfigs = playbackConfigs;
        },
        setIsLoading(state, isLoading) {
            state.isLoading = isLoading;
        }
    },
    actions: {
        laodPlaybackConfigs() {
            // TODO: IMplement me....
        }
    }
}