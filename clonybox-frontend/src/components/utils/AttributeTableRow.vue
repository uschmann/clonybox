<template>
  <tr>
    <td class="font-weight-bold">{{label}}</td>
    <td>{{value ? value : '-'}}</td>
  </tr>
</template>

<script>
export default {
  name: "AttributeTableRow",
  props: [
      'label',
      'value'
  ]
}
</script>

<style scoped>

</style>