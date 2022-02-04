import Vue from 'vue'
import VueRouter from 'vue-router'
import Welcome from '../components/Welcome.vue'
import Login from './../components/Login.vue'
import Register from './../components/Register.vue'
import HomePage from './../components/HomePage.vue'
import Success from './../components/Success.vue'
import Fail from './../components/Fail.vue'
import Error from './../components/Error.vue'
import PaymentOptions from './../components/PaymentOptions.vue'
import MyOrders from './../components/MyOrders.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Welcome',
    component: Welcome
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/homePage',
    name: 'HomePage',
    component: HomePage
  },
  {
    path: '/order/success/:id',
    name: 'Success',
    component: Success
  },
  {
    path: '/order/fail/:id',
    name: 'Fail',
    component: Fail
  },
  {
    path: '/order/error/:id',
    name: 'Error',
    component: Error
  },
  {
    path: '/about',
    name: 'About',
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/paymentOptions',
    name: 'PaymentOptions',
    component: PaymentOptions
  },
  {
    path: '/my-orders',
    name: 'MyOrders',
    component: MyOrders
  }
]

const router = new VueRouter({
  routes
})

export default router
