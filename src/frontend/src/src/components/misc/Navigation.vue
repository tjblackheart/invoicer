<template>
  <nav
    class="navbar has-shadow is-fixed-top"
    role="navigation"
    aria-label="main navigation"
  >
    <div class="container is-fluid">
      <div class="navbar-brand">
        <div class="navbar-item">
          <router-link to="/">
            <b>Invoicer</b>
          </router-link>
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
            to="/invoices"
            :class="['navbar-item', {
              'is-active': currentPath.includes('/invoice')
            }]"
            @click.native="toggleMenu"
          >
            Invoices
          </router-link>
          <router-link
            to="/customers"
            :class="['navbar-item', {
              'is-active': currentPath.includes('/customers')
            }]"
            @click.native="toggleMenu"
          >
            Customers
          </router-link>
          <router-link
            to="/settings/user"
            :class="['navbar-item', {
              'is-active': currentPath.includes('/settings')
            }]"
            @click.native="toggleMenu"
          >
            Settings
          </router-link>
        </div>

        <div class="navbar-end">
          <div class="navbar-item is-hidden-touch">
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

    currentPath () {
      return this.$route.path
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

<style lang="scss" scoped>
.navbar-brand {
  .navbar-item {
    a {
      color: #363636;
      font-size: 20px;
    }
  }
}

@media screen and (max-width: 1023px) {
  .container {
    &.is-fluid {
      padding-left: 10px;
      padding-right: 10px;
    }
  }

  .navbar-menu {
    box-shadow: none !important;
    border-top: 1px solid #f9f9f9;
  }

  .navbar-end {
    div.navbar-item:last-child {
      padding-top: 20px;
      padding-bottom: 20px;
    }
  }
}
</style>
