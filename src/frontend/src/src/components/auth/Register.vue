<template>
  <form @submit.prevent="submit">
    <h3 class="title">
      Register
    </h3>
    <p class="subtitle">
      Create an account
    </p>

    <hr>

    <div v-if="!created">
      <div class="content">
        <b-input
          id="u.username"
          v-model.trim="u.username"
          :has-error="v('username')"
          :helptext="v('username') ? 'Please enter a valid username.' : ''"
          placeholder="Username"
          autofocus
          @blur="$v.u.username.$touch()"
        />

        <b-input
          id="u.email"
          v-model.trim="u.email"
          :has-error="v('email')"
          :helptext="v('email') ? 'Please enter a valid email address.' : ''"
          placeholder="Email"
          type="email"
          @blur="$v.u.email.$touch()"
        />

        <b-input
          id="u.password"
          v-model="u.password"
          :has-error="v('password')"
          :helptext="v('password') ? 'The password should be at least 8 characters long.' : ''"
          type="password"
          placeholder="Password"
          @blur="$v.u.password.$touch()"
        />

        <b-input
          id="u.repeat_passwd"
          v-model="u.repeat_password"
          :has-error="v('repeat_password')"
          :helptext="v('repeat_password') ? 'The passwords should match.' : ''"
          type="password"
          placeholder="Repeat password"
          @blur="$v.u.repeat_password.$touch()"
        />

        <message />

        <hr>

        <button
          type="submit"
          :disabled="busy || hasErrors"
          :class="['button is-block is-fullwidth', {'is-primary': !hasErrors, 'is-loading': busy}]"
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
      <message />
      <hr>
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
import { required, minLength, sameAs, email } from 'vuelidate/lib/validators'

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

  computed: {
    hasErrors () {
      return this.$v.$invalid
    }
  },

  created () {
    this.clearMessage()
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    async submit () {
      this.clearMessage()

      if (this.hasErrors) {
        return
      }

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

    v (field) {
      return this.$v.u[field].$error
    }
  },

  beforeRouteLeave (to, from, next) {
    this.clearMessage()
    next()
  },

  validations: {
    u: {
      username: {
        required,
        minLength: minLength(3)
      },
      email: {
        required,
        email
      },
      password: {
        required,
        minLength: minLength(8)
      },
      repeat_password: {
        required,
        sameAsPassword: sameAs('password')
      }
    }
  }
}
</script>
