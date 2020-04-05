<template>
  <div>
    <message />

    <b-select
      id="c.type"
      v-model="contact.type"
      label="Type"
      :options="types"
    />

    <b-input
      id="c.value"
      v-model.trim="contact.value"
      label="Value"
    />
  </div>
</template>

<script>
import Message from '@/components/misc/Message'
import BInput from '@/components/fields/Input'
import BSelect from '@/components/fields/Select'

export default {
  components: {
    Message,
    BInput,
    BSelect
  },

  props: {
    contact: {
      type: Object,
      default: () => {
        return {
          type: 'Email',
          value: '',
        }
      },
    },
  },

  data () {
    return {
      types: [
        { value: 'Email', text: 'Email' },
        { value: 'Phone', text: 'Phone' },
        { value: 'Other', text: 'Other' }
      ],
      error: null,
    }
  },
}
</script>
