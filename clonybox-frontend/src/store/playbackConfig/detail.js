import {playbackConfigApi} from "@/api";
import playbackConfig from "@/store/playbackConfig/index";
import parseSPotifyUrl from "@/utils/parseSpotifyUrl";
import parseSpotifyUrl from "@/utils/parseSpotifyUrl";

export default {
    namespaced: true,
    state: {
        playbackConfig: null,
        isSaving: true,
        spotifyUrl: "",
        isLoading: false,
        isAssigningFromUrl: false,
        isAssigningFromItem: false
    },
    mutations: {
        setIsAssigningFromUrl(state, isAssigning) {
            state.isAssigningFromUrl = isAssigning;
        },
        setIsAssigningFromItem(state, isAssigning) {
            state.isAssigningFromItem = isAssigning;
        },
        setIsLoading(state, isLoading) {
            state.isLoading = isLoading;
        },
        setPlaybackConfig(state, playbackConfig) {
            state.playbackConfig = playbackConfig;
        },
        updateSpotifyUrl(state, url) {
            state.spotifyUrl = url;
        }
    },
    actions: {
        loadPlaybackConfig({commit}, id) {
            commit('setIsLoading', true);

            return playbackConfigApi.show(id).then(playbackConfig => {
                commit('setPlaybackConfig', playbackConfig);
            }).finally(() => {
                commit('setIsLoading', false);
            });
        },
        assignFromSpotifyItem({commit, state}, item) {
            const playbackConfig = {
                ...state.playbackConfig, ...{
                    spotify_id: item.item.id,
                    type: item.type
                }
            }

            commit('setIsAssigningFromItem', true);
            return playbackConfigApi.update(playbackConfig).then(playbackConfig => {
                commit('setPlaybackConfig', playbackConfig);
            }).finally(() => {
                commit('setIsAssigningFromItem', false);
            });
        },
        assignFromSpotifyUrl({commit, state}) {
            const data = parseSpotifyUrl(state.spotifyUrl)
            const playbackConfig = {
                ...state.playbackConfig, ...{
                    spotify_id: data.id,
                    type: data.type
                }
            }

            commit('setIsAssigningFromUrl', true);
            return playbackConfigApi.update(playbackConfig).then(playbackConfig => {
                commit('setPlaybackConfig', playbackConfig);
            }).finally(() => {
                commit('setIsAssigningFromUrl', false);
            });
        }
    }
}