<template>
  <div>
    <h3 class="title is-5">
      <icon name="user-cog" /> Change password
    </h3>

    <div class="fieldset">
      <b-input
        id="s.old_password"
        v-model="$v.oldPassword.$model"
        label="Your old password"
        type="password"
        :has-error="$v.oldPassword.$error"
        :helptext="$v.oldPassword.$error ? 'Please enter a password.' : ''"
        @input="validate('oldPassword')"
      />

      <b-input
        id="s.new_password"
        v-model="$v.newPassword.$model"
        label="New password"
        type="password"
        :has-error="$v.newPassword.$error"
        :helptext="$v.newPassword.$error ? 'The password should be at least 8 characters long.' : ''"
        @input="validate('newPassword')"
      />

      <b-input
        id="s.confirm_password"
        v-model="$v.confirmPassword.$model"
        label="Confirm password"
        type="password"
        :has-error="$v.confirmPassword.$error"
        :helptext="$v.confirmPassword.$error ? 'The passwords should match.' : ''"
        @input="validate('confirmPassword')"
      />

      <div class=" content has-text-right">
        <button
          :class="['button is-primary', {'is-loading': busy}]"
          :disabled="disabled"
          @click.prevent="updatePassword()"
        >
          Update password
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import http from '@/modules/http'
import 'vue-awesome/icons/user-cog'
import Icon from 'vue-awesome/components/Icon'
import BInput from '@/components/fields/Input'
const { required, minLength, sameAs } = require('vuelidate/lib/validators')

export default {
  components: {
    Icon,
    BInput,
  },

  props: {
    value: {
      type: Object,
      required: true,
    }
  },

  data () {
    return {
      oldPassword: '',
      newPassword: '',
      confirmPassword: '',
      busy: false,
    }
  },

  computed: {
    disabled () {
      const hasValues = this.oldPassword !== '' && this.newPassword !== '' && this.confirmPassword !== ''
      return this.busy || this.$v.$anyError || !hasValues
    }
  },

  methods: {
    validate (field) {
      let count = 0
      this.$v[field].$touch()

      Object.keys(this.$v).forEach(k => {
        if (k === field && this.$v[field].$error)  {
          count++
        }
      })

      this.$emit('error', {
        view: 'password',
        errors: this.$v.$anyError,
        count,
      })
    },

    async updatePassword () {
      this.busy = true
      this.$store.commit('clearMessage')

      try {
        const payload = {
          current: this.oldPassword,
          new: this.newPassword,
          confirm: this.confirmPassword,
        }

        await http.updatePassword(this.value.uuid, payload)

        this.oldPassword = ''
        this.newPassword = ''
        this.confirmPassword = ''
        this.$v.$reset()

        this.$store.commit('setMessage', {
          text: 'Your password was updated.',
          style: 'is-success',
        })
      } catch (err) {
        this.$store.commit('setMessage', {
          text: err,
          style: 'is-danger',
        })
      } finally {
        this.busy = false
      }
    }
  },

  validations: {
    oldPassword: {
      required
    },
    newPassword: {
      required,
      minLength: minLength(8),
    },
    confirmPassword: {
      required,
      sameAsPassword: sameAs('newPassword'),
    }
  }
}
</script>

<style>

</style>
