import {playbackConfigApi} from "@/api";
import notification from "@/utils/notification";

export default {
    namespaced: true,
    state: {
        playbackConfigs: null,
        isLoading: false,
    },
    mutations: {
        setPlaybackConfigs(state, playbackConfigs) {
            state.playbackConfigs = playbackConfigs;
        },
        setIsLoading(state, isLoading) {
            state.isLoading = isLoading;
        },
        removePlaybackConfig(state, playbackConfig) {
            if(!state.playbackConfigs) {
                return;
            }
            state.playbackConfigs = state.playbackConfigs.filter(p => p.id != playbackConfig.id);
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
        },
        deletePlaybackConfig({commit}, playbackConfig) {
            return playbackConfigApi.delete(playbackConfig.id).then(playbackConfig => {
                commit('removePlaybackConfig', playbackConfig);
                notification.success('The Playback-config was deleted');
            })
        }
    }
}