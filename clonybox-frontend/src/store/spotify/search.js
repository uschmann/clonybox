import {spotifyApi} from "@/api";

export default {
    namespaced: true,
    state: {
        query: "",
        result: null,
        isLoading: false
    },
    mutations: {
        reset(state) {
            state.query = "";
            state.result = null;
        },
        setIsLoading(state, isLoading) {
            state.isLoading = isLoading;
        },
        setResult(state, result) {
            state.result = result;
        },
        updateQuery(state, query) {
            state.query = query;
        }
    },
    actions: {
        search({commit, state}) {
            commit('setIsLoading', true);

            return spotifyApi.search(state).then(result => {
                commit('setResult', result);
            }).finally(() => {
                commit('setIsLoading', false);
            });
        }
    }
}