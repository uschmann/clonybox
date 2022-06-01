<template>
  <v-container fluid>
    <template v-if="playbackConfig !== null && !isLoading">
      <v-row>
        <v-col>
          <PlaybackConfigInfo/>
        </v-col>
        <v-col>
          <AssignSpotifyItem/>
        </v-col>
      </v-row>
    </template>
  </v-container>
</template>

<script>
import {mapState, mapActions, mapMutations} from 'vuex'
import SpotifySearchDialog from "@/components/spotify/SpotifySearchDialog";
import AssignSpotifyItem from "@/components/playbackConfig/AssignSpotifyItem";
import PlaybackConfigInfo from "@/components/playbackConfig/PlaybackConfigInfo";

export default {
  name: "PlaybackConfigDetailView",
  components: {
    PlaybackConfigInfo,
    SpotifySearchDialog,
    AssignSpotifyItem
  },
  computed: {
    ...mapState('playbackConfig/detail', [
      'playbackConfig',
      'isLoading',
    ])
  },
  methods: {
    ...mapActions('playbackConfig/detail', [
      'loadPlaybackConfig',
    ])
  },
  mounted() {
    this.loadPlaybackConfig(this.$route.params.id);
  }
}
</script>

<style scoped>

</style>