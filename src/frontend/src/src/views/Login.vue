<template>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="columns is-centered">
          <div class="column is-mobile is-half is-one-quarter-desktop">
            <transition
              name="fade"
              mode="out-in"
            >
              <router-view />
            </transition>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  data () {
    return {}
  },
}
</script>
