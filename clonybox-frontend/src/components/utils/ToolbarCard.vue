<template>
  <v-card>
    <v-toolbar color="primary lighten-1" dark flat dense>
      <slot name="title">
        <v-toolbar-title>
          <v-icon v-if="icon">
            {{icon}}
          </v-icon>
          {{title}}
        </v-toolbar-title>
      </slot>
      <v-spacer/>
      <v-toolbar-items>
        <slot name="toolbar-items"></slot>
      </v-toolbar-items>

      <template #extension v-if="$slots.extension">
        <slot name="extension"></slot>
      </template>
    </v-toolbar>
    <v-progress-linear indeterminate v-if="loading"/>
    <slot></slot>
  </v-card>
</template>

<script>
export default {
  name: "ToolbarCard",
  props: {
    title: {
      type: String
    },
    icon: {
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