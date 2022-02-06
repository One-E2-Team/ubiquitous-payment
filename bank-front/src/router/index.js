import Vue from 'vue'
import VueRouter from 'vue-router'
import Payment from './../components/Payment.vue'
import Register from './../components/Register.vue'
import Login from './../components/Login.vue'
import MyProfile from './../components/MyProfile.vue'
import Transactions from './../components/Transactions.vue'

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/about',
    name: 'About',
    component: () =>
      import ( /* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/payment',
    name: 'Payment',
    component: Payment
  },
  {
    path: '/my-profile',
    name: 'MyProfile',
    component: MyProfile
  },
  {
    path: '/transactions',
    name: 'Transactions',
    component: Transactions
  },
]

const router = new VueRouter({
  routes
})

export default router