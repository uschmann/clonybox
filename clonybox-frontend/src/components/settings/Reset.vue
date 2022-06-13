<template>
  <ToolbarCard title="Reset device">
    <v-container>
      <v-row>
        <v-col>
          <v-alert type="error" outlined>
            You will loose all settings and your spotify account will be unlinked.
          </v-alert>
          <v-btn color="red" block dark :loading="isResetting" @click="resetDevice">
            Reset device
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </ToolbarCard>
</template>

<script>
import ToolbarCard from "@/components/utils/ToolbarCard";
import {mapActions, mapState} from "vuex";

export default {
  name: "Reset",
  components: {
    ToolbarCard
  },
  computed: {
    ...mapState('settings/reset', [
      'isResetting'
    ])
  },
  methods: {
    ...mapActions('settings/reset', [
      'resetDevice'
    ])
  }
}
</script>

<style scoped>

</style>