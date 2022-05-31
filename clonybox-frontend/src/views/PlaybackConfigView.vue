<template>
  <v-container fluid>
    <v-row v-if="playbackConfigs !== null">
      <v-col>
        <PlaybackConfigList/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {mapState, mapActions} from 'vuex';
import PlaybackConfigList from "@/components/playbackConfig/PlaybackConfigList";

export default {
  name: "PlaybackConfigView",
  components: {
    PlaybackConfigList
  },
  computed: {
    ...mapState('playbackConfig/list', [
        'playbackConfigs',
        'isLoading'
    ])
  },
  methods: {
    ...mapActions('playbackConfig/list', [
        'loadPlaybackConfigs'
    ])
  },
  mounted() {
    this.loadPlaybackConfigs();
  }
}
</script>

<style scoped>

</style>