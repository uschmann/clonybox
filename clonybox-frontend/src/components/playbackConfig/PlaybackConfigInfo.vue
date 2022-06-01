<template>
  <ToolbarCard title="Information">
    <v-container>
      <v-row>
        <v-col class="shrink"
               v-if="playbackConfig.metadata && playbackConfig.metadata.images && playbackConfig.metadata.images.length > 0">
          <v-img max-width="200"
                 :src="playbackConfig.metadata.images[0].url"/>
        </v-col>
        <v-col>
          <AttributeTable>
            <AttributeTableRow label="RFID"
                               :value="playbackConfig.rfid"/>
            <AttributeTableRow label="Name"
                               :value="playbackConfig.name"/>
            <AttributeTableRow label="Type"
                               :value="playbackConfig.type"/>
          </AttributeTable>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-btn v-if="playbackConfig.metadata"
                 :href="playbackConfig.metadata.external_urls.spotify"
                 target="_blank"
                 block
                 outlined
                 color="green">
            <v-icon>
              mdi-spotify
            </v-icon>
            View on spotify
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </ToolbarCard>
</template>

<script>
import {mapState} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";
import AttributeTable from "@/components/utils/AttributeTable";
import AttributeTableRow from "@/components/utils/AttributeTableRow";

export default {
  name: "PlaybackConfigInfo",
  components: {
    ToolbarCard,
    AttributeTable,
    AttributeTableRow
  },
  computed: {
    ...mapState('playbackConfig/detail', [
      'playbackConfig'
    ])
  }
}
</script>

<style scoped>

</style>