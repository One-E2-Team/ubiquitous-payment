<template>
<v-container>
  <v-row align="center" justify="center">
      <v-col cols="auto">
          <h2>Set payment types for your Web-shop</h2>
      </v-col>
  </v-row>
    <v-row v-for="po in paymentOptions" :key="po">
        <v-col></v-col>
        <v-col>
            <v-checkbox v-model="selected" :label="po" :value="po"></v-checkbox>
        </v-col>
        <v-col></v-col>
    </v-row>
    <v-row align="center" justify="center">
      <v-col cols="auto">
          <v-btn color="primary" @click="setPaymentTypes()">Save</v-btn>
      </v-col>
  </v-row>
  </v-container>

</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
  export default {
    data() {return {
      paymentOptions: [],
      selected: [],
    }},
    mounted(){
        this.getPaymentOptions();
    },
    methods: {
      getPaymentOptions() {
          axios({
            method: "get",
            url: comm.Protocol + "://" + comm.PSPserver +"/api/psp/payment-types/my",
            headers: comm.getHeader()
          }).then((response) => {
            if (response.status == 200) {
                this.paymentOptions = response.data.paymentOptions;
                this.selected = response.data.myPaymentOptions;
            }
          }).catch((response) => {
              console.log(response.data);
          })
        },

        setPaymentTypes(){
            axios({
            method: "put",
            url: comm.Protocol + "://" + comm.PSPserver +"/api/psp/payment-types",
            data: JSON.stringify(this.selected),
            headers: comm.getHeader()
          }).then((response) => {
            if (response.status == 200) {
                alert("Your payment types were successfully updated.")
            }
          }).catch(() => {
              alert("We are sorry, some error has occurred.")
          })
        }
      }
  }
</script>