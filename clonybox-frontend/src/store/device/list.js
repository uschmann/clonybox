import {deviceApi} from "@/api";

export default {
    namespaced: true,
    state: {
        devices: null,
        isLoading: false
    },
    mutations: {
        setDevices(state, devices) {
            state.devices = devices;
        },
        setIsLoading(state, isLoading) {
            state.isLoading = isLoading;
        }
    },
    actions: {
        loadDevices({commit}) {
            commit('setIsLoading', true);

            return deviceApi.index().then(devices => {
                commit('setDevices', devices);
            }).finally(() => {
                commit('setIsLoading', false);
            });
        }
    }
}