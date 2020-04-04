<template>
  <div>
    <div class="is-clearfix">
      <div class="is-pulled-left">
        <h1 class="title is-4">
          Customers
        </h1>
        <h2 class="subtitle is-6">
          All customers
        </h2>
      </div>

      <div class="is-pulled-right">
        <router-link
          :to="{name: 'customer_create'}"
          class="button is-primary">
          Create
        </router-link>
      </div>
    </div>

    <hr>

    <message />

    <search
      placeholder="Number, Name, Zipcode ..."
      type="customers"
    />

    <div class="table-container">
      <table
        v-if="filteredItems"
        class="table is-fullwidth">
        <thead>
          <tr>
            <th>Number</th>
            <th>Company</th>
            <th>City</th>
            <th class="has-text-right">
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="customer in filteredItems"
            :key="customer.id">
            <td> {{ customer.number }} </td>
            <td> {{ customer.address.company }} </td>
            <td> {{ customer.address.city }}, {{ customer.address.zip }} </td>
            <td class="has-text-right">
              <router-link :to="{ name: 'customer_details', params: {id:customer.id} }">
                Edit
              </router-link>
              <!-- &middot;
              <a @click.prevent="remove(customer.id)">
                Remove
              </a> -->
            </td>
          </tr>
          <tr v-if="customers.length == 0">
            <td colspan="5">
              No customers found.
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
import Search from '@/components/misc/Search'

export default {
  components: {
    Message,
    Search,
  },

  data () {
    return {
      customers: [],
      busy: false,
    }
  },

  computed: {
    filters () {
      return this.$store.getters.filters
    },

    filteredItems () {
      let items = this.customers

      if (this.filters.length) {
        items = this.customers.filter(c => {
          return this.filters.filter(v => {
            return c.number.includes(v)
              || c.address.company.toLowerCase().includes(v)
              || c.address.zip.includes(v)
          }).length > 0
        })
      }

      return items
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
        this.customers = await http.fetchCustomers()
      } catch (error) {
        this.setMessage({
          text: error.message,
          style: 'is-danger',
        })
      }
    },

    async remove (id) {
      if (!window.confirm('Are you sure?')) {
        return
      }

      try {
        await http.removeCustomer(id)
        this.load()
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
  },

  beforeRouteEnter (to, from, next) {
    next(vm => {
      if (from.name && !from.name.includes('customer')) {
        vm.$store.commit('filterValue', '')
      }
    })
  },
}
</script>

<style></style>
