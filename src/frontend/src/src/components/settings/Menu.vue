<template>
  <aside class="menu">
    <ul
      v-for="(item, index) in items"
      :key="index"
      class="menu-list"
    >
      <li>
        <a
          href
          :class="['is-clearfix', { 'is-active': item.active }]"
          @click.prevent="$emit('select', item.view)"
        >
          <span class="is-pulled-left">
            {{ item.title }}
          </span>
          <span
            v-if="errorCount(item.view) > 0"
            class="tag is-pulled-right is-danger"
          >
            {{ errorCount(item.view) }}
          </span>
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

    errorCount (view) {
      const errors = this.errors.find(e => e.view === view)
      if (errors) {
        return errors.count
      }

      return 0
    }
  },
}
</script>
