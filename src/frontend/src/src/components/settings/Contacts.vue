<template>
  <div>
    <h3 class="title is-5">
      <icon name="envelope" /> Contacts
    </h3>

    <div class="fieldset">
      <b-input
        v-model="$v.value.settings.email.$model"
        label="Email"
        id="s.email"
        :has-error="$v.value.settings.email.$error"
        :helptext="$v.value.settings.email.$error ? 'Please enter a valid email.' : ''"
        @input="validate('email')"
      />

      <b-input
        v-model="$v.value.settings.phone.$model"
        label="Phone"
        id="s.phone"
        :has-error="$v.value.settings.phone.$error"
        :helptext="$v.value.settings.phone.$error ? 'Please enter a phone.' : ''"
        @input="validate('phone')"
      />
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/envelope'
const { required, email } = require('vuelidate/lib/validators')
import BInput from '@/components/fields/Input'

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
