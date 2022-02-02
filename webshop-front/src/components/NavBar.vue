<template>
  <div id="nav">
    <v-container>
      <v-row align="center" justify="space-between">
        <v-app-bar
          color="blue"
          dense
        >
        <v-row align="center" justify="space-between">
          <v-col>
            <v-btn @click="goToHomePage()" class="mx-2">
              <v-icon large>
                mdi-home-outline
              </v-icon>
            </v-btn>
            <v-spacer></v-spacer>
          </v-col>
          <v-col v-if="hasRole('SELLER')" @click="goToPaymentOptions()" cols="12" sm="2"><v-btn>Payment options</v-btn></v-col>
        </v-row>
        </v-app-bar>
      </v-row>
    </v-container>
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
     hasRole(role){
       return comm.hasRole(role);
     },
      logout() {
      }
    }
}
</script>