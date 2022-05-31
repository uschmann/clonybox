<template>
  <v-list>
    <NavigationItem icon="mdi-speaker-wireless"
                    title="Spotify devices"
                    :to="{name: 'device'}"/>
    <NavigationItem icon="mdi-playlist-music"
                    title="Playback configuration"
                    :to="{name: 'playbackConfig'}"/>
  </v-list>
</template>

<script>
import NavigationItem from "@/components/navigation/NavigationItem";

export default {
  name: "Navigation",
  components: {
    NavigationItem
  }
}
</script>

<style scoped>

</style>