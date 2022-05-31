import {playbackConfigApi} from "@/api";

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
        loadPlaybackConfigs({commit}) {
            commit('setIsLoading', true);

            return playbackConfigApi.index().then(playbackConfigs => {
                commit('setPlaybackConfigs', playbackConfigs);
            }).finally(() => {
                commit('setIsLoading', false);
            })
        }
    }
}