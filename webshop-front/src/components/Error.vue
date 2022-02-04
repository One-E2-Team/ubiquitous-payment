<template>
    <v-container>
      <v-row align="center" justify="center">
          <v-col cols="12">
            <template>
                <div class="red accent-4 text-center">
                    <span class="white--text">Purchase error</span>
                </div>
            </template>
          </v-col>
      </v-row>
      <v-row align="center" justify="center">
          <v-col cols="auto">
            <h2>Something went wrong! An error has occured with your order, ID: <b>{{pspId}}</b>.</h2>
          </v-col>
      </v-row>
      <v-row align="center" justify="center">
          <v-col cols="auto">
            <h3>Looks like we encountered an error. Please try again.</h3>
          </v-col>
      </v-row>
      <v-row align="center" justify="center">
          <v-col cols="auto">
            <h3>If you continue to have issues, try another payment method.</h3>
          </v-col>
      </v-row>
      <v-row align="center" justify="center">
          <v-col cols="auto">
           <v-btn color="primary" @click="goToHomePage()">Try again</v-btn>
          </v-col>
      </v-row>
    </v-container>
</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
 export default {
    name: 'Error',
    mounted(){
        var pathParts = window.location.href.split("/");
        this.pspId = pathParts[pathParts.length - 1];

        axios({
                method: "put",
                url: comm.WSprotocol +'://' + comm.WSserver + '/api/psp-order/' + this.pspId + "/ERROR",
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
        goToHomePage(){
            this.$router.push({name: "HomePage"});
        }
    }
 }
</script>