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
              v-model="value.settings.invoice_number_prefix"
              type="text"
              class="input">
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
              v-model.number="value.settings.next_invoice_number"
              type="number"
              step="1"
              class="input">
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
              v-model="value.settings.customer_number_prefix"
              type="text"
              class="input">
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
              v-model.number="value.settings.next_customer_number"
              type="number"
              step="1"
              class="input">
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
              v-model.number="value.settings.tax_number"
              type="text"
              class="input">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/cogs'

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
}
</script>
