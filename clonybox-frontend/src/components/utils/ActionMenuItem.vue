<template>
  <v-list-item v-bind="$attrs" v-on="$listeners" :disabled="loading">
    <template v-if="loading === false">
      <v-list-item-action v-if="icon !== null">
        <v-icon>
          {{icon}}
        </v-icon>
      </v-list-item-action>
      <v-list-item-title>
        {{label}}
      </v-list-item-title>
    </template>
    <template v-else>
      <v-list-item-action>
        <v-progress-circular indeterminate
                             color="primary"/>
      </v-list-item-action>
      <v-list-item-title>
        {{label}}
      </v-list-item-title>
    </template>
  </v-list-item>
</template>

<script>
export default {
  name: "ActionMenuItem",
  props: {
    icon: {
      type: String,
      default: null
    },
    label: {
      type: String
    },
    loading: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style scoped>

</style>