<template>
  <v-app>
    <v-app-bar app color="warning" dark>
      <v-container>
        <v-row justify="space-between">
          <v-col cols="auto">
            <h2>BANK</h2>
          </v-col>
          <v-col cols="auto">
            <v-btn @click="logout()" outlined v-if="getJwtToken() != null"
              ><img width="30" height="30" src="./assets/logout.png"
            /></v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import eventBus from "./plugins/eventBus.js";
import * as comm from "./configuration/communication.js";

export default {
  name: "App",

  data: () => ({
    sessionActivationCounter: 0,
    isUserLogged: true,
  }),

  mounted() {
    eventBus.$on("login", () => {
      this.isUserLogged = true;
      this.checkSessionActivity();
    });
    eventBus.$on("logout", () => {
      this.isUserLogged = false;
    });
  },

  methods: {
    logout() {
      comm.logOut();
      this.$router.push({ name: "Welcome" });
      location.reload();
    },
    getJwtToken() {
      return comm.getJWTToken();
    },
    async checkSessionActivity() {
      var checkInterval = setInterval(function () {
        if (document.hidden) {
          this.sessionActivationCounter++;
        } else {
          this.sessionActivationCounter = 0;
        }
        if (this.sessionActivationCounter == 200) {
          sessionStorage.removeItem("JWT");
          clearInterval(checkInterval);
        }
      }, 3000);
    },
  },
};
</script>
