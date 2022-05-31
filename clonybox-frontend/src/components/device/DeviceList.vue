<template>
  <ToolbarCard title="Spotify devices">
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
            <ActionMenuItem label="Anzeigen"
                            icon="mdi-eye"/>
          </ActionMenu>
        </v-list-item-action>
      </v-list-item>
    </v-list>
  </ToolbarCard>
</template>

<script>
import {mapState} from 'vuex'
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
        'devices'
    ])
  }
}
</script>

<style scoped>

</style>