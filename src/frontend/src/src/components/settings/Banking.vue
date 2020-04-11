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
        v-model="formattedIBAN"
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
import 'vue-awesome/icons/university'
import Icon from 'vue-awesome/components/Icon'
import BInput from '@/components/fields/Input'

const {
  required,
  helpers,
  minLength,
  maxLength
} = require('vuelidate/lib/validators')

const iban = helpers.regex('iban', /^[A-Z]{2}[0-9\s]+$/)

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

  data () {
    return {
      formattedIBAN: '',
    }
  },

  watch: {
    value () {
      this.$emit('input', this.value)
    },

    formattedIBAN (val) {
      const parts = val
        .toUpperCase()
        .split('')
        .map(v => v.trim())
        .join('')
        .match(/.{1,4}/g) || []

      this.formattedIBAN = parts.join(' ')
      this.value.settings.iban = this.formattedIBAN
    }
  },

  mounted () {
    this.formattedIBAN = this.value.settings && this.value.settings.iban || ''
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
        bank: {
          required
        },
        iban: {
          required,
          iban,
          // keep it simple here.
          // see https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
          minLength: minLength(15),
          maxLength: maxLength(32)
        },
        bic: {
          required
        },
      },
    },
  },
}
</script>
