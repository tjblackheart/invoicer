<template>
  <div>
    <h3
      class="title is-5">
      <icon name="cogs" /> Numbers
    </h3>

    <div class="fieldset">
      <div class="field">
        <label class="label">
          Invoicenumber - Prefix
        </label>

        <div class="control">
          <div class="control">
            <input
              v-model.trim="$v.value.settings.invoice_number_prefix.$model"
              type="text"
              :class="['input', { 'is-danger': $v.value.settings.invoice_number_prefix.$error }]"
              @keyup="validate('invoice_number_prefix')">
          </div>
        </div>
      </div>

      <div class="field">
        <label class="label">
          Next invoice number
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model.trim.number="$v.value.settings.next_invoice_number.$model"
              type="number"
              step="1"
              :class="['input', { 'is-danger': $v.value.settings.next_invoice_number.$error }]"
              @keyup="validate('next_invoice_number')">
          </div>
        </div>
      </div>

      <div class="field">
        <label class="label">
          Customernumber - Prefix
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model.trim="$v.value.settings.customer_number_prefix.$model"
              type="text"
              :class="['input', { 'is-danger': $v.value.settings.customer_number_prefix.$error }]"
              @keyup="validate('customer_number_prefix')">
          </div>
        </div>
      </div>

      <div class="field">
        <label class="label">
          Next customer number
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model.trim.number="$v.value.settings.next_customer_number.$model"
              type="number"
              step="1"
              :class="['input', { 'is-danger': $v.value.settings.next_customer_number.$error }]"
              @keyup="validate('next_customer_number')">
          </div>
        </div>
      </div>

      <div class="field">
        <label class="label">
          VAT ID
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model.trim="$v.value.settings.tax_number.$model"
              type="text"
              :class="['input', { 'is-danger': $v.value.settings.tax_number.$error }]"
              @keyup="validate('tax_number')">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/cogs'
const { required, numeric, alphaNum } = require('vuelidate/lib/validators')

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
