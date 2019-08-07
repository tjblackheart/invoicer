<template>
  <aside class="menu">
    <ul
      v-for="(item, index) in items"
      :key="`${item.view}-${index}`"
      class="menu-list">
      <li>
        <a
          href
          :class="{ 'is-active': item.active, 'has-text-danger': hasError(item.view) }"
          @click.prevent="$emit('select', item.view)">
          {{ item.title }}
        </a>
      </li>
    </ul>
  </aside>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      required: true,
    },
    errors: {
      type: Array,
      default: () => [],
    },
  },

  methods: {
    hasError (view) {
      return this.errors.some(e => e.view === view && e.errors)
    },
  },
}
</script>
