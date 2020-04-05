<template>
  <div>
    <h3
      class="title is-5"
    >
      <icon name="cogs" /> Numbers
    </h3>

    <div class="fieldset">
      <b-input
        id="s.i.prefix"
        v-model="$v.value.settings.invoice_number_prefix.$model"
        label="Invoice number prefix"
        :has-error="$v.value.settings.invoice_number_prefix.$error"
        :helptext="$v.value.settings.invoice_number_prefix.$error ? 'Please enter a prefix.' : ''"
        @input="validate('invoice_number_prefix')"
      />

      <b-input
        id="s.i.next"
        v-model.number="$v.value.settings.next_invoice_number.$model"
        type="number"
        label="Next invoice number"
        :has-error="$v.value.settings.next_invoice_number.$error"
        :helptext="$v.value.settings.next_invoice_number.$error ? 'Please enter a number.' : ''"
        @input="validate('next_invoice_number')"
      />

      <b-input
        id="s.c.prefix"
        v-model="$v.value.settings.customer_number_prefix.$model"
        label="Customer number prefix"
        :has-error="$v.value.settings.customer_number_prefix.$error"
        :helptext="$v.value.settings.customer_number_prefix.$error ? 'Please enter a prefix.' : ''"
        @input="validate('customer_number_prefix')"
      />

      <b-input
        id="s.c.next"
        v-model.number="$v.value.settings.next_customer_number.$model"
        label="Next customer number"
        type="number"
        :has-error="$v.value.settings.next_customer_number.$error"
        :helptext="$v.value.settings.next_customer_number.$error ? 'Please enter a number.' : ''"
        @input="validate('next_customer_number')"
      />

      <b-input
        id="s.vat"
        v-model="$v.value.settings.tax_number.$model"
        label="VAT ID"
        :has-error="$v.value.settings.tax_number.$error"
        :helptext="$v.value.settings.tax_number.$error ? 'Please enter a VAT ID.' : ''"
        @input="validate('tax_number')"
      />
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/cogs'
const { required, numeric, alphaNum } = require('vuelidate/lib/validators')
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
      this.$emit('error', { view: 'numbers', errors: this.$v.$anyError })
    },
  },

  validations: {
    value: {
      settings: {
        invoice_number_prefix: { required, alphaNum },
        next_invoice_number: { required, numeric },
        customer_number_prefix: { required, alphaNum },
        next_customer_number: { required, numeric },
        tax_number: { required },
      },
    },
  },
}
</script>
