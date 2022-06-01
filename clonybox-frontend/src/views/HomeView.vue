<template>
  <v-container fluid>
    <v-row>
      <v-col class="text-center pt-16">
        <v-icon size="64" color="green">
          mdi-spotify
        </v-icon>
        <h2 class="font-weight-light" style="line-height: 0.9em">
          Hi {{user.display_name}},
          <br><small>Welcome to clonybox! 👋</small>
        </h2>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import {mapState} from 'vuex'

  export default {
    name: 'Home',
    computed: {
      ...mapState('auth', [
          'user'
      ])
    }
  }
</script>
