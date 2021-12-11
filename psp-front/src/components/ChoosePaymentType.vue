<template>
  <v-container>
    <v-row class="text-center">
      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Choose a payment type
        </h1>
      </v-col>
    </v-row>
    <v-row justify="center" v-for="p in this.paymentTypes" :key="p.name">
         <v-btn
                color="success"
                elevation="8"
                large
                @click="choosePaymentType(p)"
                >
                {{p}}
                </v-btn><br/>
        </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
  export default {
    name: 'HelloWorld',
    mounted(){
        var pathParts = window.location.href.split("/");
        this.transactionId = pathParts[pathParts.length - 1];
        this.getPaymentTypes();
    },
    data() {return {
      paymentTypes: [],
      transactionId : ''
    }},
    methods: {
     getPaymentTypes(){
       axios({
                method: "get",
                url: comm.Protocol +'://' + comm.PSPserver + '/api/psp/payments/' + this.transactionId,
            }).then(response => {
              if(response.status==200){
                this.paymentTypes = response.data;
              }
            }).catch((response) => {
              console.log(response.data)
            });
    
     }, 
     choosePaymentType(p){
         let data = {
             id : this.transactionId,
             name : p
         }
        axios({
                method: "post",
                url: comm.Protocol +'://' + comm.PSPserver + '/api/psp/select-payment',
                data : JSON.stringify(data)
            }).then(response => {
              if(response.status==200){
                console.log(response.data);
                window.location.href = response.data;
              }
            }).catch((response) => {
              console.log(response.data)
            });
    }
  }
  }
</script>