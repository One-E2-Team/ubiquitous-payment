<template>
    <v-container>
      <v-row align="center" justify="center">
          <v-col cols="12" sm="8">
            <h1>Order with ID {{pspId}} is successfully completed!</h1>
          </v-col>
      </v-row>
    </v-container>
</template>

<script>
 export default {
    name: 'Success',
    mounted(){
        var pathParts = window.location.href.split("/");
        this.pspId = pathParts[pathParts.length - 1];
    },
    data() {
        return{
            pspId : ''
        }
    }
 }
</script>
