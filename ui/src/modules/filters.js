import Vue from 'vue'

Vue.filter('money', (val, cur = 'EUR') => {
  const locale = process.env.VUE_APP_LOCALE || 'de-DE'
  return (val).toLocaleString(locale, {
    style: 'currency',
    currency: cur,
  })
})

Vue.filter('date', val => {
  return new Date(val).toLocaleDateString()
})
