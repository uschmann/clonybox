<template>
  <ToolbarCard title="Playback configurations">
    <v-list>
      <template v-for="playbackConfig in playbackConfigs">
        <SpotifyListItem :key="'playbackconfig-' + playbackConfig.id"
                         v-if="playbackConfig.metadata"
                         :type="playbackConfig.type"
                         :item="playbackConfig.metadata"
                         :to="{name: 'playbackConfig.detail', params: {id: playbackConfig.id}}"/>
        <v-list-item :key="'playbackconfig-' + playbackConfig.id"
                     v-else
                     :to="{name: 'playbackConfig.detail', params: {id: playbackConfig.id}}">
          <v-list-item-icon>
            <v-icon>
              mdi-card-bulleted-outline
            </v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>
              {{ playbackConfig.name ? playbackConfig.name : 'Not assigned' }}
            </v-list-item-title>
            <v-list-item-subtitle>
              RFID: {{ playbackConfig.rfid }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
  </ToolbarCard>
</template>

<script>
import {mapState} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";
import SpotifyListItem from "@/components/spotify/SpotifyListItem";

export default {
  name: "PlaybackConfigList",
  components: {
    ToolbarCard,
    SpotifyListItem
  },
  computed: {
    ...mapState('playbackConfig/list', [
      'playbackConfigs'
    ])
  }
}
</script>

<style scoped>

</style>