<template>
  <span v-if="!busy">
    &middot;
    <a
      href
      :class="css"
      @click.prevent="printInvoice">
      {{ text }}
    </a>
  </span>
  <span v-else>
    <font-awesome-icon icon="circle-notch" spin />
  </span>
</template>

<script>
import http from '@/modules/http'

export default {
  props: {
    id: {
      type: String,
      required: true,
    },
    css: {
      type: String,
      default: ''
    },
    text: {
      type: String,
      default: 'PDF',
    },
  },

  data () {
    return {
      busy: false,
    }
  },

  methods: {
    async printInvoice () {
      this.busy = true
      const r = await http.printInvoice(this.id)
      this.busy = false

      let a = document.createElement('a')
      a.href = `data:application/octet-stream;base64,${r.content}`
      a.download = r.filename
      a.click()
    },
  }
}
</script>
