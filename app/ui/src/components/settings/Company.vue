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
              v-model="value.settings.company"
              type="text"
              class="input">
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
                  v-model="value.settings.first_name"
                  type="text"
                  class="input">
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
                  v-model="value.settings.last_name"
                  type="text"
                  class="input">
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
                  v-model="value.settings.street"
                  type="text"
                  class="input">
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
                  v-model="value.settings.number"
                  type="text"
                  class="input">
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
                  v-model="value.settings.zip"
                  type="text"
                  class="input">
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
                  v-model="value.settings.city"
                  type="text"
                  class="input">
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
              v-model="value.settings.country"
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
import 'vue-awesome/icons/address-card'

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
