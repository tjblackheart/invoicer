<template>
  <div>
    <h3 class="title is-5">
      <icon name="envelope" /> Contacts
    </h3>

    <div class="fieldset">
      <b-input
        id="s.email"
        v-model="$v.value.settings.email.$model"
        label="Email"
        :has-error="$v.value.settings.email.$error"
        :helptext="$v.value.settings.email.$error ? 'Please enter a valid email.' : ''"
        @input="validate('email')"
      />

      <b-input
        id="s.phone"
        v-model="$v.value.settings.phone.$model"
        label="Phone"
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
      let count = 0
      Object.keys(this.$v.value.settings).forEach(k => {
        if (this.$v.value.settings[k].$error) {
          count++
        }
      })

      this.$v.value.settings[field].$touch()
      this.$emit('error', {
        view: 'contacts',
        errors: this.$v.$anyError,
        count,
      })
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
