<template>
  <div class="table-container">
    <table class="table is-fullwidth">
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
          <td data-label="Number">
            {{ invoice.number }}
          </td>

          <td data-label="Date">
            {{ invoice.date|date }}
          </td>

          <td
            :title="`${invoice.due_days} days`"
            data-label="Due date"
          >
            {{ dueDate(invoice) }}
          </td>

          <td data-label="Customer">
            <router-link :to="{ name: 'customer_details', params: { id:invoice.customer.id }}">
              {{ invoice.customer.address.company }}
            </router-link>
          </td>

          <td data-label="Items">
            {{ invoice.items.length }}
          </td>

          <td data-label="Total (Net)">
            <strong>{{ invoice.total_net | money(invoice.currency) }} </strong>
          </td>

          <td data-label="Paid">
            <span v-if="invoice.is_paid">
              {{ invoice.paid_at|date }}
            </span>
            <span v-else>
              <span v-if="invoice.is_cancelled">CANCELLED</span>
              <span v-else>-</span>
            </span>
          </td>

          <td
            class="has-text-right action-items"
            data-label="Actions"
          >
            <span>
              <router-link :to="{name: 'invoice_details', params: {id:invoice.id}}">
                View
              </router-link>
            </span>

            <span>
              <router-link :to="{name: 'invoice_copy', params: {id:invoice.id}}">
                Copy
              </router-link>
            </span>

            <span v-if="!invoice.is_cancelled">
              <a
                href
                @click.prevent="$emit('cancel', invoice.id)"
              >
                Cancel
              </a>
            </span>

            <span v-if="!invoice.is_paid && !invoice.is_cancelled">
              <a
                href
                @click.prevent="$emit('payment', invoice.id)"
              >
                Toggle payment
              </a>
            </span>

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
</template>

<script>
import dayjs from 'dayjs'
import PrintLink from '@/components/invoice/PrintLink'

export default {
  components: {
    PrintLink
  },

  props: {
    pagedItems: {
      type: Array,
      default: () => [],
    },
    filteredItems: {
      type: Array,
      default: () => [],
    }
  },

  methods: {
    dueDate (invoice) {
      return dayjs(invoice.date).add(invoice.due_days, 'days').format('DD.MM.YYYY')
    },
  }
}
</script>
