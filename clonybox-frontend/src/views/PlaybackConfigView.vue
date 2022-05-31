<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <PlaybackConfigList/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import PlaybackConfigList from "@/components/playbackConfig/PlaybackConfigList";

export default {
  name: "PlaybackConfigView",
  components: {
    PlaybackConfigList
  }
}
</script>

<style scoped>

</style>