<template>
  <div>
    <div class="is-clearfix">
      <div class="is-pulled-left">
        <h1 class="title is-4"> Invoices </h1>
        <h2 class="subtitle is-6"> All invoices </h2>
      </div>

      <div class="is-pulled-right">
        <router-link
          :to="{name: 'invoice_create'}"
          class="button is-primary">
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

    <div class="table-container">
      <table
        v-if="invoices"
        class="table is-fullwidth">
        <thead>
          <tr>
            <!-- <th>ID</th> -->
            <th>Number</th>
            <th>Date</th>
            <th>Due</th>
            <th>Customer</th>
            <th>Items</th>
            <th>Total (Net)</th>
            <th>Paid</th>
            <th class="has-text-right">
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(invoice, i) in pagedItems"
            :key="i"
            :class="{'is-cancelled': invoice.is_cancelled}"
          >
            <!-- <td> {{ invoice.id }} </td> -->
            <td> {{ invoice.number }} </td>
            <td> {{ invoice.date|date }} </td>
            <td
              :title="`${invoice.due_days} days`">
              {{ dueDate(invoice) }}
            </td>
            <td>
              <router-link :to="{
                name: 'customer_details',
                params: { id:invoice.customer.id }
              }">
                {{ invoice.customer.address.company }}
              </router-link>
            </td>
            <td> {{ invoice.items.length }} </td>
            <td> <strong>{{ invoice.total_net | money(invoice.currency) }} </strong> </td>
            <td>
              <span v-if="invoice.is_paid">
                {{ invoice.paid_at|date }}
              </span>
              <span v-else>
                <span v-if="invoice.is_cancelled">CANCELLED</span>
                <span v-else>-</span>
              </span>
            </td>
            <td class="has-text-right">
              <router-link :to="{name: 'invoice_details', params: {id:invoice.id}}">
                View
              </router-link>

              <span v-if="!invoice.is_cancelled">
                &middot;
                <a href @click.prevent="cancel(invoice.id)">Cancel</a>
              </span>

              <span v-if="!invoice.is_paid && !invoice.is_cancelled">
                &middot;
                <a
                  @click.prevent="paymentModal(invoice.id)">
                  Toggle payment
                </a>
              </span>

              &middot;
              <print-link :id="invoice.id" />
            </td>
          </tr>
          <tr v-if="filteredItems.length == 0">
            <td colspan="7">
              Nothing found.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <hr>

    <pager
      v-if="pages > 1"
      :current="currentPage"
      :pages="pages"
      @paged="currentPage = $event"
    />

    <modal
      v-if="showModal"
      title="Toggle payment"
      :error="paymentError"
      @close="close">
      <div slot="content">
        <b-input
          v-model="paymentDate"
          type="date"
          label="Payment date"
          id="p.date"
        />
      </div>

      <button
        slot="action"
        class="button is-primary"
        :class="{'is-loading': busy}"
        :disabled="busy"
        @click.prevent="togglePayment">
        Toggle
      </button>
    </modal>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import dayjs from 'dayjs'
import Message from '@/components/misc/Message'
import Modal from '@/components/modals/Modal.vue'
import BInput from '@/components/fields/Input'
import PrintLink from './PrintLink'
import Search from '@/components/misc/Search'
import Pager from '@/components/misc/Pager'

export default {
  components: {
    Message,
    Modal,
    PrintLink,
    BInput,
    Search,
    Pager
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
            return i.number.includes(v) || i.customer.address.company.toLowerCase().includes(v)
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
    }
  },

  created () {
    this.clearMessage()
    this.load()
    this.paymentDate = dayjs().format('YYYY-MM-DD')
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    async load () {
      try {
        this.invoices = await http.fetchInvoices()
        this.invoices.sort((a, b) => a.id < b.id ? 1 : -1)
      } catch (error) {
        this.setMessage({
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
        this.setMessage({
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

    dueDate (invoice) {
      return dayjs(invoice.date).add(invoice.due_days, 'days').format('DD.MM.YYYY')
    },
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
