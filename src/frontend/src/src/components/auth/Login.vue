<template>
  <form @submit.prevent="submit">
    <h3 class="title">
      Login
    </h3>
    <p class="subtitle">
      Please login to proceed.
    </p>

    <hr>

    <div class="content">
      <b-input
        id="c.email"
        v-model="credentials.email"
        placeholder="Email"
        type="email"
        autofocus
      />

      <b-input
        id="c.passwd"
        v-model="credentials.password"
        placeholder="Password"
        type="password"
      />

      <message />

      <hr>

      <button
        type="submit"
        :class="{'is-loading': busy}"
        :disabled="busy"
        class="button is-block is-primary is-fullwidth"
      >
        Login
      </button>
    </div>

    <p class="has-text-grey has-text-centered">
      <router-link :to="{ name: 'register' }">
        Register new account
      </router-link>
      <!--<a href="#">Forgot your password?</a>-->
    </p>
  </form>
</template>

<script>
import http from '@/modules/http'
import Message from '@/components/misc/Message.vue'
import BInput from '@/components/fields/Input'

export default {
  components: {
    Message,
    BInput
  },

  data () {
    return {
      credentials: {
        email: '',
        password: '',
      },
      busy: false,
      error: null,
      message: null,
    }
  },

  methods: {
    async submit () {
      try {
        this.busy = true
        this.$store.commit('clearMessage')

        const r = await http.login(this.credentials)
        this.$store.commit('login', { token: r.token, user: r.user })

        if (r.user.settings.user_id === 0) {
          this.$store.commit('setMessage', {
            text: 'Please review your application settings.',
            style: 'is-warning',
          })
          this.$router.push('/settings/banking')
        } else {
          this.$router.push('/')
        }
      } catch (err) {
        this.credentials.password = null
        this.$store.commit('setMessage', {
          text: err.message,
          style: 'is-warning',
        })
      } finally {
        this.busy = false
      }
    },
  },
}
</script>
