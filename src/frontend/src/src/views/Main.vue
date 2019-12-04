<template>
  <div>
    <navigation />
    <section class="section body">
      <div class="container">
        <transition
          name="fade"
          mode="out-in">
          <router-view />
        </transition>
      </div>
    </section>
  </div>
</template>

<script>
import Navigation from '@/components/misc/Navigation.vue'

export default {
  components: {
    Navigation,
  },

  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    },
  },

  created () {
    const html = document.getElementById('html')
    html.classList.add('has-navbar-fixed-top')
  },
}
</script>
