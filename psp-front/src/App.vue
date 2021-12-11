<template>
  <v-app>
    <v-app-bar
      app
      color="accent"
      dark
    >
    <div>
      <h2>PSP</h2>
    </div>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  }),
};
</script>
