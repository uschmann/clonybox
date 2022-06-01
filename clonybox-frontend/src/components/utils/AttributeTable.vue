<template>
  <v-simple-table>
    <template v-slot:default>
      <tbody>
        <slot></slot>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>
export default {
  name: "AttributeTable"
}
</script>

<style scoped>

</style>