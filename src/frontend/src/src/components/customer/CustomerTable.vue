<template>
  <div class="table-container">
    <table class="table is-fullwidth">
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
          v-for="customer in pagedItems"
          :key="customer.id"
        >
          <td data-label="Number">
            {{ customer.number }}
          </td>

          <td data-label="Company">
            {{ customer.address.company }}
          </td>

          <td data-label="City">
            {{ customer.address.city }}, {{ customer.address.zip }}
          </td>

          <td
            data-label="Actions"
            class="has-text-right action-items"
          >
            <span>
              <router-link :to="{ name: 'customer_details', params: {id:customer.id} }">
                Edit
              </router-link>
            </span>
            <!-- &middot;
            <a @click.prevent="$emit('remove', customer.id)">
              Remove
            </a> -->
          </td>
        </tr>
        <tr v-if="!filteredItems.length">
          <td colspan="5">
            Nothing found.
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  props: {
    pagedItems: {
      type: Array,
      default: () => [],
    },
    filteredItems: {
      type: Array,
      default: () => [],
    }
  }
}
</script>
