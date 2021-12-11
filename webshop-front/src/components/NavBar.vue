<template>
  <div id="nav">
    <v-container>
      <v-row align="center" justify="center">
        <v-app-bar
          color="blue"
          dense
        >
        <v-col cols="12" sm="4">
          <v-btn @click="goToHomePage()" class="mx-2">
            <v-icon large>
              mdi-home-outline
            </v-icon>
          </v-btn>
          <v-spacer></v-spacer>
        </v-col>
        </v-app-bar>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
    name: "NavBar",
    mounted(){
    },
    methods: {
     goToHomePage(){
       this.$router.push({name: "HomePage"})
     },
      logout() {
      }
    }
}
</script>