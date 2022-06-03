<template>
  <ToolbarCard title="Spotify devices">
    <template #toolbar-items>
      <v-btn icon
             :loading="isLoading"
             @click="loadDevices">
        <v-icon>
          mdi-refresh
        </v-icon>
      </v-btn>
    </template>
    <v-list>
      <v-list-item v-for="device in devices" :key="'device-' + device.id">
        <v-list-item-icon>
          <v-icon>
            {{device.type | deviceIcon}}
          </v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            {{device.name}}
          </v-list-item-title>
          <v-list-item-subtitle>
            {{device.type}}
          </v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-action>
          <ActionMenu>
            <ActionMenuItem label="Set as default"
                            @click="setAsDefault(device)"
                            icon="mdi-eye"/>
          </ActionMenu>
        </v-list-item-action>
      </v-list-item>
    </v-list>
  </ToolbarCard>
</template>

<script>
import {deviceApi} from "@/api";
import {mapState, mapActions} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";
import ActionMenu from "@/components/utils/ActionMenu";
import ActionMenuItem from "@/components/utils/ActionMenuItem";
import deviceIcon from "@/filters/deviceIcon";

export default {
  name: "DeviceList",
  components: {
    ToolbarCard,
    ActionMenu,
    ActionMenuItem
  },
  filters: {
    deviceIcon
  },
  computed: {
    ...mapState('device/list', [
        'devices',
        'isLoading'
    ])
  },
  methods: {
    ...mapActions('device/list', [
         'loadDevices'
    ]),
    setAsDefault(device) {
      deviceApi.setDefault(device.id)
    }
  }
}
</script>

<style scoped>

</style>