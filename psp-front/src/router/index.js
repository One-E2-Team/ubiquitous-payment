import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ChoosePaymentType from '../components/ChoosePaymentType.vue'
import RegisterWebShop from '../components/RegisterWebShop.vue'
import Login from '../components/Login.vue'
import SetPaymentTypes from '../components/SetPaymentTypes.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/choose-payment-type/:transactionId',
    name: 'ChoosePaymentType',
    component: ChoosePaymentType
  },
  {
    path: '/register',
    name: 'RegisterWebShop',
    component: RegisterWebShop
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/payment-types',
    name: 'SetPaymentTypes',
    component: SetPaymentTypes
  },
]

const router = new VueRouter({
  routes
})

export default router
