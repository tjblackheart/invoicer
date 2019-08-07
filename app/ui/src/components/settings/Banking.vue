<template>
  <div>
    <h3
      class="title is-5">
      <icon name="university" /> Bank
    </h3>

    <div class="fieldset">
      <div class="field">
        <label class="label">
          Bank
        </label>
        <div class="control">
          <input
            v-model="$v.value.settings.bank.$model"
            type="text"
            :class="['input', { 'is-danger': $v.value.settings.bank.$error }]"
            @keyup="validate('bank')">
        </div>
      </div>
      <div class="field">
        <label class="label">
          IBAN
        </label>
        <div class="control">
          <input
            v-model="$v.value.settings.iban.$model"
            type="text"
            :class="['input', { 'is-danger': $v.value.settings.iban.$error }]"
            @keyup="validate('iban')">
        </div>
      </div>
      <div class="field">
        <label class="label">
          BIC
        </label>
        <div class="control">
          <input
            v-model="$v.value.settings.bic.$model"
            type="text"
            :class="['input', { 'is-danger': $v.value.settings.bic.$error }]"
            @keyup="validate('bic')">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/university'
const { required } = require('vuelidate/lib/validators')

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
      this.$emit('error', { key: 'banking', errors: this.$v.$anyError })
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
