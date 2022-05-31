<template>
    <ToolbarCard title="Playback configurations">
      <v-list>
        <v-list-item v-for="playbackConfig in playbackConfigs" :key="'playbackconfig' + playbackConfig.id"
                     :to="{name: 'playbackConfig.detail', params: {id: playbackConfig.id}}">
          <v-list-item-avatar>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>
              {{playbackConfig.name ? playbackConfig.name : 'Not assigned'}}
            </v-list-item-title>
            <v-list-item-subtitle>
              RFID: {{playbackConfig.rfid}}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </ToolbarCard>
</template>

<script>
import {mapState} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";

export default {
  name: "PlaybackConfigList",
  components: {
    ToolbarCard
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