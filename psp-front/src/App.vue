<template>
  <v-app>
    <v-app-bar app color="accent" dark>
    <v-container>
      <v-row justify="space-between">
        <v-col cols="auto">
          <h2>PSP</h2>
        </v-col>
        <v-col>
          <v-btn color="success" @click="goToRegisterPage()" v-if="getJwtToken() == null">Register WebShop</v-btn>
        </v-col>
        <v-col cols="auto">
            <v-btn @click="goToSetPaymentTypes()" color="primary" v-if="getJwtToken() != null">Set payment types</v-btn>
        </v-col>
         <v-col cols="auto">
            <v-btn @click="logout()" outlined v-if="getJwtToken() != null"><img width="30" height="30" src="./assets/logout.png"/></v-btn>
        </v-col>
        <v-col cols="auto">
            <v-btn @click="gotToLoginPage()" outlined  v-if="getJwtToken() == null">Log in</v-btn>
        </v-col>
      </v-row>
    </v-container>
    </v-app-bar>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
import * as comm from './configuration/communication.js'
import eventBus from './plugins/eventBus.js'
export default {
  name: 'App',

  data: () => ({
    sessionActivationCounter : 0,
    isUserLogged : true
  }),

  mounted() {
    eventBus.$on('login', () => {
      this.isUserLogged = true;
      this.checkSessionActivity();
    });
      eventBus.$on('logout', () => {
      this.isUserLogged = false;
    })
  },

  methods : {
    goToRegisterPage(){
      this.$router.push({name: "RegisterWebShop"})
    },
    logout() {
        comm.logOut();
        this.$router.push({name: "Welcome"});
        location.reload();
      },
    getJwtToken(){
      return comm.getJWTToken();
    },
    gotToLoginPage(){
      this.$router.push({name: "Login"});
    },
    goToSetPaymentTypes(){
      this.$router.push({name: "SetPaymentTypes"});
    },
    async checkSessionActivity(){
      var checkInterval = setInterval(function(){ 
          if (document.hidden) {
            this.sessionActivationCounter ++;
          }else{
            this.sessionActivationCounter = 0;
          }
          if (this.sessionActivationCounter == 200){
            sessionStorage.removeItem("JWT");
            clearInterval(checkInterval);
          }
      }, 3000);
    }
  }
};
</script>
