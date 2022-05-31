<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" clipped app>
      <Navigation/>
    </v-navigation-drawer>

    <v-app-bar
      app
      color="primary"
      dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

import Navigation from "@/components/navigation/Navigation";

export default {
  name: 'App',
  components: {
    Navigation
  },
  data: () => ({
    drawer: null
  }),
};
</script>
