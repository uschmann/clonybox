import Vue from 'vue'
import Vuex from 'vuex'

import device from "./device"
import playbackConfig from "./playbackConfig"

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
    playbackConfig
  }
})
