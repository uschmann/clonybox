<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <DeviceList v-if="devices !== null"/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import {mapState, mapActions} from 'vuex'
import DeviceList from "@/components/device/DeviceList";

export default {
  name: "DeviceView",
  components: {
      DeviceList
  },
  computed: {
    ...mapState('device/list', [
        'isLoading',
        'devices'
    ])
  },
  methods: {
    ...mapActions('device/list', [
        'loadDevices'
    ])
  },
  mounted() {
    this.loadDevices();
    console.log("FOOOOO")
  }
}
</script>

<style scoped>

</style>