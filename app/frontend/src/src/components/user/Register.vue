<template>
  <form @submit.prevent="submit">
    <h3 class="title">
      Register
    </h3>
    <p class="subtitle">
      Create an account
    </p>

    <message />

    <div
      v-if="!created"
      class="box">
      <form>
        <div class="field">
          <div class="control">
            <input
              v-model.trim="u.username"
              class="input is-medium"
              type="email"
              placeholder="Your name"
              autofocus>
          </div>
        </div>

        <div class="field">
          <div class="control">
            <input
              v-model.trim="u.email"
              class="input is-medium"
              type="email"
              placeholder="Email">
          </div>
        </div>

        <hr>

        <div class="field">
          <div class="control">
            <input
              v-model.trim="u.password"
              class="input is-medium"
              type="password"
              placeholder="Password">
          </div>
        </div>

        <div class="field">
          <div class="control">
            <input
              v-model.trim="u.repeat_password"
              class="input is-medium"
              type="password"
              placeholder="Repeat password">
          </div>
        </div>

        <hr>

        <div class="columns">
          <div class="column">
            <button
              :class="{'is-loading': busy}"
              :disabled="busy"
              class="button is-block is-info is-medium is-fullwidth"
              @click.prevent="submit">
              Create
            </button>
          </div>
          <div class="column">
            <button
              class="button is-block is-medium is-fullwidth"
              @click.prevent="$router.go(-1)">
              Cancel
            </button>
          </div>
        </div>
      </form>
    </div>
    <div v-else>
      <router-link to="/login">
        Back to login
      </router-link>
    </div>
  </form>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Message from '@/components/misc/Message'

export default {
  components: {
    Message,
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
