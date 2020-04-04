<template>
  <div>
    <message />

    <b-input
      v-model.number="item.amount"
      type="number"
      label="Amount"
      id="i.amount"
    />

    <b-select
      v-model="item.unit"
      label="Unit"
      id="i.unit"
      :options="units"
    />

    <b-input
      v-model.number="item.price_per_unit"
      type="number"
      label="Price per unit"
      id="i.ppu"
    />

    <b-input
      v-model.number="item.vat"
      type="number"
      label="VAT in %"
      id="i.vat"
    />

    <b-textarea
      v-model.trim="item.description"
      label="Description"
      id="i.desc"
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
