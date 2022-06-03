<template>
  <v-list-item v-bind="$attrs">
    <v-list-item-avatar :rounded="false">
      <v-img :src="imageUrl"/>
    </v-list-item-avatar>
    <v-list-item-content>
      <v-list-item-title>
        {{ item.metadata.name }}
      </v-list-item-title>
      <v-list-item-subtitle>
        {{ subtitle }}
      </v-list-item-subtitle>
    </v-list-item-content>
    <v-list-item-action>
      <ActionMenu>
        <ActionMenuItem label="Löschen"
                        icon="mdi-trash-can"
                        @click="deletePlaybackConfig(item)"/>
      </ActionMenu>
    </v-list-item-action>
  </v-list-item>
</template>

<script>
import {mapActions} from 'vuex'
import ActionMenu from "@/components/utils/ActionMenu";
import ActionMenuItem from "@/components/utils/ActionMenuItem";

export default {
  name: "SpotifyListItem",
  components: {
    ActionMenu,
    ActionMenuItem
  },
  props: {
    item: {
      type: Object,
      required: true,
    },
    type: {
      type: String,
      required: true
    }
  },
  methods: {
      ...mapActions('playbackConfig/list', [
          'deletePlaybackConfig'
      ])
  },
  computed: {
    imageUrl() {
      switch (this.type) {
        case "album":
        case "playlist":
          return this.item.metadata.images[0].url;
        case "track":
          return this.item.metadata.album.images[0].url;
      }
    },
    subtitle() {
      switch (this.type) {
        case "album":
        case "track":
          return this.item.metadata.artists[0].name;
        case "playlist":
          return `From: ${this.item.metadata.owner.display_name} | ${this.item.metadata.tracks.total} tracks`;
      }
    }
  }
}
</script>

<style scoped>

</style>