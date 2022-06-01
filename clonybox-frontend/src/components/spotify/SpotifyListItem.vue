<template>
  <v-list-item v-bind="$attrs">
    <v-list-item-avatar :rounded="false">
      <v-img :src="imageUrl"/>
    </v-list-item-avatar>
    <v-list-item-content>
      <v-list-item-title>
        {{ item.name }}
      </v-list-item-title>
      <v-list-item-subtitle>
        {{ subtitle }}
      </v-list-item-subtitle>
    </v-list-item-content>
  </v-list-item>
</template>

<script>
export default {
  name: "SpotifyListItem",
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
  computed: {
    imageUrl() {
      switch (this.type) {
        case "album":
        case "playlist":
          return this.item.images[0].url;
        case "track":
          return this.item.album.images[0].url;
      }
    },
    subtitle() {
      switch (this.type) {
        case "album":
        case "track":
          return this.item.artists[0].name;
        case "playlist":
          return `From: ${this.item.owner.display_name} | ${this.item.tracks.total} tracks`;
      }
    }
  }
}
</script>

<style scoped>

</style>