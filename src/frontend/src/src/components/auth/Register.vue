<template>
  <form @submit.prevent="submit">
    <h3 class="title"> Register </h3>
    <p class="subtitle"> Create an account </p>

    <message />

    <div
      v-if="!created"
      class="box"
    >
        <b-input
          v-model.trim="u.username"
          id="u.username"
          placeholder="Username"
          autofocus
        />

        <b-input
          v-model.trim="u.email"
          id="u.email"
          placeholder="Email"
          type="email"
        />

        <b-input
          v-model="u.password"
          id="u.password"
          type="password"
          placeholder="Password"
        />

        <b-input
          v-model="u.repeat_password"
          id="u.repeat_passwd"
          type="password"
          placeholder="Repeat password"
        />

        <hr>

        <div class="columns">
          <div class="column">
            <button
              type="submit"
              :class="{'is-loading': busy}"
              :disabled="busy"
              class="button is-block is-primary is-fullwidth"
            > Create
            </button>
          </div>

          <div class="column">
            <button
              class="button is-block is-fullwidth"
              @click.prevent="$router.go(-1)"
            > Cancel
            </button>
          </div>

        </div>
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
