<template>
  <v-container>
    <v-row class="text-center">
      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Choose a payment type
        </h1>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
  export default {
    name: 'HelloWorld',
    props: ['transactionId'],
    created(){
        console.log(this.transactionId);
        this.getPaymentTypes();
    },
    data() {return {
      paymentTypes: [],
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
  }
  }
</script>