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
          class="button is-primary"
        >
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

    <customer-table
      :paged-items="pagedItems"
      :filtered-items="filteredItems"
      @remove="remove($event)"
    />

    <hr>

    <pager
      v-if="pages > 1"
      :current="currentPage"
      :pages="pages"
      @paged="currentPage = $event"
    />
  </div>
</template>

<script>
import { mapMutations } from 'vuex'
import http from '@/modules/http'
import Message from '@/components/misc/Message'
import Search from '@/components/misc/Search'
import Pager from '@/components/misc/Pager'
import CustomerTable from '@/components/customer/CustomerTable'

export default {
  components: {
    Message,
    Search,
    Pager,
    CustomerTable,
  },

  data () {
    return {
      customers: [],
      busy: false,
      currentPage: 1,
      perPage: 10,
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
            return c.number.toLowerCase().includes(v)
              || c.address.company.toLowerCase().includes(v)
              || c.address.zip.includes(v)
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
  },

  methods: {
    ...mapMutations([ 'setMessage', 'clearMessage' ]),

    async load () {
      try {
        this.customers = await http.fetchCustomers()
        this.customers.sort((a, b) => a.id < b.id ? 1 : -1)
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
  },

  beforeRouteEnter (to, from, next) {
    next(vm => {
      if (from.name && !from.name.includes('customer')) {
        vm.$store.commit('filterValue', '')
        vm.$store.commit('showCancelled', true)
      }
    })
  },
}
</script>

<style></style>
