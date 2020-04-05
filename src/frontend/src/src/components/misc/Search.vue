<template>
  <div class="content">
    <b-input
      id="search"
      v-model="filterValue"
      type="text"
      :placeholder="placeholder"
      @escape="filterValue = ''"
    />

    <div class="is-clearfix">
      <div class="is-pulled-left">
        <b-checkbox
          v-if="type === 'invoice'"
          v-model="showCancelled"
          label="Show cancelled invoices"
        />
      </div>
      <div
        v-if="hasFilters"
        class="is-pulled-right"
      >
        <a
          href
          @click.prevent="resetFilters"
        > Reset all filters </a>
      </div>
    </div>
  </div>
</template>

<script>
import BInput from '@/components/fields/Input'
import BCheckbox from '@/components/fields/Checkbox'

export default {

  components: {
    BInput,
    BCheckbox,
  },
  props: {
    placeholder: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      required: true,
    }
  },

  computed: {
    showCancelled: {
      get () {
        return this.$store.getters.showCancelled
      },
      set (val) {
        this.$store.commit('showCancelled', val)
      }
    },

    filterValue: {
      get () {
        return this.$store.getters.filterValue
      },
      set (val) {
        this.$store.commit('filterValue', val)
      }
    },

    filters () {
      return this.$store.getters.filters
    },

    hasFilters () {
      return this.showCancelled === false || this.filters.length
    }
  },

  methods: {
    resetFilters () {
      this.showCancelled = true
      this.filterValue = ''
    }
  }
}
</script>

<style>

</style>
