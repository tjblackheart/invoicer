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

    <b-input
      type="text"
      id="search"
      v-model="filterValue"
      placeholder="Number, Name, Zipcode ..."
      @escape="resetSearch"
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
import BInput from '@/components/fields/Input'

export default {
  components: {
    Message,
    BInput,
  },

  data () {
    return {
      customers: [],
      busy: false,
      filterValue: '',
      filterValues: [],
      filteredItems: [],
    }
  },

  watch: {
    '$route': 'load',

    filterValue () {
      if (this.filterValue === '') {
        this.resetSearch()
        return
      }

      this.filterValues = this.filterValue.split(' ').filter(v => v.trim() !== '')
      this.search()
    },

    customers () {
      this.filteredItems = this.customers
    }
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

    search () {
      this.filteredItems = this.customers.filter(c => {
        return this.filterValues.filter(v => {
          return c.number.includes(v)
            || c.address.company.toLowerCase().includes(v)
            || c.address.zip.includes(v)
        }).length > 0
      })
    },

    resetSearch () {
      this.filterValue = ''
      this.filterValues = []
      this.filteredItems = this.customers
    }
  },
}
</script>

<style></style>
