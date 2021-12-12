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
     <v-row>
      <div v-if="isPaymentSelected">
        <v-row justify="center">
          <h2>You will be automaticly redirected when we register a transaction.</h2> 
        </v-row>
        <v-row justify="center">
          <v-progress-circular
            :size="80"
            color="primary"
            indeterminate
          ></v-progress-circular>
        </v-row>
      </div>
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
      transactionId : '',
      isPaymentSelected : false
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
       this.isPaymentSelected = true;
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
                window.open(response.data.redirectUrl, '_blank');
                if (p == "bitcoin"){
                  this.bitcoinAsyncFunc(data.id)
                }
              }
            }).catch((response) => {
              console.log(response.data)
            });
      },
      async bitcoinAsyncFunc(id){
          /*let kurc = function wait(ms) {
            var start = Date.now(),
                now = start;
            while (now - start < ms) {
              now = Date.now();
            }
          }*/
          const delay = ms => new Promise(res => setTimeout(res, ms));
          var k = true;
          while(k){
            //kurc(5000);
            await delay(5000);
            axios({
              method: "get",
              url: comm.Protocol +'://' + comm.PSPserver + '/api/psp/check-for-payment/bitcoin/' + id,
            }).then(resp => {
              if(resp.status == 200){
                if (resp.data.paymentCaptured){
                  window.location.href = resp.data.successUrl;
                }
              }
            })
          }
  }
  }}
</script>