import Vue from 'vue'
import VueRouter from 'vue-router'
import Welcome from '../components/Welcome.vue'
import Login from './../components/Login.vue'
import Register from './../components/Register.vue'
import HomePage from './../components/HomePage.vue'
import Success from './../components/Success.vue'
import Fail from './../components/Fail.vue'

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
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
