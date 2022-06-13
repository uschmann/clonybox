import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import DeviceView from "@/views/DeviceView";
import PlaybackConfigView from "@/views/PlaybackConfigView";
import PlaybackConfigDetailView from "@/views/PlaybackConfigDetailView";
import SettingsView from "@/views/SettingsView";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/device',
    name: 'device',
    component: DeviceView
  },
  {
    path: '/playback-config',
    name: 'playbackConfig',
    component: PlaybackConfigView
  },
  {
    path: '/playback-config/:id',
    name: 'playbackConfig.detail',
    component: PlaybackConfigDetailView
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsView
  },
]

const router = new VueRouter({
  routes
})

export default router
