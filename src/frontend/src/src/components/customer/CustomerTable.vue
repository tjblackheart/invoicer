<template>
  <div class="table-container">
    <table
      v-if="filteredItems"
      class="table is-fullwidth"
    >
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
          <td> {{ customer.number }} </td>
          <td> {{ customer.address.company }} </td>
          <td> {{ customer.address.city }}, {{ customer.address.zip }} </td>

          <td class="has-text-right">
            <router-link :to="{ name: 'customer_details', params: {id:customer.id} }">
              Edit
            </router-link>
            <!-- &middot;
            <a @click.prevent="$emit('remove', customer.id)">
              Remove
            </a> -->
          </td>
        </tr>
        <tr v-if="filteredItems.length == 0">
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
