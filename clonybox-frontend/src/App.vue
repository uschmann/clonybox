<template>
  <v-app>

    <template v-if="!isLoadingInfo">
      <template v-if="info.user !== null">
        <v-navigation-drawer v-model="drawer" app>
          <Navigation/>
        </v-navigation-drawer>

        <v-app-bar
            app
            color="primary"
            dark>
          <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
        </v-app-bar>

        <v-main>
          <router-view v-if="!isLoadingInfo"/>
        </v-main>

        <NotificationSnackbar/>
      </template>
      <template v-else>
        <div class="fill-height d-flex justify-center align-center">
          <div class="text-center">
            <h3>Welcome to <strong>Clonybox</strong></h3>
            <p>Please link your spotify account to continue</p>
            <v-btn color="green"
                   :href="info.auth_url"
                   outlined>
              <v-icon>
                mdi-spotify
              </v-icon>
              Connect to spotify
            </v-btn>
          </div>
        </div>
      </template>
    </template>
    <template v-else>
      <div class="d-flex align-center justify-center fill-height">
        <v-progress-circular indeterminate color="primary"/>
      </div>
    </template>
  </v-app>
</template>

<script>
import {mapState, mapActions} from 'vuex'
import Navigation from "@/components/navigation/Navigation";
import NotificationSnackbar from "@/components/utils/NotificationSnackbar";

export default {
  name: 'App',
  components: {
    Navigation,
    NotificationSnackbar
  },
  computed: {
    ...mapState([
      'isLoadingInfo',
      'info'
    ]),
  },
  methods: {
    ...mapActions([
      'loadInfo'
    ])
  },
  data: () => ({
    drawer: null
  }),
  mounted() {
    this.loadInfo();
  }
};
</script>
