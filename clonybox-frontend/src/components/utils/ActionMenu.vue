<template>
  <v-menu v-on="$listeners">
    <template #activator="{on}">
      <v-btn icon v-on="on" @click.stop.prevent>
        <v-icon>
          {{icon}}
        </v-icon>
      </v-btn>
    </template>
    <v-list>
      <slot>
      </slot>
    </v-list>
  </v-menu>
</template>

<script>
export default {
  name: "ActionMenu",
  props: {
    icon: {
      type: String,
      default: 'mdi-dots-vertical'
    }
  }
}
</script>

<style scoped>

</style>