<template>
  <div>
    <div class="is-clearfix">
      <div class="is-pulled-left">
        <h1 class="title is-4">
          Invoices
        </h1>
        <h2 class="subtitle is-6">
          Details for <b>{{ invoice.number }}</b>
        </h2>
      </div>
      <div class="is-pulled-right">
        <button
          :class="{'is-loading': busy}"
          :disabled="busy"
          class="button is-primary"
          @click.prevent="printPdf">
          Print to PDF
        </button>
      </div>
    </div>

    <hr>

    <message />

    <div class="box">
      <div class="box-content">
        <div class="columns is-mobile">
          <div class="column is-7">
            <div class="logo">
              TODO: logo
            </div>

            <div class="content">
              <p> <strong> {{ address.company }} </strong> </p>
              <p>
                <span v-if="address.first_name">
                  {{ address.first_name }}
                </span>
                <span v-if="address.last_name">
                  {{ address.last_name }}
                </span> <br>
                {{ address.street }} {{ address.number }} <br>
                {{ address.zip }} {{ address.city }} <br>
              </p>
            </div>
          </div>
          <div class="column is-5">
            <div class="content has-text-right">
              <p class="has-text-grey">
                <small>
                  <strong>{{ user.settings.company }}</strong><br>
                  {{ fullname }} <br>
                  {{ user.settings.street }} {{ user.settings.number }}<br>
                  {{ user.settings.zip }} {{ user.settings.city }} <br>
                  {{ user.settings.country }}
                </small>
              </p>

              <p class="has-text-grey">
                <small>
                  {{ user.settings.phone }}<br>
                  {{ email }}
                </small>
              </p>
            </div>
          </div>
        </div>

        <nav class="level is-mobile has-background-white-ter">
          <div class="level-item has-text-centered">
            <div>
              <p class="heading">
                Invoice
              </p> <p class="title is-6">
                {{ invoice.number }}
              </p>
            </div>
          </div>
          <div class="level-item has-text-centered">
            <div>
              <p class="heading">
                Customer
              </p> <p class="title is-6">
                {{ invoice.customer.number }}
              </p>
            </div>
          </div>
          <div class="level-item has-text-centered">
            <div>
              <p class="heading">
                Date
              </p> <p class="title is-6">
                {{ invoice.date|date }}
              </p>
            </div>
          </div>
        </nav>

        <hr>

        <div class="table-container">
          <table class="table is-fullwidth">
            <thead>
              <tr>
                <th>Pos</th>
                <th>Qty</th>
                <th>Desc</th>
                <th>Price/Unit</th>
                <th class="has-text-right">
                  Total (Net)
                </th>
                <th class="has-text-right">
                  VAT
                </th>
                <th class="has-text-right">
                  Total (Gross)
                </th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="(i, index) in invoice.items"
                :key="i.id">
                <td> {{ index + 1 }} </td>
                <td> {{ i.amount }} {{ i.unit }} </td>
                <td> {{ i.description }} </td>
                <td> {{ i.price_per_unit | money(invoice.currency) }} </td>
                <td class="has-text-right">
                  {{ totalNet(i) | money(invoice.currency) }}
                </td>
                <td class="has-text-right">
                  {{ i.vat }}%
                </td>
                <td class="has-text-right">
                  {{ totalGross(i) | money(invoice.currency) }}
                </td>
              </tr>
            </tbody>
          </table>

          <hr>
        </div>

        <div class="columns">
          <div class="column is-7" />
          <div class="column is-5">
            <ul class="list-group">
              <li class="list-group-item">
                <b>Subtotal:</b>
                <p class="is-pulled-right">
                  {{ totals.net|money(invoice.currency) }}
                </p>
              </li>
              <li class="list-group-item">
                <b>VAT:</b>
                <p class="is-pulled-right">
                  {{ totals.vat|money(invoice.currency) }}
                </p>
              </li>
              <li class="list-group-item">
                <b>Total:</b>
                <p class="is-pulled-right">
                  <b> {{ totals.gross|money(invoice.currency) }} </b>
                </p>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <hr>

    <div class="has-text-right">
      <router-link
        :to="{name: 'invoice_list'}"
        class="button">
        Back
      </router-link>
    </div>

    <modal
      v-if="showModal"
      @close="showModal = false">
      <div slot="content">
        <p> This is still a TODO. Sorry about that.</p>
      </div>
    </modal>
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Modal from '@/components/misc/Modal.vue'
import Message from '@/components/misc/Message'

export default {
  components: {
    Modal,
    Message,
  },

  data () {
    return {
      invoice: {
        customer: {
          number: null,
        },
      },
      user: {
        settings: {},
      },
      address: {},
      error: null,
      totals: {
        net: 0,
        vat: 0,
        gross: 0,
      },
      showModal: false,
      busy: false,
    }
  },

  computed: {
    fullname () {
      return `${this.user.settings.first_name} ${this.user.settings.last_name}`
    },

    email () {
      return this.user.settings.email || this.user.email
    },
  },

  created () {
    this.clearMessage()
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    totalNet (item) {
      const net = item.price_per_unit * item.amount
      return net
    },

    totalGross (item) {
      const net = this.totalNet(item)
      const vat = net / 100 * item.vat
      return net + vat
    },

    totalVat (item) {
      return this.totalNet(item) / 100 * item.vat
    },

    calculateTotals () {
      this.invoice.items.forEach(i => {
        let net = this.totalNet(i)
        let vat = this.totalVat(i)

        this.totals.net += net
        this.totals.vat += vat
        this.totals.gross += net + vat
      })
    },

    printPdf () {
      // TODO
      this.showModal = true
    },
  },

  beforeRouteEnter (to, from, next) {
    next(async vm => {
      try {
        const invoice = await http.fetchInvoice(to.params.id)
        const user = await http.fetchUser(vm.$store.getters.uuid)

        vm.invoice = invoice
        vm.address = invoice.customer.address
        vm.calculateTotals()
        vm.user = user
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      }
    })
  },
}
</script>

<style scoped>
  .list-group-item {
    padding: 10px 15px;
    margin-bottom: -1px;
    background-color: white;
    border: 1px solid #ddd;
  }
  .is-borderless > .list-group-item {
    border: none;
    padding: 5px 0;
  }
  .table-container {
    margin-bottom: 0;
  }
  .logo {
    padding: 0 0 40px 0;
  }
  nav.level {
    margin-top: 40px;
    padding: 20px 0;
  }
</style>
