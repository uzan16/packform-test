import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/orders',
      name: 'orders',
      component: () => import('../views/OrdersView.vue')
    },
    {
      path: '/:catchAll(.*)', // Unrecognized path automatically matches 404
      redirect: {
        name: 'orders'
      }
    }
  ]
})

export default router
