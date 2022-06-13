<template>
  <div>
    <v-toolbar flat color="black" dark>
      <v-toolbar-title>
        <v-icon color="green">
         mdi-spotify
        </v-icon>
        clonybox
      </v-toolbar-title>
    </v-toolbar>

    <v-divider></v-divider>

    <v-list>
      <NavigationItem icon="mdi-home"
                      exact
                      title="Home"
                      :to="{name: 'home'}"/>
      <NavigationItem icon="mdi-speaker-wireless"
                      title="Spotify devices"
                      :to="{name: 'device'}"/>
      <NavigationItem icon="mdi-playlist-music"
                      title="Playback configuration"
                      :to="{name: 'playbackConfig'}"/>
      <NavigationItem icon="mdi-cog"
                      title="Settings"
                      :to="{name: 'settings'}"/>
    </v-list>
  </div>
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