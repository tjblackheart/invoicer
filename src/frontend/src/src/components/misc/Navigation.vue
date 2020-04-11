<template>
  <nav
    class="navbar has-shadow is-fixed-top"
    role="navigation"
    aria-label="main navigation"
  >
    <div class="container is-fluid">
      <div class="navbar-brand">
        <div class="navbar-item">
          <b>Invoicer</b>
        </div>
        <a
          id="burger"
          :class="{'is-active':menuOpen}"
          :aria-expanded="!menuOpen"
          role="button"
          class="navbar-burger"
          aria-label="menu"
          @click="toggleMenu()"
        >
          <span aria-hidden="true" />
          <span aria-hidden="true" />
          <span aria-hidden="true" />
        </a>
      </div>
      <div
        :class="{'is-active':menuOpen}"
        class="navbar-menu"
      >
        <div class="navbar-start">
          <router-link
            :to="{name: 'invoice_list'}"
            class="navbar-item"
            @click.native="toggleMenu"
          >
            Invoices
          </router-link>
          <router-link
            :to="{name: 'customer_list'}"
            class="navbar-item"
            @click.native="toggleMenu"
          >
            Customers
          </router-link>
          <router-link
            to="/settings/user"
            class="navbar-item"
            @click.native="toggleMenu"
          >
            Settings
          </router-link>
        </div>
        <div class="navbar-end">
          <div class="navbar-item is-visible-desktop">
            <p v-if="username">
              Welcome back, <b>{{ username }}! </b>
            </p>
            <p v-else>
              Welcome back!
            </p>
          </div>
          <div class="navbar-item">
            <a
              class="button is-small is-primary"
              @click.prevent="logout"
            >
              Logout
            </a>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  data () {
    return {
      menuOpen: false,
    }
  },

  computed: {
    username () {
      return this.$store.getters.username
    },
  },

  methods: {
    toggleMenu () {
      const d = window.getComputedStyle(document.getElementById('burger')).display
      if (d !== 'none') {
        this.menuOpen = !this.menuOpen
      }
    },

    logout () {
      this.$store.commit('logout')
      this.$router.go()
    },
  },
}
</script>
