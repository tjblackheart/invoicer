<template>
  <component :is="layout" />
</template>

<script>
import MainView from '@/views/Main.vue'
import LoginView from '@/views/Login.vue'

export default {
  components: {
    MainView,
    LoginView,
  },

  computed: {
    layout () {
      return this.$store.getters.layout
    },
  },
}
</script>
