import Vue from 'vue'
import Vuex from 'vuex'

import {infoApi} from "@/api";

import device from "./device"
import playbackConfig from "./playbackConfig"
import spotify from "./spotify"
import auth from "./auth"

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    info: null,
    isLoadingInfo: true
  },
  mutations: {
    setInfo(state, info) {
      state.info = info;
    },
    setIsLoadingInfo(state, isLoading) {
      state.isLoadingInfo = isLoading;
    }
  },
  actions: {
    loadInfo({commit}) {
      commit('setIsLoadingInfo', true);

      return infoApi.index().then(info => {
        console.log(info)
        commit('setInfo', info);
        commit('auth/setUser', info.user);
      }).finally(() => {
        commit('setIsLoadingInfo', false);
      });
    }
  },
  modules: {
    device,
    playbackConfig,
    spotify,
    auth
  }
})
