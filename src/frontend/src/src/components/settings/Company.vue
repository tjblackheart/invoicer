<template>
  <div>
    <h3 class="title is-5">
      <icon name="address-card" /> Company
    </h3>

    <div class="fieldset">
      <b-input
        id="s.company"
        v-model="$v.value.settings.company.$model"
        label="Company"
        :has-error="$v.value.settings.company.$error"
        :helptext="$v.value.settings.company.$error ? 'Please enter a company.' : ''"
        @input="validate('company')"
      />

      <div class="columns">
        <div class="column">
          <b-input
            id="s.first_name"
            v-model="$v.value.settings.first_name.$model"
            label="First name"
            :has-error="$v.value.settings.first_name.$error"
            :helptext="$v.value.settings.first_name.$error ? 'Please enter a first name.' : ''"
            @input="validate('first_name')"
          />
        </div>
        <div class="column">
          <b-input
            id="s.last_name"
            v-model="$v.value.settings.last_name.$model"
            label="Last name"
            :has-error="$v.value.settings.last_name.$error"
            :helptext="$v.value.settings.last_name.$error ? 'Please enter a last name.' : ''"
            @input="validate('last_name')"
          />
        </div>
      </div>

      <div class="columns nm">
        <div class="column is-9">
          <b-input
            id="s.street"
            v-model="$v.value.settings.street.$model"
            label="Street"
            :has-error="$v.value.settings.street.$error"
            :helptext="$v.value.settings.street.$error ? 'Please enter a street.' : ''"
            @input="validate('street')"
          />
        </div>
        <div class="column">
          <b-input
            id="s.number"
            v-model="$v.value.settings.number.$model"
            label="Number"
            :has-error="$v.value.settings.number.$error"
            :helptext="$v.value.settings.number.$error ? 'Please enter a number.' : ''"
            @input="validate('number')"
          />
        </div>
      </div>

      <div class="columns">
        <div class="column is-3">
          <b-input
            id="s.zipcode"
            v-model="$v.value.settings.zip.$model"
            label="Zipcode"
            :has-error="$v.value.settings.zip.$error"
            :helptext="$v.value.settings.zip.$error ? 'Please enter a zipcode.' : ''"
            @input="validate('zip')"
          />
        </div>

        <div class="column is-9">
          <b-input
            id="s.city"
            v-model="$v.value.settings.city.$model"
            label="City"
            :has-error="$v.value.settings.city.$error"
            :helptext="$v.value.settings.city.$error ? 'Please enter a city.' : ''"
            @input="validate('city')"
          />
        </div>
      </div>

      <b-input
        id="s.country"
        v-model="$v.value.settings.country.$model"
        label="Country"
        :has-error="$v.value.settings.country.$error"
        :helptext="$v.value.settings.country.$error ? 'Please enter a country.' : ''"
        @input="validate('country')"
      />
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/address-card'
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
      this.$v.value.settings[field].$touch()
      this.$emit('error', { view: 'company', errors: this.$v.$anyError })
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
