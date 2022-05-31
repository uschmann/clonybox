import Vue from 'vue'
import Vuex from 'vuex'

import device from "./device"
import playbackConfig from "./playbackConfig"
import spotify from "./spotify"

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    device,
    playbackConfig,
    spotify
  }
})
