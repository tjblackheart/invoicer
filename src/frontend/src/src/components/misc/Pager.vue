<template>
  <nav
    class="pagination is-small"
    role="navigation"
    aria-label="pagination"
  >
    <a
      :disabled="current === 1"
      class="pagination-previous"
      @click.prevent="previous"
    >
      Previous
    </a>

    <a
      :disabled="current === pages"
      class="pagination-next"
      @click.prevent="next"
    >
      Next page
    </a>

    <ul class="pagination-list">
      <li
        v-for="p in pages"
        :key="p"
      >
        <a
          :class="['pagination-link', { 'is-current': current === p }]"
          :aria-label="`Page ${p}`"
          @click.prevent="$emit('paged', p)"
        > {{p}}
        </a>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  props: {
    current: {
      type: Number,
      required: true,
    },
    pages: {
      type: Number,
      required: true,
    }
  },

  methods: {
    next () {
      if (this.current < this.pages) {
        this.$emit('paged', this.current + 1)
      }
    },

    previous () {
      if (this.current > 1) {
        this.$emit('paged', this.current - 1)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  a { background: #fff; }
</style>
