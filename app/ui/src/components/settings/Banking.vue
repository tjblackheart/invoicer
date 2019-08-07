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
            v-model="value.settings.bank"
            type="text"
            class="input">
        </div>
      </div>
      <div class="field">
        <label class="label">
          IBAN
        </label>
        <div class="control">
          <input
            v-model="value.settings.iban"
            type="text"
            class="input">
        </div>
      </div>
      <div class="field">
        <label class="label">
          BIC
        </label>
        <div class="control">
          <input
            v-model="value.settings.bic"
            type="text"
            class="input">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/university'

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
