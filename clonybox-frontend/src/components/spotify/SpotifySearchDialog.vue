<template>
  <v-dialog :value="isOpen" persistent max-width="800">
    <ToolbarCard title="Search spotify">
      <template #toolbar-items>
        <v-btn icon @click="close">
          <v-icon>
            mdi-close
          </v-icon>
        </v-btn>
      </template>

      <v-container fluid>
        <v-row>
          <v-col>
            <v-text-field label="Search"
                          :value="query"
                          @input="updateQuery($event)"
                          outlined dense/>
            <v-btn color="primary"
                   block
                   @click="search"
                   :loading="isLoading">
              search
            </v-btn>
          </v-col>
        </v-row>


        <v-tabs v-model="tab">
          <v-tab>
            Albums
          </v-tab>
          <v-tab>
            Playlists
          </v-tab>
          <v-tab>
            Tracks
          </v-tab>
        </v-tabs>

        <v-row v-if="result !== null">

          <v-col>
            <v-tabs-items v-model="tab">
              <v-tab-item>
                <v-list>
                  <v-list-item v-for="album in result.albums.items"
                               :key="'album-' + album.id"
                               @click="select('album', album)">
                    <v-list-item-avatar :rounded="false">
                      <v-img :src="album.images[0].url"/>
                    </v-list-item-avatar>
                    <v-list-item-content>
                      <v-list-item-title>
                        {{album.name}}
                      </v-list-item-title>
                      <v-list-item-subtitle>
                        {{album.artists[0].name}}
                      </v-list-item-subtitle>
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-tab-item>

              <v-tab-item>
                <v-list>
                  <v-list-item v-for="playlist in result.playlists.items"
                               :key="'playlist-' + playlist.id"
                               @click="select('playlist', playlist)">
                    <v-list-item-avatar :rounded="false">
                      <v-img :src="playlist.images[0].url"/>
                    </v-list-item-avatar>
                    <v-list-item-content>
                      <v-list-item-title>
                        {{playlist.name}}
                      </v-list-item-title>
                      <v-list-item-subtitle>
                        From: {{playlist.owner.display_name}} | {{playlist.tracks.total}} tracks
                      </v-list-item-subtitle>
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-tab-item>

              <v-tab-item>
                <v-list>
                  <v-list-item v-for="track in result.tracks.items"
                               :key="'tracks-' + track.id"
                               @click="select('track', track)">
                    <v-list-item-avatar :rounded="false">
                      <v-img :src="track.album.images[0].url" v-if="track.album"/>
                    </v-list-item-avatar>
                    <v-list-item-content>
                      <v-list-item-title>
                        {{track.name}}
                      </v-list-item-title>
                      <v-list-item-subtitle>
                        {{track.artists[0].name}}
                      </v-list-item-subtitle>
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-tab-item>
            </v-tabs-items>
          </v-col>
        </v-row>
      </v-container>

      <v-divider/>

      <v-card-actions>
        <v-btn text @click="close">
          Cancel
        </v-btn>
        <v-spacer/>
      </v-card-actions>
    </ToolbarCard>
  </v-dialog>
</template>

<script>
import {mapState, mapActions, mapMutations} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";

export default {
  name: "SpotifySearchDialog",
  data() {
    return {
      isOpen: false,
      tab: null,
      resolve: null
    }
  },
  components: {
    ToolbarCard
  },
  computed: {
      ...mapState('spotify/search', [
          'isLoading',
          'result',
          'query'
      ])
  },
  methods: {
    select(type, item) {
      this.resolve({type, item});
      this.close();
    },
    ...mapActions('spotify/search', [
      'search'
    ]),
    ...mapMutations('spotify/search', [
      'updateQuery'
    ]),
    open() {
      this.isOpen = true;
      return new Promise(resolve => {
        this.resolve = resolve;
      });
    },
    close() {
      this.isOpen = false;
    }
  }
}
</script>

<style scoped>

</style>