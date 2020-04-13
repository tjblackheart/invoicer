<template>
  <div>
    <h1 class="title is-4">
      Invoices
    </h1>
    <h2 class="subtitle is-6">
      Create new invoice
    </h2>

    <message />

    <hr>

    <form @submit.prevent="submit">
      <div class="columns">
        <div class="column">
          <b-input
            id="i.number"
            :value="invoice.number"
            type="text"
            label="Number"
            disabled
            :helplink="{ to: '/settings/numbers', text: 'Edit ... ' }"
          />
        </div>

        <div class="column">
          <b-select
            id="i.currency"
            v-model="invoice.currency"
            label="Currency"
            :options="currencies"
          />
        </div>

        <div class="column">
          <b-select
            id="i.customer"
            v-model.number="invoice.customer_id"
            label="Customer"
            :options="customerSelect"
            :helplink="{ to: '/customers/create', text: 'Add customer ...' }"
          />
        </div>

        <div class="column">
          <b-input
            id="i.date"
            v-model="formattedDate"
            type="date"
            label="Date"
          />
        </div>

        <div class="column">
          <b-input
            id="i.due_days"
            v-model.number="invoice.due_days"
            type="number"
            label="Due days"
          />
        </div>
      </div>

      <hr>

      <div class="columns">
        <div class="column">
          <div class="field">
            <label class="label">
              Items
            </label>
            <div class="control">
              <table
                v-if="invoice.items.length != 0"
                class="table is-fullwidth"
              >
                <thead>
                  <tr>
                    <th>Amount</th>
                    <th>Price/Unit</th>
                    <th>VAT</th>
                    <th>Description</th>
                    <th class="has-text-right" />
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="(i, index) in invoice.items"
                    :key="index"
                  >
                    <td data-label="Amount">
                      {{ i.amount }}{{ i.unit }}
                    </td>
                    <td data-label="Price/Unit">
                      {{ i.price_per_unit|money(invoice.currency) }}
                    </td>
                    <td data-label="VAT">
                      {{ i.vat }}%
                    </td>
                    <td
                      data-label="Description"
                      class="content"
                    >
                      <markdown :md="i.description" />
                    </td>
                    <td
                      data-label="Actions"
                      class="has-text-right"
                    >
                      <a @click.prevent="editItem(index)">
                        Edit
                      </a> &middot;
                      <a @click.prevent="removeItem(index)">
                        Remove
                      </a>
                    </td>
                  </tr>
                </tbody>
              </table>
              <p v-else>
                No items.
              </p>
            </div>
          </div>
          <button
            class="button is-small"
            @click.prevent="addItemModal"
          >
            Add item ...
          </button>
        </div>
      </div>

      <hr>

      <div class="columns">
        <div class="column">
          <ul class="list-group">
            <li class="list-group-item">
              <b>Subtotal:</b> <p class="is-pulled-right">
                {{ totalNet|money(invoice.currency) }}
              </p>
            </li>
            <li class="list-group-item">
              <b>VAT:</b> <p class="is-pulled-right">
                {{ vat|money(invoice.currency) }}
              </p>
            </li>
            <li class="list-group-item">
              <b>Total:</b> <p class="is-pulled-right">
                <b> {{ totalGross|money(invoice.currency) }} </b>
              </p>
            </li>
          </ul>
        </div>
      </div>

      <hr>

      <div class="has-text-right">
        <button
          :class="{'is-loading': busy}"
          :disabled="busy"
          class="button is-primary"
          @click.prevent="submit"
        >
          Create
        </button> &nbsp;
        <button
          class="button"
          @click.prevent="$router.go(-1)"
        >
          Cancel
        </button>
      </div>
    </form>

    <modal
      v-if="showItemModal"
      :title="title"
      :error="itemFormError"
      @close="close"
    >
      <div slot="content">
        <item-form :item="item" />
      </div>
      <button
        slot="action"
        class="button is-primary"
        @click.prevent="addItem"
      >
        <span v-if="edit">
          Edit item
        </span>
        <span v-else>
          Add item
        </span>
      </button>
    </modal>
  </div>
</template>

<script>
import http from '@/modules/http'
import Modal from '@/components/modals/Modal'
import ItemForm from '@/components/modals/Item'
import Markdown from '@/components/misc/MarkdownText'

