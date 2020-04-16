<template>
  <div>
    <div class="is-clearfix">
      <div class="is-pulled-left">
        <h1 class="title is-4">
          Invoices
        </h1>
        <h2 class="subtitle is-6">
          All invoices
        </h2>
      </div>

      <div class="is-pulled-right">
        <router-link
          :to="{name: 'invoice_create'}"
          class="button is-primary"
        >
          Create
        </router-link>
      </div>
    </div>

    <hr>

    <message />

    <search
      placeholder="Number, Company ..."
      type="invoice"
    />

    <pager
      v-if="pages > 1"
      :current="currentPage"
      :pages="pages"
      class="is-hidden-desktop"
      @paged="onPaged($event)"
    />

    <invoice-table
      :paged-items="pagedItems"
      :filtered-items="filteredItems"
      @cancel="cancel($event)"
      @payment="paymentModal($event)"
    />

    <hr>

    <pager
      v-if="pages > 1"
      :current="currentPage"
      :pages="pages"
      @paged="onPaged($event)"
    />

    <modal
      v-if="showModal"
      title="Toggle payment"
      :error="paymentError"
      @close="close"
    >
      <div slot="content">
        <b-input
          id="p.date"
          v-model="paymentDate"
          type="date"
          label="Payment date"
        />
      </div>

      <button
        slot="action"
        class="button is-primary"
        :class="{'is-loading': busy}"
        :disabled="busy"
        @click.prevent="togglePayment"
      >
        Toggle
      </button>
    </modal>
  </div>
</template>

<script>
import http from '@/modules/http'
import dayjs from 'dayjs'
import Message from '@/components/misc/Message'
import Modal from '@/components/modals/Modal.vue'
import BInput from '@/components/fields/Input'
import Search from '@/components/misc/Search'
import Pager from '@/components/misc/Pager'
import InvoiceTable from '@/components/invoice/InvoiceTable'

export default {
  components: {
    Message,
    Modal,
    BInput,
    Search,
    Pager,
    InvoiceTable
  },

  data () {
    return {
      busy: false,
      invoices: [],
      error: null,
      showModal: false,
      activeInvoiceId: null,
      paymentError: null,
      paymentDate: null,
      currentPage: 1,
      perPage: 10,
    }
  },

  computed: {
    showCancelled () {
      return this.$store.getters.showCancelled
    },

    filters () {
      return this.$store.getters.filters
    },

    filteredItems () {
      let items = this.invoices

      if (!this.showCancelled) {
        items = items.filter(i => !i.is_cancelled)
      }

      if (this.filters.length){
        items = items.filter(i => {
          return this.filters.filter(v => {
            return i.number.toLowerCase().includes(v) || i.customer.address.company.toLowerCase().includes(v)
          }).length
        })
      }

      return items
    },

    pages () {
      return Math.ceil(this.filteredItems.length / this.perPage)
    },

    pagedItems () {
      const from = (this.currentPage - 1)  * this.perPage
      const to = this.currentPage * this.perPage
      return this.filteredItems.slice(from, to)
    }
  },

  watch: {
    '$route': 'load',

    filters () {
      this.currentPage = 1
    },

    showCancelled () {
      this.currentPage = 1
    },
  },

  created () {
    this.$store.commit('clearMessage')
    this.load()
    this.paymentDate = dayjs().format('YYYY-MM-DD')
  },

  methods: {
    async load () {
      try {
        this.invoices = await http.fetchInvoices()
        this.invoices.sort((a, b) => a.id < b.id ? 1 : -1)
      } catch (error) {
        this.$store.commit('setMessage', {
          text: error.message,
          style: 'is-danger',
        })
      }
    },

    async togglePayment () {
      this.paymentError = null
      this.busy = true

      try {
        this.invoices = await http.togglePayment({
          id: this.activeInvoiceId,
          date: this.paymentDate,
        })
        this.close()
      } catch (e) {
        this.paymentError = e.message
      } finally {
        this.busy = false
      }
    },

    async cancel(id) {
      if (!confirm("Really cancel? This can not be undone.")) {
        return
      }

      try {
        await http.cancelInvoice(id)
        this.invoices = await http.fetchInvoices()
      } catch (error) {
        this.$store.commit('setMessage', {
          text: error.message,
          style: 'is-danger',
        })
      }
    },

    paymentModal (id) {
      this.activeInvoiceId = id
      this.showModal = true
    },

    close () {
      this.activeInvoiceId = null
      this.showModal = false
      this.paymentError = null
    },

    onPaged (page) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth',
      })
      this.currentPage = page
    }
  },

  beforeRouteEnter (to, from, next) {
    next(vm => {
      if (from.name && !from.name.includes('invoice')) {
        vm.$store.commit('filterValue', '')
        vm.$store.commit('showCancelled', true)
      }
    })
  },
}
</script>
