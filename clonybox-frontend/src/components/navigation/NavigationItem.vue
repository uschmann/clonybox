<template>
  <v-list-item>
    <v-list-item-icon>
      <v-icon>
        {{icon}}
      </v-icon>
    </v-list-item-icon>
    <v-list-item-title>
      {{title}}
    </v-list-item-title>
  </v-list-item>
</template>

<script>
export default {
  name: "NavigationItem",
  props: {
    icon: {
      type: String
    },
    title: {
      type: String
    }
  }
}
</script>

<style scoped>

</style>