<template>
  <div>
    <h3
      class="title is-5">
      <icon name="address-card" /> Company
    </h3>

    <div class="fieldset">
      <div class="field">
        <label class="label">
          Name
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model="$v.value.settings.company.$model"
              type="text"
              :class="['input', { 'is-danger': $v.value.settings.company.$error }]"
              @keyup="validate('company')">
          </div>
        </div>
      </div>

      <div class="columns">
        <div class="column">
          <div class="field">
            <label class="label">
              First name
            </label>
            <div class="control">
              <div class="control">
                <input
                  v-model="$v.value.settings.first_name.$model"
                  type="text"
                  :class="['input', { 'is-danger': $v.value.settings.first_name.$error }]"
                  @keyup="validate('company')">
              </div>
            </div>
          </div>
        </div>
        <div class="column">
          <div class="field">
            <label class="label">
              Last name
            </label>
            <div class="control">
              <div class="control">
                <input
                  v-model="$v.value.settings.last_name.$model"
                  type="text"
                  :class="['input', { 'is-danger': $v.value.settings.last_name.$error }]"
                  @keyup="validate('last_name')">
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="columns">
        <div class="column is-9">
          <div class="field">
            <label class="label">
              Street
            </label>
            <div class="control">
              <div class="control">
                <input
                  v-model="$v.value.settings.street.$model"
                  type="text"
                  :class="['input', { 'is-danger': $v.value.settings.street.$error }]"
                  @keyup="validate('street')">
              </div>
            </div>
          </div>
        </div>
        <div class="column">
          <div class="field">
            <label class="label">
              Number
            </label>
            <div class="control">
              <div class="control">
                <input
                  v-model="$v.value.settings.number.$model"
                  type="text"
                  :class="['input', { 'is-danger': $v.value.settings.number.$error }]"
                  @keyup="validate('number')">
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="columns">
        <div class="column is-3">
          <div class="field">
            <label class="label">
              Zip
            </label>
            <div class="control">
              <div class="control">
                <input
                  v-model="$v.value.settings.zip.$model"
                  type="text"
                  :class="['input', { 'is-danger': $v.value.settings.zip.$error }]"
                  @keyup="validate('zip')">
              </div>
            </div>
          </div>
        </div>

        <div class="column is-9">
          <div class="field">
            <label class="label">
              City
            </label>
            <div class="control">
              <div class="control">
                <input
                  v-model="$v.value.settings.city.$model"
                  type="text"
                  :class="['input', { 'is-danger': $v.value.settings.city.$error }]"
                  @keyup="validate('city')">
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="field">
        <label class="label">
          Country
        </label>
        <div class="control">
          <div class="control">
            <input
              v-model="$v.value.settings.country.$model"
              type="text"
              :class="['input', { 'is-danger': $v.value.settings.country.$error }]"
              @keyup="validate('country')">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/address-card'
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
      this.$emit('error', { key: 'company', errors: this.$v.$anyError })
    },
  },

  validations: {
    value: {
      settings: {
        company: { required },
        first_name: { required },
        last_name: { required },
        street: { required },
        number: { required },
        zip: { required },
        city: { required },
        country: { required },
      },
    },
  },
}
</script>
