<template>
  <ToolbarCard title="Playback configurations">
    <template #toolbar-items>
      <v-btn icon
             :loading="isLoading"
             @click="loadPlaybackConfigs">
        <v-icon>
          mdi-refresh
        </v-icon>
      </v-btn>
    </template>
    <v-list>
      <template v-for="playbackConfig in playbackConfigs">
        <SpotifyListItem :key="'playbackconfig-' + playbackConfig.id"
                         v-if="playbackConfig.metadata"
                         :type="playbackConfig.type"
                         :item="playbackConfig"
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
          <v-list-item-action>
            <ActionMenu>
              <ActionMenuItem label="Löschen"
                              icon="mdi-trash-can"
                              @click="deletePlaybackConfig(playbackConfig)"/>
            </ActionMenu>
          </v-list-item-action>
        </v-list-item>
      </template>
    </v-list>
  </ToolbarCard>
</template>

<script>
import {mapState, mapActions} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";
import SpotifyListItem from "@/components/spotify/SpotifyListItem";
import ActionMenuItem from "@/components/utils/ActionMenuItem";
import ActionMenu from "@/components/utils/ActionMenu";

export default {
  name: "PlaybackConfigList",
  components: {
    ToolbarCard,
    SpotifyListItem,
    ActionMenu,
    ActionMenuItem
  },
  computed: {
    ...mapState('playbackConfig/list', [
      'playbackConfigs',
      'isLoading',
    ])
  },
  methods: {
    ...mapActions('playbackConfig/list', [
        'loadPlaybackConfigs',
        'deletePlaybackConfig'
    ])
  }
}
</script>

<style scoped>

</style>