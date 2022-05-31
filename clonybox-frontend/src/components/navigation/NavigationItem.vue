<template>
  <v-list-item :to="to">
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
    },
    to: {
      type: Object
    }
  }
}
</script>

<style scoped>

</style>