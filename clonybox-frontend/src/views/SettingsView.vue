<template>
  <v-container>
    <v-row>
      <v-col>
        <h2 class="font-weight-thin mb-5">Settings</h2>
        <Reset/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Reset from "@/components/settings/Reset";

export default {
  name: "SettingsView",
  components: {
    Reset
  }
}
</script>

<style scoped>

</style>