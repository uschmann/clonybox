import {settingsApi} from "@/api";
import notification from "@/utils/notification";

export default {
    namespaced: true,
    state: {
        isResetting: false,
    },
    mutations: {
        setIsResetting(state, isResetting) {
            state.isResetting = isResetting;
        }
    },
    actions: {
        resetDevice({commit}) {
            commit('setIsResetting', true);
            notification.success('The devices will restart now...');

            setTimeout(() => {
                location.reload()
            }, 5000);

            return settingsApi.reset()
        }
    }
}