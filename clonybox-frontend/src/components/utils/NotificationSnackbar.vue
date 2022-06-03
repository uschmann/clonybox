<template>
  <v-snackbar v-model="isOpen"
              :color="color"
              :timeout="3000"
              align="left"
              top>
    <v-icon color="white">{{icon}}</v-icon>
    {{ text }}
  </v-snackbar>

</template>

<script>
import EventBus from "../../utils/EventBus";

export default {
  name: "NotificationSnackbar",
  mounted()
  {
    EventBus.$on('alert', message =>
    {
      this.message = message;
    });
  },
  data()
  {
    return {
      message: null,
    };
  },
  computed: {
    isOpen: {
      get()
      {
        return this.message !== null;
      },
      set(value)
      {
        if (!value)
        {
          this.message = null;
        }
      }
    },
    text()
    {
      return this.message ? this.message.text : '';
    },
    color()
    {
      return this.message ? this.message.color : '';
    },
    icon()
    {
      return this.message ? this.message.icon : '';
    }
  }
}
</script>

<style scoped>

</style>