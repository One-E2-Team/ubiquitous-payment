<template>
    <v-container>
    <v-row>
        <v-col> 
        </v-col>
    </v-row>
    <v-row>
        <v-col>
            <v-data-table :headers="this.headers" :items="this.myOrders" class="elevation-1">
                <template v-slot:top>
                <v-toolbar flat>
                    <v-toolbar-title>My orders</v-toolbar-title>
                    <v-divider class="mx-4" inset vertical></v-divider>
                    <v-spacer></v-spacer>
                </v-toolbar>
                </template>
            </v-data-table>
        </v-col>
    </v-row>
</v-container>
</template>

<script>
import axios from 'axios'
import * as comm from '../configuration/communication.js'
 export default {
    name: 'MyOrders',
    data() { return {
         myOrders: [],
         headers: [
            { text: 'ID', value: 'orderId' },
            { text: 'PSP ID', value: 'pspId' },
            { text: 'Timestamp', value: 'timestamp' },
            { text: 'Product', value: 'productName' },
            { text: 'Price', value: 'productPrice' },
            { text: 'Currency', value: 'currency' },
            { text: 'Payment type', value: 'paymentType' },
            { text: 'Number of installments', value: 'numberOfInstallments' },
            { text: 'Delayed installments', value: 'delayedInstallments' },
            { text: 'Recurring type', value: 'recurringType' },
            { text: 'Order status', value: 'orderStatus' }
            
        ]
        }
    },
    mounted(){
        axios({
                method: "get",
                url: comm.WSprotocol +'://' + comm.WSserver + '/api/my-orders',
                headers: comm.getHeader()
            }).then(response => {
              if(response.status==200){
                this.myOrders = response.data;
                console.log(response.data);
              }
            }).catch((response) => {
              console.log(response.data);
            });
    }
 }
</script>