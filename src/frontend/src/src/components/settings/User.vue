<template>
  <div>
    <h3 class="title is-5">
      <icon name="user-cog" /> User
    </h3>

    <b-input
      id="s.username"
      v-model="$v.value.username.$model"
      label="Username"
      :has-error="$v.value.username.$error"
      :helptext="$v.value.username.$error ? 'Please enter a valid username.' : ''"
      @input="validate('username')"
    />

    <b-input
      id="s.email"
      v-model="$v.value.email.$model"
      label="Email"
      :has-error="$v.value.email.$error"
      :helptext="$v.value.email.$error ? 'Please enter a valid email.' : ''"
      @input="validate('email')"
    />
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/user-cog'
import BInput from '@/components/fields/Input'

const { required, minLength, email } = require('vuelidate/lib/validators')

export default {
  components: {
    Icon,
    BInput
  },

  props: {
    value: {
      type: Object,
      required: true,
    },
  },

  watch: {
    value () {
      this.$emit('input', this.value)
    },
  },

  methods: {
    validate (field) {
      let count = 0
      Object.keys(this.$v.value).forEach(k => {
        if (this.$v.value[k].$error) {
          count++
        }
      })

      this.$v.value[field].$touch()
      this.$emit('error', {
        view: 'user',
        errors: this.$v.$anyError,
        count,
      })
    },
  },

  validations: {
    value: {
      username: { required, minLength: minLength(3) },
      email: { required, email },
    },
  },
}
</script>
