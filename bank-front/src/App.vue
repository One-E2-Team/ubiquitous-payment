<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
        />

        <v-img
          alt="Vuetify Name"
          class="shrink mt-1 hidden-sm-and-down"
          contain
          min-width="100"
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-name-dark.png"
          width="100"
        />
      </div>

      <v-spacer></v-spacer>

      <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
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
    },
  }
};
</script>
