<template>
  <transition name="modal">
    <div class="modal is-active">
      <div class="modal-background" />

      <div class="modal-content">
        <div class="box">
          <slot name="content" />
          <hr>

          <div class="has-text-right">
            <slot name="action" /> &nbsp;
            <button
              class="button"
              @click.prevent="$emit('close')">
              Cancel
            </button>
          </div>
        </div>
      </div>

      <button
        class="modal-close is-large"
        aria-label="close"
        @click.prevent="$emit('close')" />
    </div>
  </transition>
</template>

<script>
export default {
  created () {
    window.addEventListener('keyup', this.onEscListener)
  },

  destroyed () {
    window.removeEventListener('keyup', this.onEscListener)
  },

  methods: {
    onEscListener (e) {
      if (e.key === 'Escape' || e.keyCode === 27) {
        this.$emit('close')
      }
    },
  },
}
</script>

<style scoped>
.modal-content {
  max-height: 100vh !important;
  max-width: 90% !important;
}

.modal-background {
  background: rgba(10,10,10,.75);
}

.modal-enter,
.modal-leave-active {
  opacity: 0;
  transform: scale(1.05);
}

.modal-enter-active,
.modal-leave-active {
  transition: all .2s ease
}
</style>