import Message from '@/components/misc/Message'
import BInput from '@/components/fields/Input'
import BSelect from '@/components/fields/Select'

import dayjs from 'dayjs'
import dayjsPluginUTC from 'dayjs-plugin-utc'
dayjs.extend(dayjsPluginUTC)

const { required } = require('vuelidate/lib/validators')

export default {

  components: {
    Modal,
    ItemForm,
    Message,
    BInput,
    BSelect,
    Markdown,
  },

  data () {
    return {
      customers: [],
      invoice: {
        number: null,
        customer_id: null,
        currency: 'EUR',
        date: dayjs().toJSON(),
        items: [],
        due_days: 10,
      },
      error: null,
      showItemModal: false,
      item: {},
      customer: {},
      user: null,
      totalNet: 0,
      vat: 0,
      totalGross: 0,
      edit: false,
      busy: false,
      itemFormError: '',
      currencies: [
        {value: 'EUR', text: 'EUR'},
      ],
    }
  },

  computed: {
    title () {
      return this.edit ? 'Edit item' : 'Add item'
    },

    formattedDate: {
      get () {
        return dayjs.utc(this.invoice.date).format('YYYY-MM-DD')
      },
      set (value) {
        this.invoice.date = dayjs.utc(value).toJSON()
      },
    },

    customerSelect () {
      return this.customers.map(c => {
        return { value: c.id, text: c.address.company }
      })
    }
  },

  created () {
    this.$store.commit('clearMessage')
    this.load()
    this.initItem()
  },

  methods: {
    async load () {
      try {
        this.user = await http.fetchUser(this.$store.getters.uuid)

        if (this.user.settings.user_id === 0) {
          this.$store.commit('setMessage', {
            text: 'Please review your application settings.',
            style: 'is-warning',
          })
          this.$router.push('/settings/banking')
        }

        this.setInvoiceNumber()
        this.$store.commit('setUser', this.user)
        this.customers = await http.fetchCustomers()

        if (this.customers[0]) {
          this.customer = this.customers[0]
          this.invoice.customer_id = this.customer.id
        }
      } catch (error) {
        this.$store.commit('setMessage', { text: error.message, style: 'is-danger' })
      }
    },

    setInvoiceNumber () {
      const prefix = this.user.settings.invoice_number_prefix
      const nr = this.user.settings.next_invoice_number
      this.invoice.number = `${prefix}${nr}`
      this.invoice.user_id = this.user.id
    },

    addItemModal () {
      this.initItem()
      this.showItemModal = true
    },

    addItem () {
      this.itemFormError = ''
      this.$v.$touch()

      if (!this.$v.$error) {
        if (!this.edit) {
          this.invoice.items.push(this.item)
        }

        this.edit = false
        this.close()
        this.calculateTotals()
      } else {
        this.itemFormError = 'Please fill in all the data.'
      }
    },

    editItem (index) {
      this.edit = true
      this.item = this.invoice.items[index]
      this.showItemModal = true
    },

    removeItem (index) {
      this.invoice.items.splice(index, 1)
      this.calculateTotals()
    },

    calculateTotals () {
      this.totalNet = 0
      this.vat = 0
      this.totalGross = 0

      this.invoice.items.map(i => {
        const net = i.price_per_unit * i.amount
        this.totalNet += net
        this.vat += net / 100 * i.vat
        this.totalGross += net * (i.vat / 100 + 1)
      })
    },

    close () {
      this.edit = false
      this.showItemModal = false
      this.itemFormError = ''
    },

    async submit () {
      this.$store.commit('clearMessage')
      this.busy = true

      if (!this.invoice.items.length) {
        this.$store.commit('setMessage', {
          text: 'Please add some items first.',
          style: 'is-warning',
        })
        this.busy = false
        return
      }

      try {
        await http.postInvoice(this.invoice)
        this.$router.push('/invoices')
      } catch (error) {
        this.$store.commit('setMessage', {
          text: error.message,
          style: 'is-danger',
        })
      } finally {
        this.busy = false
      }
    },

    initItem () {
      this.item = {
        vat: 19,
        unit: 'h',
      }
    },
  },

  validations: {
    item: {
      amount: { required },
      price_per_unit: { required },
      vat: { required },
    },
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
</style>
