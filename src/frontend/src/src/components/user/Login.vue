<template>
  <form @submit.prevent="submit">
    <h3 class="title">
      Login
    </h3>
    <p class="subtitle">
      Please login to proceed.
    </p>

    <message />

    <div class="box">
      <form>
        <div class="field">
          <div class="control">
            <input
              v-model="credentials.email"
              class="input is-medium"
              type="email"
              placeholder="Your Email"
              autofocus>
          </div>
        </div>
        <div class="field">
          <div class="control">
            <input
              v-model="credentials.password"
              class="input is-medium"
              type="password"
              placeholder="Your Password">
          </div>
        </div>
        <button
          :class="{'is-loading': busy}"
          :disabled="busy"
          class="button is-block is-info is-medium is-fullwidth"
          @click.prevent="submit">
          Login
        </button>
      </form>
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
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Message from '@/components/misc/Message.vue'

export default {
  components: {
    Message,
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
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    async submit () {
      try {
        this.busy = true
        this.clearMessage()

        const r = await http.login(this.credentials)
        this.$store.commit('login', { token: r.token, user: r.user })

        if (r.user.settings.user_id === 0) {
          this.setMessage({
            text: 'Please review your application settings.',
            style: 'is-warning',
          })
          this.$router.push('/settings')
        } else {
          this.$router.push('/')
        }
      } catch (err) {
        this.credentials.password = null
        this.setMessage({
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
