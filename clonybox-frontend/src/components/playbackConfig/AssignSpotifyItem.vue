<template>
  <ToolbarCard title="Assign Spotify item">
    <v-container fluid>
      <v-row>
        <v-col class="text-center">
          <v-btn @click="searchSpotifyItem" block color="green" outlined class="mb-3" :loading="isAssigningFromItem">
            <v-icon>
              mdi-spotify
            </v-icon>
            search on spotify
          </v-btn>
          <p class="black--text">Or</p>
          <v-row>
            <v-col class="py-0" md="8">
              <v-text-field label="Spotify-URL"
                            :value="spotifyUrl"
                            @input="updateSpotifyUrl($event)"
                            outlined dense/>
            </v-col>
            <v-col class="py-0">
              <v-btn outlined
                     :loading="isAssigningFromUrl"
                     color="primary"
                     block
                     class="mb-3"
                     @click="assignFromSpotifyUrl">
                Assign from URL
              </v-btn>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-container>
    <SpotifySearchDialog ref="spotifyDialog"/>
  </ToolbarCard>
</template>

<script>
import {mapSTate, mapMutations, mapActions, mapState} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";
import SpotifySearchDialog from "@/components/spotify/SpotifySearchDialog";

export default {
  name: "AssignSpotifyItem",
  components: {
    ToolbarCard,
    SpotifySearchDialog
  },
  computed: {
    ...mapState('playbackConfig/detail', [
      'playbackConfig',
      'isAssigningFromUrl',
      'spotifyUrl',
      'isAssigningFromItem',
    ])
  },
  methods: {
    ...mapMutations('playbackConfig/detail', [
      'updateSpotifyUrl'
    ]),
    ...mapActions('playbackConfig/detail', [
      'assignFromSpotifyUrl',
      'assignFromSpotifyItem'
    ]),
    searchSpotifyItem() {
      this.$refs.spotifyDialog.open().then(item => {
        this.assignFromSpotifyItem(item)
      });
    }
  }
}
</script>

<style scoped>

</style>