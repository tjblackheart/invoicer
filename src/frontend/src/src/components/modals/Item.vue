<template>
  <div>
    <message />

    <b-input
      id="i.amount"
      v-model.number="item.amount"
      type="number"
      label="Amount"
    />

    <b-select
      id="i.unit"
      v-model="item.unit"
      label="Unit"
      :options="units"
    />

    <b-input
      id="i.ppu"
      v-model.number="item.price_per_unit"
      type="number"
      label="Price per unit"
    />

    <b-input
      id="i.vat"
      v-model.number="item.vat"
      type="number"
      label="VAT in %"
    />

    <b-textarea
      id="i.desc"
      v-model.trim="item.description"
      label="Description"
    />
  </div>
</template>

<script>
import Message from '@/components/misc/Message'
import BInput from '@/components/fields/Input'
import BSelect from '@/components/fields/Select'
import BTextarea from '@/components/fields/Textarea'

export default {
  components: {
    Message,
    BInput,
    BSelect,
    BTextarea
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
        { value: 'h', text: 'Hours' },
        { value: 'd', text: 'Days' },
      ],
      error: null,
    }
  },
}
</script>
