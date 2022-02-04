<template>
  <div id="nav">
    <v-app-bar color="blue" dense>
      <v-container>
        <v-row align="center" justify="center">
          <v-col cols="auto">
            <v-btn @click="goToHomePage()" class="mx-2">
              <v-icon large> mdi-home-outline</v-icon>
            </v-btn>
          </v-col>
          <v-col></v-col><v-col></v-col><v-col></v-col>
          <v-col v-if="hasRole('SELLER')" @click="goToPaymentOptions()" cols="auto"><v-btn>Payment options</v-btn></v-col>
          <v-col v-if="hasRole('CUSTOMER')" @click="goToMyOrders()" cols="auto"><v-btn>My orders</v-btn></v-col>
          <v-col cols="auto">
            <v-btn @click="logout()" class="mx-2"><img width="30" height="30" src="../assets/logout.png"/></v-btn>
            <v-spacer></v-spacer>
          </v-col>
        </v-row>
        </v-container>
        </v-app-bar>
  </div>
</template>

<script>
import * as comm from '../configuration/communication.js'
export default {
    name: "NavBar",
    mounted(){
    },
    methods: {
     goToHomePage(){
       this.$router.push({name: "HomePage"})
     },
     goToPaymentOptions(){
        this.$router.push({name: "PaymentOptions"})
     },
     goToMyOrders(){
        this.$router.push({name: "MyOrders"})
     },
     hasRole(role){
       return comm.hasRole(role);
     },
      logout() {
        comm.logOut();
        this.$router.push({name: "Welcome"});
      }
    }
}
</script>