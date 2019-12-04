<template>
  <div>
    <message />

    <div class="field">
      <label class="label">
        Amount
      </label>
      <div class="control">
        <input
          v-model.number="item.amount"
          class="input"
          type="number">
      </div>
    </div>

    <div class="field">
      <label class="label">
        Unit
      </label>
      <div class="control">
        <div class="select">
          <select
            v-model="item.unit"
            class="is-fullwidth">
            <option
              v-for="(u, i) in units"
              :key="i"
              :value="u.name">
              {{ u.desc }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <div class="field">
      <label class="label">
        Price per unit
      </label>
      <div class="control">
        <input
          v-model.number="item.price_per_unit"
          class="input"
          type="number">
      </div>
    </div>

    <div class="field">
      <label class="label">
        VAT in %
      </label>
      <div class="control">
        <input
          v-model.number="item.vat"
          class="input"
          type="number">
      </div>
    </div>

    <div class="field">
      <label class="label">
        Description
      </label>
      <div class="control">
        <textarea
          v-model="item.description"
          class="textarea" />
      </div>
    </div>
  </div>
</template>

<script>
import Message from '@/components/misc/Message'

export default {
  components: {
    Message,
  },

  props: {
    item: {
      type: Object,
      required: true,
    },
  },

  data () {
    return {
      units: [
        { name: 'h', desc: 'Hours' },
        { name: 'd', desc: 'Days' },
      ],
      error: null,
    }
  },
}
</script>
