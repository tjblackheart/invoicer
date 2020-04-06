<template>
  <div>
    <h3 class="title is-5">
      <icon name="university" /> Bank
    </h3>

    <div class="fieldset">
      <b-input
        id="s.bank"
        v-model="$v.value.settings.bank.$model"
        label="Bank"
        :has-error="$v.value.settings.bank.$error"
        :helptext="$v.value.settings.bank.$error ? 'Please enter a name.' : ''"
        @input="validate('bank')"
      />

      <b-input
        id="s.iban"
        v-model="$v.value.settings.iban.$model"
        label="IBAN"
        :has-error="$v.value.settings.iban.$error"
        :helptext="$v.value.settings.iban.$error ? 'Please enter a valid IBAN.' : ''"
        @input="validate('iban')"
      />

      <b-input
        id="s.bic"
        v-model="$v.value.settings.bic.$model"
        label="BIC"
        :has-error="$v.value.settings.bic.$error"
        :helptext="$v.value.settings.bic.$error ? 'Please enter a BIC.' : ''"
        @input="validate('bic')"
      />
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/university'
const { required } = require('vuelidate/lib/validators')
import BInput from '@/components/fields/Input'

export default {
  components: {
    Icon,
    BInput,
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
        view: 'banking',
        errors: this.$v.$anyError,
        count
      })
    },
  },

  validations: {
    value: {
      settings: {
        bank: { required },
        iban: { required },
        bic: { required },
      },
    },
  },
}
</script>
