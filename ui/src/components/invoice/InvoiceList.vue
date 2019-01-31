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
        class="table is-fullwidth is-striped">
        <thead>
          <tr>
            <!--<th>ID</th>-->
            <th>Number</th>
            <th>Date</th>
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
            :key="invoice.id">
            <td> {{ invoice.number }} </td>
            <td> {{ invoice.date|date }} </td>
            <td>
              <router-link :to="{name: 'customer_details', params: {id:invoice.customer.id}}">
                {{ invoice.customer.address.company }}
              </router-link>
            </td>
            <td> {{ invoice.items.length }} </td>
            <td> <strong>{{ invoice.total_net | money(invoice.currency) }} </strong> </td>
            <td>
              <span v-if="invoice.is_paid">
                Yes
              </span>
              <span v-else>
                No
              </span>
            </td>
            <td class="has-text-right">
              <router-link :to="{name: 'invoice_details', params: {id:invoice.id}}">
                View
              </router-link>
              &middot;
              <a @click.prevent="setPaid">
                Toggle payment
              </a>
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
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Message from '@/components/misc/Message'

export default {
  components: {
    Message,
  },

  data () {
    return {
      busy: false,
      invoices: null,
      error: null,
    }
  },

  watch: {
    '$route': 'load',
  },

  created () {
    this.clearMessage()
    this.load()
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

    close () {
      this.showModal = false
    },

    setPaid () {
      // TODO
    },
  },
}
</script>
