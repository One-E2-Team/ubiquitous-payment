<template>
<v-container>
    <v-row><v-col></v-col></v-row>
    <v-row><h2>My ordered products</h2></v-row>
    <v-row justify="center">
        <v-col cols="auto">
            <v-data-table :headers="this.headers" :items="this.orderedProducts" class="elevation-1">
            </v-data-table>
        </v-col>
    </v-row>
</v-container>

</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
  export default {
    data() {return {
      headers: [
            { text: 'ID', value: 'orderId' },
            { text: 'PSP ID', value: 'pspId' },
            { text: 'Timestamp', value: 'timestamp' },
            { text: 'Product', value: 'productName' },
            { text: 'Price', value: 'productPrice' },
            { text: 'Currency', value: 'currency' },
            { text: 'Number of installments', value: 'numberOfInstallments' },
            { text: 'Delayed installments', value: 'delayedInstallments' },
            { text: 'Recurring type', value: 'recurringType' },
            { text: 'Order status', value: 'orderStatus' }
            
        ],
      orderedProducts: [],
    }},
    mounted() {
        this.getOrderedProducts();
    },
    methods: {
        getOrderedProducts(){
            axios({
                method: "get",
                url: comm.WSprotocol +'://' + comm.WSserver + '/api/my-orders/seller',
                headers: comm.getHeader()
            }).then(response => {
              if(response.status==200){
                this.orderedProducts = response.data;
              }
            }).catch(() => {
              console.log("error")
            })
        },
    }
  }
</script>