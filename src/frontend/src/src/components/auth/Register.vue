<template>
  <form @submit.prevent="submit">
    <h3 class="title">
      Register
    </h3>
    <p class="subtitle">
      Create an account
    </p>

    <hr>

    <div
      v-if="!created"
    >
      <div class="content">
        <b-input
          id="u.username"
          v-model.trim="u.username"
          placeholder="Username"
          autofocus
        />

        <b-input
          id="u.email"
          v-model.trim="u.email"
          placeholder="Email"
          type="email"
        />

        <b-input
          id="u.password"
          v-model="u.password"
          type="password"
          placeholder="Password"
        />

        <b-input
          id="u.repeat_passwd"
          v-model="u.repeat_password"
          type="password"
          placeholder="Repeat password"
        />

        <message />

        <hr>

        <button
          type="submit"
          :class="{'is-loading': busy}"
          :disabled="busy"
          class="button is-block is-primary is-fullwidth"
        >
          Register
        </button>
      </div>

      <p class="has-text-grey has-text-centered">
        <router-link :to="{ name: 'login' }">
          No thanks
        </router-link>
      </p>
    </div>

    <div v-else>
      <router-link :to="{ name: 'login' }">
        Back to login
      </router-link>
    </div>
  </form>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Message from '@/components/misc/Message'
import BInput from '@/components/fields/Input'

export default {
  components: {
    Message,
    BInput
  },

  data () {
    return {
      u: {},
      busy: false,
      created: false,
    }
  },

  created () {
    this.clearMessage()
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    async submit () {
      this.clearMessage()

      try {
        this.busy = true
        await http.createUser(this.u)
        this.setMessage({
          text: 'Account successfully created.',
        })
        this.created = true
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      } finally {
        this.busy = false
      }
    },
  },

  beforeRouteLeave (to, from, next) {
    this.clearMessage()
    next()
  },
}
</script>
