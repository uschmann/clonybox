<template>
  <v-list>
    <NavigationItem icon="mdi-speaker-wireless"
                    title="Spotify devices"/>
    <NavigationItem icon="mdi-playlist-music"
                    title="Playback configuration"/>
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