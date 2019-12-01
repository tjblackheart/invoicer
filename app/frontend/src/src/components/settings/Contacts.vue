<template>
  <div>
    <h3 class="title is-5">
      <icon name="envelope" /> Contacts
    </h3>

    <div class="fieldset">
      <div class="field">
        <label class="label">
          Company email
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model="$v.value.settings.email.$model"
              type="text"
              :class="['input', { 'is-danger': $v.value.settings.email.$error }]"
              @keyup="validate('email')">
          </div>
        </div>
      </div>

      <div class="field">
        <label class="label">
          Phone
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model="$v.value.settings.phone.$model"
              type="text"
              :class="['input', { 'is-danger': $v.value.settings.phone.$error }]"
              @keyup="validate('phone')">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/envelope'
const { required, email } = require('vuelidate/lib/validators')

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
      this.$v.value.settings[field].$touch()
      this.$emit('error', { view: 'contacts', errors: this.$v.$anyError })
    },
  },

  validations: {
    value: {
      settings: {
        email: { required, email },
        phone: { required },
      },
    },
  },
}
</script>
