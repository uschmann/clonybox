<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <h1>Details</h1>
        <v-btn @click="searchSpotifyItem">
          search
        </v-btn>
      </v-col>
    </v-row>
    <SpotifySearchDialog ref="spotifyDialog"/>
  </v-container>
</template>

<script>
import SpotifySearchDialog from "@/components/spotify/SpotifySearchDialog";

export default {
  name: "PlaybackConfigDetailView",
  components: {
    SpotifySearchDialog
  },
  methods: {
    searchSpotifyItem() {
      this.$refs.spotifyDialog.open();
    }
  }
}
</script>

<style scoped>

</style>