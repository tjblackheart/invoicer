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
          class="button is-primary">
          Create
        </router-link>
      </div>
    </div>

    <hr>

    <message />

    <div class="table-container">
      <table
        v-if="invoices"
        class="table is-fullwidth">
        <thead>
          <tr>
            <!--<th>ID</th>-->
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
            v-for="invoice in invoices"
            :key="invoice.id"
            :class="{'is-cancelled': invoice.is_cancelled}"
          >
            <td> {{ invoice.number }} </td>
            <td> {{ invoice.date|date }} </td>
            <td
              :title="`${invoice.due_days} days`">
              {{ dueDate(invoice) }}
            </td>
            <td>
              <router-link :to="{name: 'customer_details', params: {id:invoice.customer.id}}">
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
          <tr v-if="invoices.length == 0">
            <td colspan="7">
              No invoices found.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <modal
      v-if="showModal"
      title="Toggle payment"
      :error="paymentError"
      @close="close">
      <div slot="content">
        <div class="field">
          <label class="label">
            Payment date
          </label>
          <div class="control">
            <input
              v-model="paymentDate"
              class="input"
              type="date">
          </div>
        </div>
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
import Modal from '@/components/misc/Modal.vue'
import PrintLink from './PrintLink.vue'

export default {
  components: {
    Message,
    Modal,
    PrintLink,
  },

  data () {
    return {
      busy: false,
      invoices: null,
      error: null,
      showModal: false,
      activeInvoiceId: null,
      paymentError: null,
      paymentDate: null,
    }
  },

  watch: {
    '$route': 'load',
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
}
</script>
