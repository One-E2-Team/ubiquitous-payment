<template>
  <v-container>
    <v-row class="text-center">
      <v-col class="mb-4">
        <h2 class="display-2 mb-3">
          Choose a payment type
        </h2>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col>
        <v-row>
          <v-col cols="auto" v-for="p in this.paymentTypes" :key="p.name">
            <v-btn v-if='p=="qrcode" && !isQrcode' @click="choosePaymentType(p)">Pay with QR code</v-btn>
          </v-col>
        </v-row>
        <v-row align="center" justify="center" v-if="isQrcode">
              <v-col cols="auto">
                <template>
                  <qrcode-vue :value="dataForQrCode" :size="size" level="H" />
                </template>
              </v-col>
              <v-col cols="6" sm="9">
                <v-form ref="form" v-model="valid" lazy-validation>
                  <v-row align="center" justify="center">
                    <v-col cols="6" sm="6">
                      <v-text-field
                        v-model="creditCard.pan"
                        label="PAN:"
                        :rules="[rules.required, rules.pan]"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col cols="6" sm="6">
                      <v-text-field
                        v-model="creditCard.holderName"
                        label="Holder name:"
                        :rules="[rules.required]"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-row>
                  <v-row align="center" justify="center">
                    <v-col cols="12" sm="6">
                      <v-text-field
                        v-model="creditCard.validUntil"
                        label="Valid until:"
                        :rules="[rules.required, rules.cardValid]"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-text-field
                        v-model="creditCard.cvc"
                        label="CVC:"
                        :rules="[rules.required]"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-form>
                <v-row align="center" justify="center">
                  <v-col cols="12" sm="6" class="d-flex justify-space-around mb-6">
                    <v-btn color="primary" @click="pay()"> Pay </v-btn>
                  </v-col>
                </v-row>
              </v-col>
        </v-row>
      </v-col>
      <v-col cols="3">
        <v-divider vertical="true" style="margin-left: 80px;"></v-divider>
      </v-col>
      <v-col cols="auto">
          <v-row justify="center" cols="auto" v-for="p in this.paymentTypes" :key="p.name" style="height: 100px;">
            <v-col >
              <v-container>
                <v-row>
                  <v-btn v-if='p=="paypal"' @click="choosePaymentType(p)"><img width="70px" height="70px" src="../assets/paypal-logo.png" /> PAYPAL</v-btn>
                </v-row>
                <v-row>
                  <v-btn v-if='p=="bitcoin"' @click="choosePaymentType(p)"><img widtch="50px" height="50px" src="../assets/bitcoin-icon.png" /> BITCOIN</v-btn>
                </v-row>
                <v-row>
                <v-btn v-if='p=="bank"' @click="choosePaymentType(p)"><img width="50px" height="50px" src="../assets/bank-icon.png" />BANK</v-btn>
                </v-row>
              </v-container>
            </v-col>
          </v-row>
      </v-col>
    </v-row>
     <v-row>
     </v-row>
      <div v-if="isPaymentSelected && !isQrcode">
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
  </v-container>
</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
import * as validator from "../plugins/validator"
import QrcodeVue from 'qrcode.vue'
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
      isPaymentSelected : false,
      redirectUrl : '',
      isQrcode : false,
      dataForQrCode : "",
      size : 200,
      valid: true,
      rules: validator.rules,
      creditCard: {
        pan: "",
        cvc: "",
        holderName: "",
        validUntil: "",
      },
    }},
    components: {
      QrcodeVue,
    },
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
                if (p == "qrcode"){
                  this.redirectUrl = response.data.redirectUrl;
                  this.getDataForQrCode();
                }
                else{
                  window.open(response.data.redirectUrl, '_blank');
                  if (p == "bitcoin"){
                    this.bitcoinAsyncFunc(data.id)
                  }
                }
              }
            }).catch((response) => {
              console.log(response.data)
            });
      },
      getDataForQrCode(){

         axios({
                method: "get",
                url: comm.Protocol +'://' + comm.PSPserver + '/api/transaction/qrcode/' + this.transactionId,
            }).then(response => {
              if(response.status==200){
                this.dataForQrCode = JSON.stringify(response.data);
                this.isQrcode = true;
              }
            }).catch((response) => {
              console.log(response.data)
            });
      },
      pay(){
        if (!this.$refs.form.validate()) {
        return;
      }
      axios({
        method: "post",
        url: this.redirectUrl,
        data: JSON.stringify(this.creditCard),
      }).then((response) => {
        if (response.status == 200) {
          console.log(response);
          window.location.href = response.data;
        }
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