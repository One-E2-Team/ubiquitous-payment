<template>
    <v-container>
      <v-row align="center" justify="center">
          <v-col cols="auto">
            <img width="60" height="60" src="../assets/successImage.png"/>
          </v-col>
      </v-row>
      <v-row align="center" justify="center">
          <v-col cols="auto">
            <h1>THANK YOU FOR YOUR PURCHASE!</h1>
          </v-col>
      </v-row>
      <v-row align="center" justify="center" align-content="center">
          <v-col align-self="center" cols="auto">
            <h3>Your order ID is <b>{{pspId}}</b>.</h3>
          </v-col>
      </v-row>
      <v-row align="center" justify="center">
          <v-col cols="auto">
           <v-btn color="success" outlined @click="goToMyOrders()">Check your orders</v-btn>
          </v-col>
      </v-row>
    </v-container>
</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
 export default {
    name: 'Success',
    mounted(){
        var pathParts = window.location.href.split("/");
        this.pspId = pathParts[pathParts.length - 1];

        axios({
                method: "put",
                url: comm.WSprotocol +'://' + comm.WSserver + '/api/psp-order/' + this.pspId + "/FULFILLED",
            }).then(response => {
              if(response.status==200){
                console.log(response.status);
              }
            }).catch(() => {
              console.log("error");
            })
    },
    data() {
        return{
            pspId : ''
        }
    },
    methods : {
      goToMyOrders(){
        this.$router.push({name: "MyOrders"});
      }
    }
 }
</script>
