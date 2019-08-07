<template>
  <div>
    <h3 class="title is-5">
      <icon name="user-cog" /> User
    </h3>

    <div class="fieldset">
      <div class="field">
        <label class="label">
          Username
        </label>
        <div class="control">
          <input
            v-model.trim="$v.value.username.$model"
            type="text"
            :class="['input', { 'is-danger': $v.value.username.$error }]"
            @keyup="validate('username')">
        </div>
      </div>

      <div class="field">
        <label class="label">
          Email
        </label>
        <div class="control">
          <input
            v-model.trim="$v.value.email.$model"
            type="email"
            :class="['input', { 'is-danger': $v.value.email.$error }]"
            @keyup="validate('email')">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/user-cog'

const { required, minLength, email } = require('vuelidate/lib/validators')

export default {
  components: {
    Icon,
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
      this.$v.value[field].$touch()
      this.$emit('error', { key: 'user', errors: this.$v.$anyError })
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
